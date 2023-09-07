package cmd

import (
	"context"
	"fmt"
	"sync"

	"github.com/chelnak/ysmrr"
	"github.com/spf13/cobra"

	"github.com/albe2669/kattissues/internal"
)

var addMembersCmd = &cobra.Command{
	Use:   "add",
	Short: "Add members to the repository",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		sm := ysmrr.NewSpinnerManager()
		defer sm.Stop()

		ctx := context.Background()

		client := internal.CreateGithubClient()
		for _, username := range args {
			wg.Add(1)
			go func(username string) {
				defer wg.Done()

				spinner := sm.AddSpinner(fmt.Sprintf("Adding %s", username))

				err := internal.AddMember(ctx, client, username)
				if err != nil {
					spinner.ErrorWithMessage(fmt.Sprintf("Error adding %s: %s", username, err))
					return
				}

				spinner.CompleteWithMessage(fmt.Sprintf("Added %s", username))
			}(username)
		}

		sm.Start()
		wg.Wait()
	},
}
