package spotify

import (
	"log"
	"net/url"
	"strconv"
)

//go:generate ffjson $GOFILE

type GetPlaylistsRequest struct {
	Limit  int `json:"-"`
	Offset int `json:"-"`
}

func (r GetPlaylistsRequest) Request() (method, string) {
	values := url.Values{}
	limit := r.Limit
	if limit == 0 {
		limit = 20
	}
	values.Set("limit", strconv.Itoa(limit))
	values.Set("offset", strconv.Itoa(r.Offset))
	url, err := url.Parse("https://api.spotify.com/v1/me/playlists")
	if err != nil {
		panic(err)
	}

	url.RawQuery = values.Encode()

	log.Println(url.String())

	return get, url.String()
}

type GetPlaylistsResponse struct {
	PagingObject
	Items []SimplePlaylist `json:"items"`
}
