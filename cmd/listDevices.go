/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/redhat-et/jumpstarter/pkg/harness"
	"github.com/spf13/cobra"
)

// listDevicesCmd represents the listDevices command
var listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "Lists available devices",
	Long:  `Iterates over the available drivers and gets a list of devices.`,
	Run: func(cmd *cobra.Command, args []string) {
		driver := cmd.Flag("driver").Value.String()
		devices, err := harness.FindDevices(driver)
		handleErrorAsFatal(err)
		if cmd.Flag("only-names").Value.String() == "true" {
			printDeviceNames(devices)
			return
		}
		printDeviceTable(devices)

	},
}

func init() {
	rootCmd.AddCommand(listDevicesCmd)
	listDevicesCmd.Flags().StringP("driver", "d", "", "Only list devices for the specified driver")
	listDevicesCmd.Flags().Bool("only-names", false, "Only list the device names")
}

func printDeviceTable(devices []harness.Device) {
	color.Set(color.FgGreen)
	fmt.Println("Device Name\tSerial Number\tDriver\t\t\tVersion")
	color.Unset()
	for _, device := range devices {
		deviceName, err := device.Name()
		handleErrorAsFatal(err)
		deviceSerial, err := device.Serial()
		handleErrorAsFatal(err)
		deviceVersion, err := device.Version()
		fmt.Printf("%s\t%s\t%s\t%s\n",
			deviceName, deviceSerial, device.Driver().Name(), deviceVersion)
	}
}

func printDeviceNames(devices []harness.Device) {
	for _, device := range devices {
		deviceName, err := device.Name()
		handleErrorAsFatal(err)
		fmt.Printf("%s\n", deviceName)
	}
}
