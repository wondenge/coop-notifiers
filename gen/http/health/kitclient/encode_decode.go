// Code generated by goa v3.1.3, DO NOT EDIT.
//
// health go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/coop-notifiers/design

package client

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/coop-notifiers/gen/http/health/client"
	goahttp "goa.design/goa/v3/http"
)

// DecodeShowResponse returns a go-kit DecodeResponseFunc suitable for decoding
// health show responses.
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeShowResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
