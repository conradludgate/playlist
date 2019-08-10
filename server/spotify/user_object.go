package spotify

// Followers https://developer.spotify.com/documentation/web-api/reference/object-model/#followers-object
type Followers struct {
	HREF  string `json:"href"`
	Total int    `json:"total"`
}

// PrivateUser https://developer.spotify.com/documentation/web-api/reference/object-model/#user-object-private
type PrivateUser struct {
	User
	Country string `json:"country"`
	Email   string `json:"email"`
	Product string `json:"product"`
}

// User https://developer.spotify.com/documentation/web-api/reference/object-model/#user-object-public
type User struct {
	External
	DisplayName string    `json:"display_name"`
	Followers   Followers `json:"followers"`
	Images      []Image   `json:"images"`
}
