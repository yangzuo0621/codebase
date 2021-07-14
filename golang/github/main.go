package main

import (
	"context"
	"os"

	"github.com/google/go-github/v37/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func main() {
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
	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		logger.Fatalf("list repo: %s", err)
	}

	for _, repo := range repos {
		logger.Infof("%s", *repo.Name)
	}
}
