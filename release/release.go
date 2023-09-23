package release

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Release struct {
	Repo    string
	URL     string
	TarName string
	TarURL  string
	Assets  Assets
}

func NewRelease(repo string, tarName string) *Release {
	return &Release{
		Repo:    repo,
		URL:     SetURL(repo),
		TarName: tarName,
	}
}

func SetURL(repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
}

func (r *Release) GetList() []byte {
	empty := make([]byte, 0)
	res, err := http.Get(r.URL)
	if err != nil {
		log.Info().Err(err).Msg("Cannot get latest release list")
		return empty
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Info().Err(err).Msg("Cannot read response body")
		return empty
	}
	return data
}

func (r *Release) GetAssets(data []byte) []Asset {
	results := make([]Asset, 0)
	var tempRelease Release
	err := json.Unmarshal(data, &tempRelease)
	if err != nil {
		log.Info().Err(err).Msg("Cannot unmarshal data")
		return results
	}
	results = tempRelease.Assets
	return results
}

func (r *Release) SetAssets(assets []Asset) {
	r.Assets = assets
}

func (r *Release) SetTarURL(url string) {
	r.TarURL = url
}

func (r *Release) ExtractTarURL() string {
	return r.Assets.GetURL(r.TarName)
}

func (r *Release) Setup() {
	data := r.GetList()
	assets := r.GetAssets(data)
	r.SetAssets(assets)
}
