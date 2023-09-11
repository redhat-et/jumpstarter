/*
Copyright © 2023 Miguel Angel Ajo Pelayo <majopela@redhat.com
*/
package cmd

import (
	"github.com/redhat-et/jumpstarter/pkg/runner"
	"github.com/spf13/cobra"
)

// powerCmd represents the listDevices command
var runPlaybookCmd = &cobra.Command{
	Use:   "run-playbook",
	Short: "Run a jumpstarter playbook",

	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			handleErrorAsFatal(err)
		}

		playbook := args[0]
		driver := cmd.Flag("driver").Value.String()
		disableCleanup, _ := cmd.Flags().GetBool("disable-cleanup")

		err := runner.RunPlaybook("", driver, playbook, disableCleanup)
		handleErrorAsFatal(err)
	},
}

func init() {
	rootCmd.AddCommand(runPlaybookCmd)
	runPlaybookCmd.Flags().StringP("driver", "d", "", "Only run on devices for the specified driver")
	runPlaybookCmd.Flags().BoolP("disable-cleanup", "c", false, "Disable the cleanup phase if something goes wrong")
}
