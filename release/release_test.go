package release

import (
	"os"
	"testing"

	"slices"

	"github.com/stretchr/testify/assert"
)

func TestExtractTarURL(t *testing.T) {
	testFile := "../testdata/latest_releases.json"
	data := LoadFile(t, testFile)
	r := NewRelease("", "staticcheck_linux_amd64.tar.gz")
	assets := r.GetAssets(data)
	r.SetAssets(assets)
	tarURL := r.ExtractTarURL()
	t.Logf("tarURL: %v", tarURL)
	assert.Equal(t, tarURL, "https://github.com/dominikh/go-tools/releases/download/2023.1.6/staticcheck_linux_amd64.tar.gz")
}

func LoadFile(t *testing.T, name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		t.Errorf("Cannot open test file: %s", name)
	}
	return data
}

func TestList(t *testing.T) {
	r := NewRelease("haproxytech/dataplaneapi", "")
	r.Setup()
	assert.True(t, slices.Contains(r.Assets.Names(), "dataplaneapi_2.8.1_linux_x86_64.tar.gz"))
}
