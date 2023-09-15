/*
Copyright © 2023 Miguel Angel Ajo Pelayo <majopela@redhat.com
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/redhat-et/jumpstarter/pkg/harness"
	"github.com/spf13/cobra"
)

// powerCmd represents the listDevices command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command via console",

	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			handleErrorAsFatal(err)
		}

		driver := cmd.Flag("driver").Value.String()
		wait, err := cmd.Flags().GetInt("wait")
		handleErrorAsFatal(err)

		device, err := harness.FindDevice(driver, args[0])
		serial, err := device.Console()
		handleErrorAsFatal(err)

		result, err := runCommand(serial, strings.Join(args[1:], " "), wait)
		handleErrorAsFatal(err)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("driver", "d", "", "Only devices for the specified driver")
	runCmd.Flags().IntP("wait", "w", 2, "Wait seconds before trying to get a response")
}

func runCommand(console harness.ConsoleInterface, cmd string, wait int) (string, error) {
	buf := make([]byte, 1024)
	// clear the input buffer first
	console.SetReadTimeout(20 * time.Millisecond)
	console.Read(buf)

	if _, err := console.Write([]byte(cmd + "\n")); err != nil {
		return "", fmt.Errorf("runCommand %s, sending command: %w", cmd, err)
	}

	time.Sleep(time.Duration(wait) * time.Second)
	n, err := console.Read(buf)
	if err != nil {
		return "", fmt.Errorf("runCommand %s, reading response: %w", cmd, err)
	}

	all := string(buf[:n])
	lines := strings.Split(all, "\n")
	if len(lines) > 1 {
		// the first line is the command we just sent, so we skip it
		return strings.Join(lines[1:], "\n"), nil
	} else {
		return "", nil
	}
}
