package cmd

import (
	"fmt"

	"github.com/practical-coder/target/release"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.Flags().String("repo", "dominikh/go-tools", "Github Repository Name")
	getCmd.Flags().String("file", "staticcheck_linux_amd64.tar.gz", "Asset file name to download")
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get / Download release asset file",
	Example: `
		tg get --repo='haproxytech/dataplaneapi' --file='dataplaneapi_2.8.1_linux_x86_64.tar.gz'
		tg get --repo='dominikh/go-tools' --file='staticcheck_linux_amd64.tar.gz'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatal().Err(err).Msg("Asset filename missing")
		}
		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			log.Fatal().Err(err).Msg("Github repository name missing")
		}
		r := release.NewRelease(repo, file)
		r.Setup()
		a := r.Assets.FindByName(r.TarName)
		a.Get(fmt.Sprintf("./%s", r.TarName))
	},
}
