package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"

	"github.com/conradludgate/playlist/server/spotify"
)

var (
	SPOTIFY_CLIENT_ID     string
	SPOTIFY_CLIENT_SECRET string

	redirectURI string
)

func (state *State) GetAccessToken() (string, error) {
	if time.Now().After(state.Expires) {
		requestData := spotify.RefreshTokenRequest{
			ClientID:     SPOTIFY_CLIENT_ID,
			ClientSecret: SPOTIFY_CLIENT_SECRET,
			RefreshToken: state.RefreshToken,
		}
		resp, err := requestData.Send()
		if err != nil {
			return "", err
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		if resp.StatusCode != http.StatusOK {
			return "", errors.New(string(b))
		}

		data := spotify.RefreshTokenResponse{}
		if err := data.UnmarshalJSON(b); err != nil {
			return "", err
		}

		state.AccessToken = data.AccessToken
		state.Expires = time.Now().Add(time.Duration(data.ExpiresIn) * time.Second)
	}

	return state.AccessToken, nil
}

func init() {
	godotenv.Load()
	redirectURL, err := url.Parse(os.Getenv("HOSTNAME"))
	if err != nil {
		panic(err.Error())
	}

	redirectURL.Path = filepath.Join(redirectURL.Path, "spotify_callback")
	redirectURI = redirectURL.String()
}

func LoginRequest(r *http.Request, state *State) response {
	scopes := []string{"user-read-private"}

	spotify_url, err := url.Parse("https://accounts.spotify.com/authorize")
	if err != nil {
		log.Println(err.Error())
		return response{
			status: http.StatusInternalServerError,
			err:    errors.New("Could not process login request at this time"),
		}
	}

	query := spotify_url.Query()
	query.Set("response_type", "code")
	query.Set("client_id", SPOTIFY_CLIENT_ID)
	for _, scope := range scopes {
		query.Add("scope", scope)
	}

	query.Set("redirect_uri", redirectURI)
	spotify_url.RawQuery = query.Encode()

	headers := make(map[string]string)
	headers["Location"] = spotify_url.String()

	return response{status: http.StatusSeeOther, headers: headers}
}

func SpotifyCallback(r *http.Request, state *State) response {
	requestData := spotify.AccessTokenRequest{
		ClientID:     SPOTIFY_CLIENT_ID,
		ClientSecret: SPOTIFY_CLIENT_SECRET,
		Code:         r.FormValue("code"),
		RedirectURI:  redirectURI,
	}
	resp, err := requestData.Send()
	if err != nil {
		return response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	if resp.StatusCode != http.StatusOK {
		return response{
			status: http.StatusBadRequest,
			err:    errors.New(string(b)),
		}
	}

	data := spotify.AccessTokenResponse{}
	if err := data.UnmarshalJSON(b); err != nil {
		return response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	state.AccessToken = data.AccessToken
	state.RefreshToken = data.RefreshToken
	state.Expires = time.Now().Add(time.Duration(data.ExpiresIn) * time.Second)

	return response{status: 200}
}

func Me(r *http.Request, state *State) response {
	accessToken, err := state.GetAccessToken()
	if err != nil {
		return response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	requestData := spotify.MeRequest{AccessToken: accessToken}
	resp, err := requestData.Send()
	if err != nil {
		return response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	if resp.StatusCode != http.StatusOK {
		return response{
			status: http.StatusBadRequest,
			err:    errors.New(string(b)),
		}
	}

	data := spotify.MeResponse{}
	if err := data.UnmarshalJSON(b); err != nil {
		return response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	fmt.Println(data)

	return response{status: 200}
}
