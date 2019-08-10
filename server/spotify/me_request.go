package spotify

//go:generate ffjson $GOFILE

// MeRequest
//
// Request for the current user
// Makes a GET request to https://api.spotify.com/v1/me
// https://developer.spotify.com/documentation/web-api/reference/users-profile/get-current-users-profile/
type MeRequest struct {
}

// MeResponse
//
// Response object for [/#MeRequest](MeRequest)
type MeResponse struct {
	PrivateUser
	Error
}

// Request conforms to Request interface
func (MeRequest) Request() (method, string) {
	return get, "https://api.spotify.com/v1/me"
}
