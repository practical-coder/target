package release

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Release struct {
	Repo    string
	URL     string
	TarName string
	TarURL  string
	Assets  []Asset
}

type Asset struct {
	URL                string    `json:"url"`
	Name               string    `json:"name"`
	ContentType        string    `json:"content_type"`
	Size               int       `json:"size"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadURL string    `json:"browser_download_url"`
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
	var tarURL string
	for _, a := range r.Assets {
		if a.Name == r.TarName {
			tarURL = a.BrowserDownloadURL
			break
		}
	}
	return tarURL
}

func (r *Release) Setup() {
	data := r.GetList()
	assets := r.GetAssets(data)
	r.SetAssets(assets)
	url := r.ExtractTarURL()
	r.SetTarURL(url)
}

func (r *Release) ListAssets() []string {
	urls := make([]string, 0, len(r.Assets))
	for _, a := range r.Assets {
		urls = append(urls, a.Name)
	}

	return urls
}
