package cmd

import (
	"github.com/randaalex/finance_bot/cmd/cli"
	"github.com/spf13/cobra"
)

var runBotCmd = &cobra.Command{
	Use: "bot",
	Run: cli.BotHandler,
}

func init() {
	rootCmd.AddCommand(runBotCmd)
}
