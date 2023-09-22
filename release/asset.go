package release

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

type Asset struct {
	URL                string    `json:"url"`
	Name               string    `json:"name"`
	ContentType        string    `json:"content_type"`
	Size               int       `json:"size"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadURL string    `json:"browser_download_url"`
}

func (a Asset) Get(path string) error {
	return GetFile(a.URL, path)
}

type Assets []Asset

func (assets Assets) URLs() []string {
	urls := make([]string, 0, len(assets))
	for _, a := range assets {
		urls = append(urls, a.BrowserDownloadURL)
	}

	return urls
}

func (assets Assets) Names() []string {
	names := make([]string, 0, len(assets))
	for _, a := range assets {
		names = append(names, a.Name)
	}

	return names
}

func (assets Assets) FindByName(name string) Asset {
	var result Asset
	for _, a := range assets {
		if a.Name == name {
			result = a
			break
		}
	}
	return result
}

func (assets Assets) GetURL(name string) string {
	a := assets.FindByName(name)
	return a.BrowserDownloadURL
}

func GetFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		log.Info().Err(err).Str("filepath", filepath).Msg("Create file error")
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		e := fmt.Errorf("bad status: %s", resp.Status)
		log.Info().Err(e).Str("url", url).Msg("cannot download from the given url")
		return e
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Info().Err(err).Str("url", url).Str("filepath", filepath).Msg("Cannot copy body to filepath")
		return err
	}

	return nil
}
