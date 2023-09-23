package cmd

import (
	"github.com/practical-coder/target/release"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List release assets",
	Example: `
		target list --repo='haproxytech/dataplaneapi'
		target list --repo='dominikh/go-tools'
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
