package cmd

import (
	"fmt"
	"go/format"

	"github.com/albe2669/kattissues/internal/kattis"
	"github.com/spf13/cobra"
)

var addContestsCmd = &cobra.Command{
	Use:   "add",
	Short: "Add contest", // TODO: Add description
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contest, err := kattis.GetContest(args[0])
		if err != nil {
			cmd.PrintErrln(fmt.Sprintf("Error getting contest: %s", err))
			return
		}

		if contest.Status == kattis.ContestNotStarted {
			cmd.PrintErrln("Can't get problems as contest has not started")
			return
		}

		fmt.Println(contest)
	},
}
