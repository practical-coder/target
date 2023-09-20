package main

import (
	"github.com/practical-coder/target/release"
)

func main() {
	r := release.NewRelease(
		"https://api.github.com/repos/dominikh/go-tools/releases/latest",
		"staticcheck_linux_amd64.tar.gz",
	)
	r.Setup()
	// tarURL := r.TarURL()
	// Get ReleaseUrl
	// Grep browser_download_url with TarName
	// Get that file
	// unpack tar archive
}
