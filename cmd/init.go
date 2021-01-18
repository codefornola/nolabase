package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/markbates/pkger"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize the system",
	Long:  `Initializes the system. Run once after install.`,
	Run: func(cmd *cobra.Command, args []string) {
		seed, err := pkger.Open("/data/config/config.seed.yaml")
		if err != nil {
			log.Fatal(err)
		}
		defer seed.Close()

		buf := bytes.NewBuffer(nil)
		io.Copy(seed, buf)
		fmt.Println(buf.String())

		dir, err := homedir.Dir()
		if err != nil {
			log.Fatal("Could not find home directory")
		}

		path := path.Join(dir, ".nolabase.yaml")
		f, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = io.Copy(seed, f)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Wrote config file to ", path)
	},
}
