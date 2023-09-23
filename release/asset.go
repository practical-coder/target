package release

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"text/template"
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
	return GetFile(a.BrowserDownloadURL, path)
}

func (a Asset) Format(format string) {
	tmpl, err := template.New("asset").Parse(format)
	if err != nil {
		log.Logger.Info().Err(err).Msg("asset template error")
	}
	err = tmpl.Execute(os.Stdout, a)
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("asset template execute error")
	}
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

func (assets Assets) FindByPattern(pattern string) Asset {
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal().Err(err).Str("pattern", pattern).Msg("FindByPattern error")
	}
	var result Asset
	for _, a := range assets {
		if r.MatchString(a.Name) {
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

func (assets Assets) Format(format string) {
	for _, a := range assets {
		a.Format(format)
		fmt.Println()
	}
}
