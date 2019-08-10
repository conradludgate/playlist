package exchanges

import (
	"github.com/conradludgate/playlist/server/exchanges/spotify"
)

//go:generate ffjson $GOFILE

// Playlist is a single playlist type
type Playlist struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Images []spotify.Image
}

// PlaylistResponse is some playlists and the part of a user object
type PlaylistResponse struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Playlists []Playlist `json:"playlists"`
}
