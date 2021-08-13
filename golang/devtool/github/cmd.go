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
	c.AddCommand(createPullRequestCommand())
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
				logger.Fatalf("list repo failed: %s", err)
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
				logger.Fatalf("get repo %s failed: %s", flagRepository, err)
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

// https://docs.github.com/en/rest/reference/pulls#create-a-pull-request
// ./devtool github pr-create --owner yangzuo0621 --repository codebase --base main --head zuya/create-pr-command --title "add create pr command" --body "add create pull request command"
func createPullRequestCommand() *cobra.Command {
	var (
		flagRepository string
		flagOwner      string
		flagBase       string
		flagHead       string
		flagTitle      string
		flagBody       string
	)
	c := &cobra.Command{
		Use: "pr-create",
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

			pr, _, err := client.PullRequests.Create(ctx, flagOwner, flagRepository, &hub.NewPullRequest{
				Title: &flagTitle,
				Head:  &flagHead,
				Base:  &flagBase,
				Body:  &flagBody,
			})
			if err != nil {
				logger.Fatalf("create PR failed: %s", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "ID= %d Number=%d\n", *pr.ID, *pr.Number)
			return nil
		},
	}

	c.Flags().StringVar(&flagRepository, "repository", "", "repository name")
	c.Flags().StringVar(&flagOwner, "owner", "", "owner name")
	c.Flags().StringVar(&flagBase, "base", "", "The name of the branch you want the changes pulled into.")
	c.Flags().StringVar(&flagHead, "head", "", "The name of the branch where your changes are implemented.")
	c.Flags().StringVar(&flagTitle, "title", "", "The name of pull request.")
	c.Flags().StringVar(&flagBody, "body", "", "The body of pull request.")
	c.MarkFlagRequired("repository")
	c.MarkFlagRequired("owner")
	c.MarkFlagRequired("base")
	c.MarkFlagRequired("head")
	return c
}
