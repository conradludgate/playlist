package spotify

type Album struct {
	SimpleAlbum
	ExternalIDs ExternalID `json:"external_ids"`

	Copyrights []Copyright `json:"copyrights"`
	Genres     []string    `json:"genres"`
	Label      string      `json:"copyrights"`
	Popularity int         `json:"copyrights"`
	Tracks     struct {
		PagingObject
		Items []SimpleTrack `json:"items"`
	} `json:"tracks"`
}

// copyrights	array of copyright objects	The copyright statements of the album.
// external_ids	an external ID object	Known external IDs for the album.
// genres	array of strings	A list of the genres used to classify the album. For example: "Prog Rock" , "Post-Grunge". (If not yet classified, the array is empty.)
// label	string	The label for the album.
// popularity	integer	The popularity of the album. The value will be between 0 and 100, with 100 being the most popular. The popularity is calculated from the popularity of the album’s individual tracks.
// tracks	array of simplified track objects inside a paging object	The tracks of the album.

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

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type Restrictions struct {
	Reason string `json:"reason"`
}
