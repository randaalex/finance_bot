package cmd

import (
	"github.com/randaalex/finance_bot/cmd/cli"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use: "parse",
	Run: cli.ParseFileHandler,
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
