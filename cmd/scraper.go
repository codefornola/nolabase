package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/bhelx/nolabase/infra"
	"github.com/bhelx/nolabase/scraper"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/cobra"
)

var (
	jobId int
)

func init() {
	var scraperNames []string
	for name := range scraper.AllScrapers {
		scraperNames = append(scraperNames, strings.ToLower(name))
	}
	msg := fmt.Sprintf("Comma separated list of scrapers to apply to the command. Choose any or all of: (%s)", strings.Join(scraperNames, ","))
	scraperCommand.Flags().StringVar(&scrapers, "scrapers", "", msg)
	scraperCommand.Flags().IntVar(&jobId, "job-id", 0, "Use this option to run just a single job")
	rootCmd.AddCommand(scraperCommand)
}

var scraperCommand = &cobra.Command{
	Use:   "scraper",
	Short: "Control the scraper ",
	Long:  `scraper`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := infra.NewConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		manager, err := scraper.NewScraperManager(config.Database.String())
		if err != nil {
			log.Fatal(err)
		}

		if jobId != 0 {
			pool, err := pgxpool.Connect(context.Background(), config.Database.String())
			if err != nil {
				log.Fatal(err)
			}
			jobRepo := infra.NewJobRepo(pool)
			j, err := jobRepo.GetJob(jobId)
			if err != nil {
				log.Fatal(err)
			}
			newJob, err := manager.RunJob(j)
			if newJob != nil {
				log.Println("Job tried to enqueue new job ", newJob)
			}
			if err != nil {
				log.Fatal(err)
			}
		} else {
			var scraperNames []string
			for _, name := range strings.Split(scrapers, ",") {
				scraperNames = append(scraperNames, name)
			}
			manager.Enqueue(scraperNames)
		}
	},
}
