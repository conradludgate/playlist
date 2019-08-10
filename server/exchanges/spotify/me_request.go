package spotify

//go:generate ffjson $GOFILE

// MeRequest request for the current user
// Makes a GET request to https://api.spotify.com/v1/me
// https://developer.spotify.com/documentation/web-api/reference/users-profile/get-current-users-profile/
type MeRequest struct {
}

// MeResponse repsponse object for MeRequest
type MeResponse struct {
	PrivateUser
	Error
}

func (MeRequest) request() (method, string, error) {
	return get, "https://api.spotify.com/v1/me", nil
}
