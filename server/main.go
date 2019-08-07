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
	r.HandleFunc("/login", HttpMiddleware(LoginRequest, &state))
	r.HandleFunc("/spotify_callback", HttpMiddleware(SpotifyCallback, &state))
	r.HandleFunc("/me", HttpMiddleware(Me, &state))

	http.ListenAndServe(":8080", r)
}

func init() {
	godotenv.Load()
	SPOTIFY_CLIENT_ID = os.Getenv("SPOTIFY_CLIENT_ID")
	SPOTIFY_CLIENT_SECRET = os.Getenv("SPOTIFY_CLIENT_SECRET")
	// HOSTNAME = os.Getenv("HOSTNAME")
}
