package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		getCmd,
		listCmd,
		versionCmd,
	)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	rootCmd.PersistentFlags().String("repo", "haproxytech/dataplaneapi", "Github Repository Name")
	rootCmd.PersistentFlags().String("format", "{{.Name}}", "Format assets listing")
}

var rootCmd = &cobra.Command{
	Use:   "target",
	Short: "TarGET - download latest github project release",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "TarGET version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TarGET version 0.0.1")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Err(err)
		os.Exit(1)
	}
}
