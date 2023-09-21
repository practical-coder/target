package cmd

import (
	"fmt"

	"github.com/practical-coder/target/release"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.Flags().String("repo", "dominikh/go-tools", "Github Repository Name")
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
		r := release.NewRelease(repo, "")
		r.Setup()
		for _, n := range r.Assets.Names() {
			fmt.Println(n)
		}
	},
}
