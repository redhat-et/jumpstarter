/*
Copyright © 2023 Miguel Angel Ajo Pelayo <majopela@redhat.com
*/
package cmd

import (
	"github.com/redhat-et/jumpstarter/pkg/harness"
	"github.com/spf13/cobra"
)

// powerCmd represents the listDevices command
var powerOffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Powers device off",

	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			handleErrorAsFatal(err)
		}

		driver := cmd.Flag("driver").Value.String()
		device, err := harness.FindDevice(driver, args[0])
		handleErrorAsFatal(err)

		err = device.PowerOff()
		handleErrorAsFatal(err)

	},
}

func init() {
	rootCmd.AddCommand(powerOffCmd)
	powerOffCmd.Flags().StringP("driver", "d", "", "Only list devices for the specified driver")

}
