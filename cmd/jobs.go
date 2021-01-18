package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/bhelx/nolabase/infra"
	"github.com/spf13/cobra"
)

func init() {
	jobsCommand.Flags().BoolVarP(&watch, "watch", "w", false, "repeatedly watch")
	jobsCommand.Flags().StringVarP(&sql, "sql", "s", "'select * from infra.jobs;'", "sql to run")

	rootCmd.AddCommand(jobsCommand)
}

var watch bool
var sql string

var jobsCommand = &cobra.Command{
	Use:   "jobs",
	Short: "View and control jobs",
	Long:  `jobs`,
	Run: func(cmd *cobra.Command, args []string) {
		scraperNames := strings.Split(scrapers, ",")
		for _, n := range scraperNames {
			fmt.Println(n)
		}

		config, err := infra.NewConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		var cargs []string
		cargs = append(cargs, config.Database.DatabaseName)
		cargs = append(cargs, "-c")
		cargs = append(cargs, sql)
		cargs = append([]string{"-n", "1", "/usr/bin/psql"}, cargs...)
		command := exec.Command("watch", cargs...)

		command.Stdout = os.Stdout
		err = command.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}
