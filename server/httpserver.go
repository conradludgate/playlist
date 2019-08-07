package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/conradludgate/playlist/server/exchanges"
)

var (
	SPOTIFY_CLIENT_ID string
	HOSTNAME          string
)

type State struct {
}

type response struct {
	value   exchanges.FFJSONMarshaler
	status  int
	err     error
	headers map[string]string
}

type HttpHandler func(r *http.Request, state State) response

func HttpMiddleware(handler HttpHandler, state State) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		resp := handler(r, state)

		// Add or set headers?
		for key, value := range resp.headers {
			w.Header().Set(key, value)
		}

		var response exchanges.Response
		response.Status = resp.status
		w.WriteHeader(resp.status)

		if resp.err != nil {

			response.Error = resp.err
			response.Success = false

			b, err := response.MarshalJSON()
			if err != nil {
				// Should not error on marshalling
				// But I guess I never know?
				// Should replace with a 'failsafe' solution in future
				panic(err)
			}

			if _, err := w.Write(b); err != nil {
				log.Println(err)
			}

			return
		}

		response.Success = true
		response.Value = resp.value

		b, err := response.MarshalJSON()
		if err != nil {
			// Should not error on marshalling
			// But I guess I never know?
			// Should replace with a 'failsafe' solution in future
			panic(err)
		}

		if _, err := w.Write(b); err != nil {
			log.Println(err)
		}
	}
}

func LoginRequest(r *http.Request, state State) response {
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

	redirect_url, err := url.Parse(HOSTNAME)
	if err != nil {
		log.Println(err.Error())
		return response{
			status: http.StatusInternalServerError,
			err:    errors.New("Could not process login request at this time"),
		}
	}

	redirect_url.Path = filepath.Join(redirect_url.Path, "spotify_callback")

	query.Set("redirect_uri", redirect_url.String())
	spotify_url.RawQuery = query.Encode()

	headers := make(map[string]string)
	headers["Location"] = spotify_url.String()

	return response{status: http.StatusSeeOther, headers: headers}
}

func SpotifyCallback(r *http.Request, state State) response {
	log.Println(r.URL)
	return response{status: 200}
}
