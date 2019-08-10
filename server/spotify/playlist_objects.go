package spotify

import "time"

type Playlist struct {
	SimplePlaylist
	Followers   Followers `json:"followers"`
	Description string    `json:"description"`
}

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

type PlaylistTrack struct {
	AddedAt time.Time `json:"added_at`
	AddedBy User      `json:"added_by"`
	IsLocal bool      `json:"is_local"`
	Track   Track     `json:"track"`
}
