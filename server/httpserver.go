package main

import (
	"net/http"
	"time"

	"github.com/conradludgate/playlist/server/exchanges"
)

// State is the state that all HTTP Handlers might need to access
type State struct {
	AccessToken  string
	Expires      time.Time
	RefreshToken string
}

// Response is a response from all HTTPHandlers to get processed by HTTPMiddleware
type Response struct {
	value   exchanges.FFJSONMarshaler
	status  int
	err     error
	headers map[string]string
}

// HTTPHandler handler for all http requests
type HTTPHandler func(r *http.Request, state *State) Response

// HTTPMiddleware turns a HTTPHandler and State into a valid http.HandlerFunc
func HTTPMiddleware(handler HTTPHandler, state *State) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		resp := handler(r, state)

		// Add or set headers?
		for key, value := range resp.headers {
			w.Header().Set(key, value)
		}

		var response exchanges.Response
		response.Status = resp.status

		if resp.err != nil {
			response.Error = resp.err
			response.Success = false
		} else {
			response.Success = true
			response.Value = resp.value
		}

		b, err := response.MarshalJSON()
		if err != nil {
			// Should not error on marshalling
			// But I guess I never know?
			// Should replace with a 'failsafe' solution in future
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.status)
		if _, err := w.Write(b); err != nil {
			panic(err)
		}
	}
}
