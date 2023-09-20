package release

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type Release struct {
	URL      string
	TarName  string
	TarURL   string
	FileList map[string]any
}

func NewRelease(url string, tarName string) *Release {
	return &Release{
		URL:     url,
		TarName: tarName,
	}
}

func (r *Release) GetList() map[string]any {
	results := make(map[string]any)
	res, err := http.Get(r.URL)
	if err != nil {
		log.Info().Err(err).Msg("Cannot get latest release list")
		return results
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Info().Err(err).Msg("Cannot read response body")
		return results
	}

	err = json.Unmarshal(data, &results)
	if err != nil {
		log.Info().Err(err).Msg("Cannot unmarshal data")
	}
	return results
}

func (r *Release) SetList(fileList map[string]any) {
	r.FileList = fileList
}

func (r *Release) SetTarURL(url string) {
	r.TarURL = url
}

func (r *Release) ExtractTarURL() string {
	assets := r.FileList["assets"].([]any)
	var tarURL string
	for _, a := range assets {
		asset := a.(map[string]any)
		url := asset["browser_download_url"].(string)
		if strings.HasSuffix(url, r.TarName) {
			tarURL = url
			break
		}
	}
	return tarURL
}

func (r *Release) Setup() {
	list := r.GetList()
	r.SetList(list)
	url := r.ExtractTarURL()
	r.SetTarURL(url)
}
