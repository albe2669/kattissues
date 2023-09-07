package cmd

import "github.com/spf13/cobra"

var MembersCmd = &cobra.Command{
	Use:   "members",
	Short: "Manage members of the repository",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	MembersCmd.AddCommand(addMembersCmd)
}
