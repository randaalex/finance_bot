package cmd

import (
	"fmt"
	"github.com/randaalex/finance_bot/cmd/cli"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "finance_bot",
}

var runBotCmd = &cobra.Command{
	Use: "bot",
	Run: cli.BotHandler,
}

var runCronCmd = &cobra.Command{
	Use: "cron",
	Run: cli.CronHandler,
}

var runMailCmd = &cobra.Command{
	Use: "mail",
	Run: cli.MailHandler,
}

var parseCmd = &cobra.Command{
	Use: "parse",
	Run: cli.ParseFileHandler,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	SetupCloseHandler()
	cobra.CheckErr(rootCmd.Execute())
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\rCtrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(runBotCmd)
	rootCmd.AddCommand(runCronCmd)
	rootCmd.AddCommand(runMailCmd)
	rootCmd.AddCommand(parseCmd)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.SetEnvPrefix("")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()
}
