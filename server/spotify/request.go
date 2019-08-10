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
)

//go:generate ffjson $GOFILE

type JSONRequest interface {
	Request() (method, string)
}

func Send(r JSONRequest, accessToken string) (*http.Response, error) {
	client := http.Client{}
	method, url := r.Request()
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

type AccessTokenRequest struct {
	ClientID     string `json:"-"`
	ClientSecret string `json:"-"`
	Code         string `json:"-"`
	RedirectURI  string `json:"-"`
}

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

type RefreshTokenRequest struct {
	ClientID     string `json:"-"`
	ClientSecret string `json:"-"`
	RefreshToken string `json:"-"`
}

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
