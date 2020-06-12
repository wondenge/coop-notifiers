// Code generated by goa v3.1.3, DO NOT EDIT.
//
// connect go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/coop-notifiers/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// MountIFTAccountToAccountHandler configures the mux to serve the "connect"
// service "IFTAccountToAccount" endpoint.
func MountIFTAccountToAccountHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/callbacks/coop-bank/FundsTransfer/Internal/A2A/2.0.0", f)
}

// MountPesaLinkSendToAccountHandler configures the mux to serve the "connect"
// service "PesaLinkSendToAccount" endpoint.
func MountPesaLinkSendToAccountHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/callbacks/coop-bank/FundsTransfer/External/A2A/PesaLink/1.0.0", f)
}

// MountPesaLinkSendToPhoneHandler configures the mux to serve the "connect"
// service "PesaLinkSendToPhone" endpoint.
func MountPesaLinkSendToPhoneHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/callbacks/coop-bank/FundsTransfer/External/A2M/PesaLink/1.0.0", f)
}

// MountSendToMPesaHandler configures the mux to serve the "connect" service
// "SendToMPesa" endpoint.
func MountSendToMPesaHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/callbacks/coop-bank/FundsTransfer/External/A2M/Mpesa/v1.0.0", f)
}
