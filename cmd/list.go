package cmd

import (
	"github.com/practical-coder/target/release"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.Flags().String("repo", "dominikh/go-tools", "Github Repository Name")
	listCmd.Flags().String("format", "{{.Name}}\t{{.BrowserDownloadURL}}", "Format assets listing")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List release assets",
	Example: `
		tg list --repo='haproxytech/dataplaneapi'
		tg list --repo='dominikh/go-tools'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			log.Fatal().Err(err).Msg("Github repository name missing")
		}
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Info().Err(err).Msg("Format flag error")
		}
		r := release.NewRelease(repo, "")
		r.Setup()
		r.Assets.Format(format)
	},
}
