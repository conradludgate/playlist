package spotify

// PagingObject https://developer.spotify.com/documentation/web-api/reference/object-model/#paging-object
type PagingObject struct {
	HREF     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}
