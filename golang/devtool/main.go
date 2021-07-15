package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yangzuo0621/codebase/golang/devtool/github"
)

func createCommand() *cobra.Command {
	c := &cobra.Command{
		Use: "devtool",
	}
	return c
}

func main() {
	rootCmd := createCommand()
	rootCmd.AddCommand(github.CreateCommand())
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
