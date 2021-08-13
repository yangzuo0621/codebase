package github

import (
	"context"
	"fmt"
	"os"

	hub "github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func CreateCommand() *cobra.Command {
	c := &cobra.Command{
		Use: "github",
	}
	c.AddCommand(createListCommand())
	c.AddCommand(createGetCommand())
	return c
}

func createListCommand() *cobra.Command {
	c := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logrus.New()
			pat := os.Getenv("PERSONAL_ACCESS_TOKEN")
			if pat == "" {
				logger.Fatal("PERSONAL_ACCESS_TOKEN is empty")
			}

			ctx := context.Background()
			ts := oauth2.StaticTokenSource(&oauth2.Token{
				AccessToken: pat,
			})
			tc := oauth2.NewClient(ctx, ts)
			client := hub.NewClient(tc)

			repos, _, err := client.Repositories.List(ctx, "", nil)
			if err != nil {
				logger.Fatalf("list repo: %s", err)
			}

			for _, repo := range repos {
				logger.Infof("%s", *repo.Name)
			}
			return nil
		},
	}
	return c
}

func createGetCommand() *cobra.Command {
	var (
		flagRepository string
		flagOwner      string
	)
	c := &cobra.Command{
		Use: "get",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logrus.New()
			pat := os.Getenv("PERSONAL_ACCESS_TOKEN")
			if pat == "" {
				logger.Fatal("PERSONAL_ACCESS_TOKEN is empty")
			}

			ctx := context.Background()
			ts := oauth2.StaticTokenSource(&oauth2.Token{
				AccessToken: pat,
			})
			tc := oauth2.NewClient(ctx, ts)
			client := hub.NewClient(tc)

			repo, _, err := client.Repositories.Get(ctx, flagOwner, flagRepository)
			if err != nil {
				logger.Fatalf("get repo %s: %s", flagRepository)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "ID= %d Name=%s\n", *repo.ID, *repo.Name)
			return nil
		},
	}

	c.Flags().StringVar(&flagRepository, "repository", "", "repository name")
	c.Flags().StringVar(&flagOwner, "owner", "", "owner name")
	c.MarkFlagRequired("repository")
	c.MarkFlagRequired("owner")
	return c
}
