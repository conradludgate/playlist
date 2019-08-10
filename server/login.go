package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/conradludgate/playlist/server/exchanges"

	"github.com/joho/godotenv"

	"github.com/conradludgate/playlist/server/spotify"
)

var (
	spotifyClientID     string
	spotifyClientSecret string

	redirectURI string
)

// GetAccessToken returns the access token if it's currently valid, or requests a new one
func (state *State) GetAccessToken() (string, error) {
	if time.Now().After(state.Expires) {
		requestData := spotify.RefreshTokenRequest{
			ClientID:     spotifyClientID,
			ClientSecret: spotifyClientSecret,
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

// LoginRequest is the handler for logging in and getting a spotify auth code
func LoginRequest(r *http.Request, state *State) Response {
	scopes := []string{
		"user-read-private",
		"playlist-read-private",
	}

	spotifyURL, err := url.Parse("https://accounts.spotify.com/authorize")
	if err != nil {
		log.Println(err.Error())
		return Response{
			status: http.StatusInternalServerError,
			err:    errors.New("Could not process login request at this time"),
		}
	}

	query := spotifyURL.Query()
	query.Set("response_type", "code")
	query.Set("client_id", spotifyClientID)
	for _, scope := range scopes {
		query.Add("scope", scope)
	}

	query.Set("redirect_uri", redirectURI)
	spotifyURL.RawQuery = query.Encode()

	return redirect(http.StatusSeeOther, spotifyURL.String())
}

// SpotifyCallback is the handler for when spotify redirects back to us after login
func SpotifyCallback(r *http.Request, state *State) Response {
	requestData := spotify.AccessTokenRequest{
		ClientID:     spotifyClientID,
		ClientSecret: spotifyClientSecret,
		Code:         r.FormValue("code"),
		RedirectURI:  redirectURI,
	}
	resp, err := requestData.Send()
	if err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	if resp.StatusCode != http.StatusOK {
		return Response{
			status: http.StatusBadRequest,
			err:    errors.New(string(b)),
		}
	}

	data := spotify.AccessTokenResponse{}
	if err := data.UnmarshalJSON(b); err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	state.AccessToken = data.AccessToken
	state.RefreshToken = data.RefreshToken
	state.Expires = time.Now().Add(time.Duration(data.ExpiresIn) * time.Second)

	return redirect(http.StatusSeeOther, "/me")
}

// Me is the handler for testing
func Me(r *http.Request, state *State) Response {
	accessToken, err := state.GetAccessToken()
	if err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	// Get Me
	resp, err := spotify.Send(spotify.MeRequest{}, accessToken)
	if err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	if resp.StatusCode != http.StatusOK {
		return Response{
			status: http.StatusBadRequest,
			err:    errors.New(string(b)),
		}
	}

	me := spotify.MeResponse{}
	if err := me.UnmarshalJSON(b); err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	if me.IsError() {
		return Response{
			status: me.Status,
			err:    me.Error,
		}
	}

	log.Println(me.PrivateUser)

	// Get my playlists
	resp, err = spotify.Send(spotify.GetPlaylistsRequest{}, accessToken)
	if err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	if resp.StatusCode != http.StatusOK {
		return Response{
			status: http.StatusBadRequest,
			err:    errors.New(string(b)),
		}
	}

	playlists := spotify.GetPlaylistsResponse{}
	if err := playlists.UnmarshalJSON(b); err != nil {
		return Response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	plr := exchanges.PlaylistResponse{
		ID:        me.ID,
		Playlists: make([]exchanges.Playlist, 0, len(playlists.Items)),
	}

	for _, playlist := range playlists.Items {
		plr.Playlists = append(plr.Playlists, exchanges.Playlist{
			ID:     playlist.ID,
			Name:   playlist.Name,
			Images: playlist.Images,
		})
	}

	return Response{
		status: 200,
		value:  &plr,
	}
}
