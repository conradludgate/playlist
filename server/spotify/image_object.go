package spotify

// Image https://developer.spotify.com/documentation/web-api/reference/object-model/#image-object
type Image struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}
