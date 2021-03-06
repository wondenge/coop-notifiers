// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the connect service.
//
// Command:
// $ goa gen github.com/wondenge/coop-notifiers/design

package client

// IFTAccountToAccountConnectPath returns the URL path to the connect service IFTAccountToAccount HTTP endpoint.
func IFTAccountToAccountConnectPath() string {
	return "/callbacks/coop-bank/FundsTransfer/Internal/A2A/2.0.0"
}

// PesaLinkSendToAccountConnectPath returns the URL path to the connect service PesaLinkSendToAccount HTTP endpoint.
func PesaLinkSendToAccountConnectPath() string {
	return "/callbacks/coop-bank/FundsTransfer/External/A2A/PesaLink/1.0.0"
}

// PesaLinkSendToPhoneConnectPath returns the URL path to the connect service PesaLinkSendToPhone HTTP endpoint.
func PesaLinkSendToPhoneConnectPath() string {
	return "/callbacks/coop-bank/FundsTransfer/External/A2M/PesaLink/1.0.0"
}

// SendToMPesaConnectPath returns the URL path to the connect service SendToMPesa HTTP endpoint.
func SendToMPesaConnectPath() string {
	return "/callbacks/coop-bank/FundsTransfer/External/A2M/Mpesa/v1.0.0"
}
