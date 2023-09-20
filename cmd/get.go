package cmd

import (
	"fmt"

	"github.com/practical-coder/target/release"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "GET release file",
	Run: func(cmd *cobra.Command, args []string) {
		r := release.NewRelease(
			"https://api.github.com/repos/dominikh/go-tools/releases/latest",
			"staticcheck_linux_amd64.tar.gz",
		)
		r.Setup()
		fmt.Println(r.TarURL)
	},
}
