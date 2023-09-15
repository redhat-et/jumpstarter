/*
Copyright © 2023 Miguel Angel Ajo Pelayo <majopela@redhat.com
*/
package cmd

import (
	"fmt"
	"regexp"

	"github.com/redhat-et/jumpstarter/pkg/harness"
	"github.com/spf13/cobra"
)

// powerCmd represents the listDevices command
var createAnsibleInventory = &cobra.Command{
	Use:   "create-ansible-inventory",
	Short: "Create an ansible inventory file",

	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			handleErrorAsFatal(err)
		}

		driver := cmd.Flag("driver").Value.String()
		user := cmd.Flag("user").Value.String()

		device, err := harness.FindDevice(driver, args[0])
		serial, err := device.Console()
		handleErrorAsFatal(err)

		result, err := runCommand(serial, "ip a show dev eth0", 2)
		handleErrorAsFatal(err)

		ip, err := extractSrcIPAddress(result)

		fmt.Print("---\nboards:\n  hosts:\n")
		fmt.Printf("    %s:\n", device.Name())
		fmt.Printf("      ansible_host: %s\n", ip)
		fmt.Printf("      ansible_user: %s\n", user)
		fmt.Printf("      ansible_become: yes\n")
		fmt.Printf("      ansible_ssh_common_args: '-o StrictHostKeyChecking=no'\n")

	},
}

func init() {
	rootCmd.AddCommand(createAnsibleInventory)
	createAnsibleInventory.Flags().StringP("driver", "d", "", "Only devices for the specified driver")
	createAnsibleInventory.Flags().StringP("user", "u", "root", "The user for the ansible inventory file")
}

func extractSrcIPAddress(input string) (string, error) {
	re := regexp.MustCompile(`inet (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return matches[1], nil
	} else {
		return "", fmt.Errorf("No src IP address found")
	}
}
