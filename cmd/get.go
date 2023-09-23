package cmd

import (
	"fmt"

	"github.com/practical-coder/target/release"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.Flags().String("file", "", "Asset file name to download")
	getCmd.Flags().String("pattern", "", "Regular expression pattern matching asset file name to download")
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get / Download release asset file",
	Example: `
		target get --repo='haproxytech/dataplaneapi' --file='dataplaneapi_2.8.1_linux_x86_64.tar.gz'
		target get --repo='dominikh/go-tools' --file='staticcheck_linux_amd64.tar.gz'
	`,
	Run: func(cmd *cobra.Command, args []string) {

		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			log.Fatal().Err(err).Msg("Github repository name missing")
		}
		pattern, err := cmd.Flags().GetString("pattern")
		if pattern != "" && err != nil {
			log.Info().Err(err).Msg("Pattern flag error")
		}

		file, err := cmd.Flags().GetString("file")
		if file != "" && err != nil {
			log.Fatal().Err(err).Msg("Asset filename error")
		}

		if file != "" && pattern != "" {
			log.Fatal().Err(err).
				Str("file", file).
				Str("pattern", pattern).
				Msg("--file and --pattern provided! Only one of those flags can be in use")
		}

		r := release.NewRelease(repo, file)
		r.Setup()

		if file != "" {
			a := r.Assets.FindByName(file)
			a.Get(fmt.Sprintf("./%s", a.Name))
			return
		}

		if pattern != "" {
			a := r.Assets.FindByPattern(pattern)
			a.Get(fmt.Sprintf("./%s", a.Name))
		}
	},
}
