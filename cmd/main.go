package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:    "members",
				Aliases: []string{"m"},
				Usage:   "member related commands",
				Subcommands: []cli.Command{
					{
						Name:    "list",
						Aliases: []string{"l"},
						Usage:   "list all members",
						Action: func(c *cli.Context) error {
							fmt.Println("Listing members")
							return nil
						},
					},
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "add a member",
						Action: func(c *cli.Context) error {
							fmt.Println("Adding a member")
							return nil
						},
					},
					{
						Name:    "remove",
						Aliases: []string{"r"},
						Usage:   "remove a member",
						Action: func(c *cli.Context) error {
							fmt.Println("Removing a member")
							return nil
						},
					},
				},
			},
			{
				Name:    "contests",
				Aliases: []string{"c"},
				Usage:   "contest related commands",
				Subcommands: []cli.Command{
					{
						Name:    "list",
						Aliases: []string{"l"},
						Usage:   "list all contests",
						Action: func(c *cli.Context) error {
							fmt.Println("Listing contests")
							return nil
						},
					},
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "add a contest. This will create issues, on github and add a folder to work in",
						Action: func(c *cli.Context) error {
							fmt.Println("Adding a contest")
							return nil
						},
					},
					{
						Name:    "remove",
						Aliases: []string{"r"},
						Usage:   "remove a contest",
						Action: func(c *cli.Context) error {
							fmt.Println("Removing a contest")
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
