package internal

import (
	"context"

	"github.com/google/go-github/v55/github"
)

func CreateGithubClient() *github.Client {
	return github.
		NewClient(nil).
		WithAuthToken(Config.Github.Token)
}

func AddMember(ctx context.Context, client *github.Client, username string) error {
	owner := Config.Github.RepoOwner
	repo := Config.Github.RepoName

	opts := &github.RepositoryAddCollaboratorOptions{
		Permission: "push",
	}

	_, _, err := client.Repositories.AddCollaborator(ctx, owner, repo, username, opts)
	return err
}

func RemoveMember(ctx context.Context, client *github.Client, username string) error {
	owner := Config.Github.RepoOwner
	repo := Config.Github.RepoName

	_, err := client.Repositories.RemoveCollaborator(ctx, owner, repo, username)
	return err
}

func ListMembers(ctx context.Context, client *github.Client) ([]*github.User, error) {
	owner := Config.Github.RepoOwner
	repo := Config.Github.RepoName

	members, _, err := client.Repositories.ListCollaborators(ctx, owner, repo, nil)
	return members, err
}
