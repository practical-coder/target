package release

import "time"

type Asset struct {
	URL                string    `json:"url"`
	Name               string    `json:"name"`
	ContentType        string    `json:"content_type"`
	Size               int       `json:"size"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadURL string    `json:"browser_download_url"`
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

func (assets Assets) GetURL(name string) string {
	var url string
	for _, a := range assets {
		if a.Name == name {
			url = a.BrowserDownloadURL
			break
		}
	}
	return url
}
