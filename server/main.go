package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	state := State{}

	r := mux.NewRouter()
	r.HandleFunc("/login", HTTPMiddleware(LoginRequest, &state))
	r.HandleFunc("/spotify_callback", HTTPMiddleware(SpotifyCallback, &state))
	r.HandleFunc("/me", HTTPMiddleware(Me, &state))

	http.ListenAndServe(":8080", r)
}

func init() {
	godotenv.Load()
	spotifyClientID = os.Getenv("SPOTIFY_CLIENT_ID")
	spotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	// HOSTNAME = os.Getenv("HOSTNAME")
}
