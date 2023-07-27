/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/redhat-et/jumpstarter/pkg/harness"
	"github.com/spf13/cobra"
)

// listDriversCmd represents the listDrivers command
var listDriversCmd = &cobra.Command{
	Use:   "list-drivers",
	Short: "List available HIL drivers",
	Run: func(cmd *cobra.Command, args []string) {
		drivers := harness.GetDrivers()
		for _, driver := range drivers {
			fmt.Printf("%s\n\t%s\n", driver.Name(), driver.Description())
		}
	},
}

func init() {
	rootCmd.AddCommand(listDriversCmd)
}
