package spotify

type Artist struct {
	SimpleArtist

	Followers  Followers `json:"followers"`
	Genres     []string  `json:"followers"`
	Images     []Image   `json:"followers"`
	Popularity int       `json:"followers"`
}

type SimpleArtist struct {
	External

	Name string `json:"name"`
}
