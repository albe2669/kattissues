package cmd

import "github.com/spf13/cobra"

var ContestsCmd = &cobra.Command{
	Use:   "contests",
	Short: "Manage contests", // TODO: Add description
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ContestsCmd.AddCommand(addContestsCmd)
}
