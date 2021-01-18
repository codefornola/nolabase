package cmd

import (
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	cfgFile  string
	logLevel string
	logOut   string
	scrapers string
	rootCmd  = &cobra.Command{
		Use:   "nolabase",
		Short: "The tool to control and maintain a nolabase instance",
		Long:  `Tools for controlling and maintaining a nolabase instance such as the scrapers and loaders.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Could not find home directory for this platform", err)
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", path.Join(home, ".nolabase.yaml"), "Path to the yaml config file.")

	//rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "The level to run the logger. Available: (panic,fatal,error,warn,info,debug,trace)")
	// level, err := log.ParseLevel(logLevel)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.SetLevel(log.DebugLevel)

	rootCmd.PersistentFlags().StringVar(&logOut, "log-output", "stdout", "The output location for the logger")

	if logOut == "stdout" {
		log.SetOutput(os.Stdout)
	} else {
		file, err := os.OpenFile(logOut, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Fatal(err)
		}
	}

	log.SetFormatter(&prefixed.TextFormatter{})
}
