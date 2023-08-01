/*
Copyright © 2023 Miguel Angel Ajo Pelayo <majopela@redhat.com
*/
package cmd

import (
	"github.com/redhat-et/jumpstarter/pkg/harness"
	"github.com/spf13/cobra"
)

// powerCmd represents the listDevices command
var setDiskImageCmd = &cobra.Command{
	Use:   "set-disk-image",
	Short: "Writes a disk image to the target device attached storage",

	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			handleErrorAsFatal(err)
		}

		driver := cmd.Flag("driver").Value.String()
		device, err := harness.FindDevice(driver, args[0])
		handleErrorAsFatal(err)

		err = device.SetDiskImage(args[1])
		handleErrorAsFatal(err)

	},
}

func init() {
	rootCmd.AddCommand(setDiskImageCmd)
	setDiskImageCmd.Flags().StringP("driver", "d", "", "Only list devices for the specified driver")
	// add fixed possition argument device-id

}
