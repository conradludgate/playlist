package spotify

import (
	"io"
	"net/http"
	"net/url"

	fflib "github.com/pquerna/ffjson/fflib/v1"
)

type ffJSONUnmarshaler interface {
	UnmarshalJSON(b []byte) error
}

type ffJSONMarshaler interface {
	MarshalJSON() ([]byte, error)
	MarshalJSONBuf(fflib.EncodingBuffer) error
}

type method uint8

const (
	get method = iota
	post
	put
	patch
	delete
	none
)

//go:generate ffjson $GOFILE

type jsonRequest interface {
	request() (method, string, error)
}

// Send sends off a request, returning the response or an error
func Send(r jsonRequest, accessToken string) (*http.Response, error) {
	method, url, err := r.request()
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	var methodString string
	var body io.Reader

	if method == get {
		methodString = "GET"
	}

	if method == post {
		methodString = "POST"
	}

	if m, ok := r.(ffJSONMarshaler); ok {
		buf := &fflib.Buffer{}
		if err := m.MarshalJSONBuf(buf); err != nil {
			return nil, err
		}
		if buf.Len() > 2 {
			body = buf
		}
	}

	req, err := http.NewRequest(methodString, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	return client.Do(req)
}

// AccessTokenRequest object containing the neccesary information to request an access token for the first time
type AccessTokenRequest struct {
	ClientID     string `json:"-"`
	ClientSecret string `json:"-"`
	Code         string `json:"-"`
	RedirectURI  string `json:"-"`
}

// AccessTokenResponse response for AccessTokenRequest
type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// Send sends off the request for an access token
func (r *AccessTokenRequest) Send() (*http.Response, error) {
	values := url.Values{}
	values.Set("client_id", r.ClientID)
	values.Set("client_secret", r.ClientSecret)
	values.Set("grant_type", "authorization_code")
	values.Set("code", r.Code)
	values.Set("redirect_uri", r.RedirectURI)

	return http.PostForm("https://accounts.spotify.com/api/token", values)
}

// RefreshTokenRequest object containing the neccesary information to request a new access token
type RefreshTokenRequest struct {
	ClientID     string `json:"-"`
	ClientSecret string `json:"-"`
	RefreshToken string `json:"-"`
}

// RefreshTokenResponse response for RefreshTokenRequest
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
}

// Send sends off the request for a new access token
func (r RefreshTokenRequest) Send() (*http.Response, error) {
	values := url.Values{}
	values.Set("client_id", r.ClientID)
	values.Set("client_secret", r.ClientSecret)
	values.Set("grant_type", "refresh_token")
	values.Set("refresh_token", r.RefreshToken)

	return http.PostForm("https://accounts.spotify.com/authorize", values)
}
