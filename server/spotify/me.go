package spotify

import (
	"net/http"
)

//go:generate ffjson $GOFILE

// ffjson: skip
type MeRequest struct {
	AccessToken string
}

// ffjson: noencoder
type MeResponse struct {
	DisplayName  string            `json:"display_name"`
	Email        string            `json:"email"`
	ExternalURLs map[string]string `json:"external_urls"`
	HREF         string            `json:"href"`
	ID           string            `json:"id"`
	Images       []struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
	Product string `json:"product"`
	Type    string `json:"type"`
	URI     string `json:"uri"`
}

func (r MeRequest) Send() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+r.AccessToken)
	return client.Do(req)
}
