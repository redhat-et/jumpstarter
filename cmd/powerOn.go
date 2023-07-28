/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/redhat-et/jumpstarter/pkg/harness"
	"github.com/spf13/cobra"
)

// powerCmd represents the listDevices command
var powerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Powers device on",

	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			handleErrorAsFatal(err)
		}

		driver := cmd.Flag("driver").Value.String()
		device, err := harness.FindDevice(driver, args[0])
		handleErrorAsFatal(err)

		err = device.PowerOn()
		handleErrorAsFatal(err)

	},
}

func init() {
	rootCmd.AddCommand(powerOnCmd)
	powerOnCmd.Flags().StringP("driver", "d", "", "Only list devices for the specified driver")
	// add fixed possition argument device-id

}
