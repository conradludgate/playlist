package spotify

import (
	"net/url"
	"path"
	"strconv"
)

//go:generate ffjson $GOFILE

// GetPlaylistsRequest https://developer.spotify.com/console/get-current-user-playlists/
//
// If no UserID is provided, it gets the current user's playlists
// https://developer.spotify.com/console/get-playlists/
type GetPlaylistsRequest struct {
	Limit  int    `json:"-"`
	Offset int    `json:"-"`
	UserID string `json:"-"`
}

func (r GetPlaylistsRequest) request() (method, string, error) {
	values := url.Values{}
	limit := r.Limit
	if limit == 0 {
		limit = 20
	}
	values.Set("limit", strconv.Itoa(limit))
	values.Set("offset", strconv.Itoa(r.Offset))

	var err error
	var u *url.URL

	if len(r.UserID) == 0 {
		u, err = url.Parse("https://api.spotify.com/v1/me/playlists")
	} else {
		u, err = url.Parse(path.Join("https://api.spotify.com/v1/users", r.UserID, "playlists"))
	}

	if err != nil {
		return none, "", err
	}

	u.RawQuery = values.Encode()

	return get, u.String(), nil
}

// GetPlaylistsResponse is the response for GetPlaylistsRequest
type GetPlaylistsResponse struct {
	PagingObject
	Items []SimplePlaylist `json:"items"`
}
