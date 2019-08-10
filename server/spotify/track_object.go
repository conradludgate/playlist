package spotify

type Track struct {
	SimpleTrack
	ExternalIDs ExternalID `json:"external_ids"`

	Album      SimpleAlbum `json:"album"`
	Popularity int         `json:"popularity"`
}

type SimpleTrack struct {
	External

	Artists      []SimpleArtist `json:"artists"`
	Markets      []string       `json:"available_markets"`
	DiscNumber   int            `json:"disc_number"`
	DurationMS   int            `json:"duration_ms"`
	Explicit     bool           `json:"explicit"`
	IsPlayable   bool           `json:"is_playable"`
	LinkedFrom   External       `json:"linked_from"`
	Restrictions Restrictions   `json:"restrictions"`
	Name         string         `json:"name"`
	PreviewURL   string         `json:"preview_url"`
	TrackNumber  int            `json:"track_number"`
	IsLocal      bool           `json:"is_local"`
}
