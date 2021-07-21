package cli

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
	"github.com/spf13/cobra"
)

func CronHandler(cmd *cobra.Command, args []string) {
	c := cron.New()
	_ = c.AddFunc("*/5 * * * * *", mailHandler)
	c.Start()

	fmt.Println("Cron scheduler started:")

	for {
		time.Sleep(10 * time.Second)
	}
}

func mailHandler() {
	fmt.Println("Every 5 secs job")
}
