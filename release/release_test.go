package release

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExtractTarURL(t *testing.T) {
	testFile := "../testdata/latest_releases.json"
	data := LoadFile(t, testFile)
	list := make(map[string]any)
	err := json.Unmarshal(data, &list)
	if err != nil {
		t.Errorf("Cannot unmarshal test file: %s", testFile)
	}
	r := NewRelease("", "staticcheck_linux_amd64.tar.gz")
	r.SetList(list)
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
