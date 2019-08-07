package spotify

import (
	"net/http"
	"net/url"
)

//go:generate ffjson $GOFILE

type Request interface {
	Send() (*http.Response, error)
}

// ffjson: skip
type AccessTokenRequest struct {
	ClientID     string
	ClientSecret string
	Code         string
	RedirectURI  string
}

// ffjson: noencoder
type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (r *AccessTokenRequest) Send() (*http.Response, error) {
	values := url.Values{}
	values.Set("client_id", r.ClientID)
	values.Set("client_secret", r.ClientSecret)
	values.Set("grant_type", "authorization_code")
	values.Set("code", r.Code)
	values.Set("redirect_uri", r.RedirectURI)

	return http.PostForm("https://accounts.spotify.com/api/token", values)
}

// ffjson: skip
type RefreshTokenRequest struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
}

// ffjson: noencoder
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
}

func (r RefreshTokenRequest) Send() (*http.Response, error) {
	values := url.Values{}
	values.Set("client_id", r.ClientID)
	values.Set("client_secret", r.ClientSecret)
	values.Set("grant_type", "refresh_token")
	values.Set("refresh_token", r.RefreshToken)

	return http.PostForm("https://accounts.spotify.com/authorize", values)
}
