// Code generated by goa v3.1.3, DO NOT EDIT.
//
// swagger go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/coop-notifiers/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/swagger/swagger.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer) {
	mux.Handle("GET", "/swagger/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../gen/http/openapi.json")
	}))
}
