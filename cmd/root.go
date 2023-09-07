package cmd

import (
	"fmt"
	contests "github.com/albe2669/kattissues/cmd/contests"
	members "github.com/albe2669/kattissues/cmd/members"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kattissues",
	Short: "kattissues is a CLI tool for managing Kattis contests through Github Issues",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(contests.ContestsCmd)
	rootCmd.AddCommand(members.MembersCmd)
}
