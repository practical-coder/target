package cmd

import (
	"fmt"

	"github.com/practical-coder/target/release"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.Flags().String("file", "staticcheck_linux_amd64.tar.gz", "Asset file name to download")
	getCmd.Flags().String("url", "https://api.github.com/repos/dominikh/go-tools/releases/latest", "Release URL")
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "GET release file",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatal().Err(err).Msg("Asset filename missing")
		}
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatal().Err(err).Msg("Release URL missing")
		}
		r := release.NewRelease(url, file)
		r.Setup()
		fmt.Println(r.TarURL)
	},
}
