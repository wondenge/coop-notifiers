// Code generated by goa v3.1.3, DO NOT EDIT.
//
// connect go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/coop-notifiers/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/coop-notifiers/gen/http/connect/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeIFTAccountToAccountResponse returns a go-kit EncodeResponseFunc
// suitable for encoding connect IFTAccountToAccount responses.
func EncodeIFTAccountToAccountResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeIFTAccountToAccountResponse(encoder)
}

// DecodeIFTAccountToAccountRequest returns a go-kit DecodeRequestFunc suitable
// for decoding connect IFTAccountToAccount requests.
func DecodeIFTAccountToAccountRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeIFTAccountToAccountRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeIFTAccountToAccountError returns a go-kit EncodeResponseFunc suitable
// for encoding errors returned by the connect IFTAccountToAccount endpoint.
func EncodeIFTAccountToAccountError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeIFTAccountToAccountError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodePesaLinkSendToAccountResponse returns a go-kit EncodeResponseFunc
// suitable for encoding connect PesaLinkSendToAccount responses.
func EncodePesaLinkSendToAccountResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodePesaLinkSendToAccountResponse(encoder)
}

// DecodePesaLinkSendToAccountRequest returns a go-kit DecodeRequestFunc
// suitable for decoding connect PesaLinkSendToAccount requests.
func DecodePesaLinkSendToAccountRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodePesaLinkSendToAccountRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodePesaLinkSendToAccountError returns a go-kit EncodeResponseFunc
// suitable for encoding errors returned by the connect PesaLinkSendToAccount
// endpoint.
func EncodePesaLinkSendToAccountError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodePesaLinkSendToAccountError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodePesaLinkSendToPhoneResponse returns a go-kit EncodeResponseFunc
// suitable for encoding connect PesaLinkSendToPhone responses.
func EncodePesaLinkSendToPhoneResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodePesaLinkSendToPhoneResponse(encoder)
}

// DecodePesaLinkSendToPhoneRequest returns a go-kit DecodeRequestFunc suitable
// for decoding connect PesaLinkSendToPhone requests.
func DecodePesaLinkSendToPhoneRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodePesaLinkSendToPhoneRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodePesaLinkSendToPhoneError returns a go-kit EncodeResponseFunc suitable
// for encoding errors returned by the connect PesaLinkSendToPhone endpoint.
func EncodePesaLinkSendToPhoneError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodePesaLinkSendToPhoneError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeSendToMPesaResponse returns a go-kit EncodeResponseFunc suitable for
// encoding connect SendToMPesa responses.
func EncodeSendToMPesaResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeSendToMPesaResponse(encoder)
}

// DecodeSendToMPesaRequest returns a go-kit DecodeRequestFunc suitable for
// decoding connect SendToMPesa requests.
func DecodeSendToMPesaRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeSendToMPesaRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeSendToMPesaError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the connect SendToMPesa endpoint.
func EncodeSendToMPesaError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeSendToMPesaError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}