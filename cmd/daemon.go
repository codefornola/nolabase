package cmd

import (
	"log"

	"github.com/bhelx/nolabase/infra"
	"github.com/bhelx/nolabase/scraper"
	"github.com/spf13/cobra"
)

var concurrency int

func init() {
	daemonCommand.Flags().IntVarP(&concurrency, "concurrency", "c", 8, "The max number of workers in the pool.")
	rootCmd.AddCommand(daemonCommand)
}

var daemonCommand = &cobra.Command{
	Use:   "daemon",
	Short: "Control the scraper worker pool",
	Long:  `This command allows you to start, stop, and check status of the scraper worker pool daemon.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := infra.NewConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		manager, err := scraper.NewScraperManager(config.Database.String())
		if err != nil {
			log.Fatal(err)
		}

		jobs := make(chan infra.Job, 2)
		go manager.StartDatabaseEnqueuer(jobs)
		manager.StartSchedulerDaemon(concurrency, jobs)
	},
}
