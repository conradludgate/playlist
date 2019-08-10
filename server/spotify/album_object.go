package spotify

// Album https://developer.spotify.com/documentation/web-api/reference/object-model/#album-object-full
type Album struct {
	SimpleAlbum
	ExternalIDs ExternalID `json:"external_ids"`

	Copyrights []Copyright `json:"copyrights"`
	Genres     []string    `json:"genres"`
	Label      string      `json:"label"`
	Popularity int         `json:"popularity"`
	Tracks     struct {
		PagingObject
		Items []SimpleTrack `json:"items"`
	} `json:"tracks"`
}

// SimpleAlbum https://developer.spotify.com/documentation/web-api/reference/object-model/#album-object-simplified
type SimpleAlbum struct {
	External

	AlbumGroup string         `json:"album_group"`
	AlbumType  string         `json:"album_type"`
	Artists    []SimpleArtist `json:"artists"`
	Markets    []string       `json:"available_markets"`
	Images     []Image        `json:"images"`
	Name       string         `json:"name"`

	ReleastDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`

	Restrictions Restrictions `json:"rescriptions"`
}

// Copyright https://developer.spotify.com/documentation/web-api/reference/object-model/#copyright-object
type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// Restrictions reason why is_playable would be false
type Restrictions struct {
	Reason string `json:"reason"`
}
