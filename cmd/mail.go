package cmd

import (
	"github.com/randaalex/finance_bot/cmd/cli"
	"github.com/spf13/cobra"
)

var runMailCmd = &cobra.Command{
	Use: "mail",
	Run: cli.MailHandler,
}

func init() {
	rootCmd.AddCommand(runMailCmd)
}
