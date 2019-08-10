package spotify

// Artist https://developer.spotify.com/documentation/web-api/reference/object-model/#artist-object-full
type Artist struct {
	SimpleArtist

	Followers  Followers `json:"followers"`
	Genres     []string  `json:"genres"`
	Images     []Image   `json:"images"`
	Popularity int       `json:"popularity"`
}

// SimpleArtist https://developer.spotify.com/documentation/web-api/reference/object-model/#artist-object-simplified
type SimpleArtist struct {
	External

	Name string `json:"name"`
}
