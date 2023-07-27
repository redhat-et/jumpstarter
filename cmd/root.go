/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jumpstarter",
	Short: "Jumpstarter is a tool to perform HIL testing on Edge and Embedded devices",
	Long: `This commandline tool is the interface to the jumpstarter drivers
which enable the control of Edge and Embedded devices from a host.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

func handleErrorAsFatal(err error) {
	if err != nil {
		color.Set(color.FgRed)
		fmt.Printf("Error: %s\n", err)
		color.Unset()
		os.Exit(1)
	}
}
