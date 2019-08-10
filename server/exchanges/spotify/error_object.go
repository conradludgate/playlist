package spotify

import (
	"fmt"
	"net/http"
)

// Error Standard Spotify error type
// https://developer.spotify.com/documentation/web-api/#response-schema
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// IsError Easy test to check if error was provided
func (e Error) IsError() bool {
	// https://developer.spotify.com/documentation/web-api/#response-status-codes
	return e.Status != 0 && e.Status >= 300
}

func (e Error) Error() string {
	return fmt.Sprintf("spotify: [%d] %s: %s", e.Status, http.StatusText(e.Status), e.Message)
}
