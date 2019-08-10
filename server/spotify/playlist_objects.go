package spotify

import "time"

// Playlist https://developer.spotify.com/documentation/web-api/reference/object-model/#playlist-object-full
type Playlist struct {
	SimplePlaylist
	Followers   Followers `json:"followers"`
	Description string    `json:"description"`
}

// SimplePlaylist https://developer.spotify.com/documentation/web-api/reference/object-model/#playlist-object-simplified
type SimplePlaylist struct {
	External

	Collaborative bool    `json:"collaborative"`
	Images        []Image `json:"images"`
	Name          string  `json:"name"`
	Owner         User    `json:"owner"`
	Public        bool    `json:"public"`
	SnapshotID    string  `json:"snapshot_id"`
	Tracks        struct {
		PagingObject
		Items []SimpleTrack `json:"items"`
	} `json:"tracks"`
}

// PlaylistTrack https://developer.spotify.com/documentation/web-api/reference/object-model/#playlist-track-object
type PlaylistTrack struct {
	AddedAt time.Time `json:"added_at"`
	AddedBy User      `json:"added_by"`
	IsLocal bool      `json:"is_local"`
	Track   Track     `json:"track"`
}
