// Code generated by goa v3.1.3, DO NOT EDIT.
//
// connect HTTP client types
//
// Command:
// $ goa gen github.com/wondenge/coop-notifiers/design

package client

import (
	"unicode/utf8"

	connect "github.com/wondenge/coop-notifiers/gen/connect"
	goa "goa.design/goa/v3/pkg"
)

// IFTAccountToAccountRequestBody is the type of the "connect" service
// "IFTAccountToAccount" endpoint HTTP request body.
type IFTAccountToAccountRequestBody struct {
	// Your unique transaction request message identifier
	MessageReference string `form:"MessageReference" json:"MessageReference" xml:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime string `form:"MessageDateTime" json:"MessageDateTime" xml:"MessageDateTime"`
	// Message Response Code
	MessageCode string `form:"MessageCode" json:"MessageCode" xml:"MessageCode"`
	// Message Code description
	MessageDescription string                                        `form:"MessageDescription" json:"MessageDescription" xml:"MessageDescription"`
	Source             *SourceAccountCallbackRequestRequestBody      `form:"Source" json:"Source" xml:"Source"`
	Destinations       *DestinationAccountCallbackRequestRequestBody `form:"Destinations" json:"Destinations" xml:"Destinations"`
}

// PesaLinkSendToAccountRequestBody is the type of the "connect" service
// "PesaLinkSendToAccount" endpoint HTTP request body.
type PesaLinkSendToAccountRequestBody struct {
	// Your unique transaction request message identifier
	MessageReference string `form:"MessageReference" json:"MessageReference" xml:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime string `form:"MessageDateTime" json:"MessageDateTime" xml:"MessageDateTime"`
	// Message Response Code
	MessageCode string `form:"MessageCode" json:"MessageCode" xml:"MessageCode"`
	// Message Code description
	MessageDescription string `form:"MessageDescription" json:"MessageDescription" xml:"MessageDescription"`
	// Your callback URL that will receive transaction processing results
	CallBackURL  *string                                       `form:"CallBackUrl,omitempty" json:"CallBackUrl,omitempty" xml:"CallBackUrl,omitempty"`
	Source       *SourceAccountCallbackRequestRequestBody      `form:"Source" json:"Source" xml:"Source"`
	Destinations *DestinationAccountCallbackRequestRequestBody `form:"Destinations" json:"Destinations" xml:"Destinations"`
}

// PesaLinkSendToPhoneRequestBody is the type of the "connect" service
// "PesaLinkSendToPhone" endpoint HTTP request body.
type PesaLinkSendToPhoneRequestBody struct {
	// Your unique transaction request message identifier
	MessageReference string `form:"MessageReference" json:"MessageReference" xml:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime string `form:"MessageDateTime" json:"MessageDateTime" xml:"MessageDateTime"`
	// Message Response Code
	MessageCode string `form:"MessageCode" json:"MessageCode" xml:"MessageCode"`
	// Message Code description
	MessageDescription string                                          `form:"MessageDescription" json:"MessageDescription" xml:"MessageDescription"`
	Source             *SourceAccountCallbackRequestRequestBody        `form:"Source" json:"Source" xml:"Source"`
	Destinations       []*DestinationAccountCallbackRequestRequestBody `form:"Destinations" json:"Destinations" xml:"Destinations"`
}

// SendToMPesaRequestBody is the type of the "connect" service "SendToMPesa"
// endpoint HTTP request body.
type SendToMPesaRequestBody struct {
	// Your unique transaction request message identifier
	MessageReference string `form:"MessageReference" json:"MessageReference" xml:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime string `form:"MessageDateTime" json:"MessageDateTime" xml:"MessageDateTime"`
	// Message Response Code
	MessageCode string `form:"MessageCode" json:"MessageCode" xml:"MessageCode"`
	// Message Code description
	MessageDescription string `form:"MessageDescription" json:"MessageDescription" xml:"MessageDescription"`
	// Your callback URL that will receive transaction processing results
	CallBackURL  *string                                       `form:"CallBackUrl,omitempty" json:"CallBackUrl,omitempty" xml:"CallBackUrl,omitempty"`
	Source       *SourceAccountCallbackRequestRequestBody      `form:"Source" json:"Source" xml:"Source"`
	Destinations *DestinationAccountCallbackRequestRequestBody `form:"Destinations" json:"Destinations" xml:"Destinations"`
}

// IFTAccountToAccountPartialSuccessResponseBody is the type of the "connect"
// service "IFTAccountToAccount" endpoint HTTP response body for the
// "partial_success" error.
type IFTAccountToAccountPartialSuccessResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// IFTAccountToAccountFullFailureResponseBody is the type of the "connect"
// service "IFTAccountToAccount" endpoint HTTP response body for the
// "full_failure" error.
type IFTAccountToAccountFullFailureResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// PesaLinkSendToAccountPartialSuccessResponseBody is the type of the "connect"
// service "PesaLinkSendToAccount" endpoint HTTP response body for the
// "partial_success" error.
type PesaLinkSendToAccountPartialSuccessResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// PesaLinkSendToAccountFullFailureResponseBody is the type of the "connect"
// service "PesaLinkSendToAccount" endpoint HTTP response body for the
// "full_failure" error.
type PesaLinkSendToAccountFullFailureResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// PesaLinkSendToPhonePartialSuccessResponseBody is the type of the "connect"
// service "PesaLinkSendToPhone" endpoint HTTP response body for the
// "partial_success" error.
type PesaLinkSendToPhonePartialSuccessResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// PesaLinkSendToPhoneFullFailureResponseBody is the type of the "connect"
// service "PesaLinkSendToPhone" endpoint HTTP response body for the
// "full_failure" error.
type PesaLinkSendToPhoneFullFailureResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SendToMPesaPartialSuccessResponseBody is the type of the "connect" service
// "SendToMPesa" endpoint HTTP response body for the "partial_success" error.
type SendToMPesaPartialSuccessResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SendToMPesaFullFailureResponseBody is the type of the "connect" service
// "SendToMPesa" endpoint HTTP response body for the "full_failure" error.
type SendToMPesaFullFailureResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SourceAccountCallbackRequestRequestBody is used to define fields on request
// body types.
type SourceAccountCallbackRequestRequestBody struct {
	// Posting account number
	AccountNumber string `form:"AccountNumber" json:"AccountNumber" xml:"AccountNumber"`
	// Transaction Amount
	Amount uint64 `form:"Amount" json:"Amount" xml:"Amount"`
	// Posting account currency in ISO Currency Code
	TransactionCurrency string `form:"TransactionCurrency" json:"TransactionCurrency" xml:"TransactionCurrency"`
	// Posting account transaction narration
	Narration string `form:"Narration" json:"Narration" xml:"Narration"`
	// Posting leg response code
	ResponseCode string `form:"ResponseCode" json:"ResponseCode" xml:"ResponseCode"`
	// Posting leg response description
	ResponseDescription string `form:"ResponseDescription" json:"ResponseDescription" xml:"ResponseDescription"`
}

// DestinationAccountCallbackRequestRequestBody is used to define fields on
// request body types.
type DestinationAccountCallbackRequestRequestBody struct {
	// Unique posting reference for the transaction leg
	ReferenceNumber string `form:"ReferenceNumber" json:"ReferenceNumber" xml:"ReferenceNumber"`
	// Recipient M-Pesa mobile number
	MobileNumber string `form:"MobileNumber" json:"MobileNumber" xml:"MobileNumber"`
	// Transaction Amount
	Amount int `form:"Amount" json:"Amount" xml:"Amount"`
	// Transaction posting narration
	Narration string `form:"Narration" json:"Narration" xml:"Narration"`
	// Co-operative Bank's processed transaction number concatenated with M-Pesa
	// transaction number
	TransactionID *string `form:"TransactionID,omitempty" json:"TransactionID,omitempty" xml:"TransactionID,omitempty"`
	// Posting leg response code
	ResponseCode *string `form:"ResponseCode,omitempty" json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// Posting leg response description
	ResponseDescription *string `form:"ResponseDescription,omitempty" json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
}

// NewIFTAccountToAccountRequestBody builds the HTTP request body from the
// payload of the "IFTAccountToAccount" endpoint of the "connect" service.
func NewIFTAccountToAccountRequestBody(p *connect.IFTAccountToAccountCallbackRequest) *IFTAccountToAccountRequestBody {
	body := &IFTAccountToAccountRequestBody{
		MessageReference:   p.MessageReference,
		MessageDateTime:    p.MessageDateTime,
		MessageCode:        p.MessageCode,
		MessageDescription: p.MessageDescription,
	}
	if p.Source != nil {
		body.Source = marshalConnectSourceAccountCallbackRequestToSourceAccountCallbackRequestRequestBody(p.Source)
	}
	if p.Destinations != nil {
		body.Destinations = marshalConnectDestinationAccountCallbackRequestToDestinationAccountCallbackRequestRequestBody(p.Destinations)
	}
	return body
}

// NewPesaLinkSendToAccountRequestBody builds the HTTP request body from the
// payload of the "PesaLinkSendToAccount" endpoint of the "connect" service.
func NewPesaLinkSendToAccountRequestBody(p *connect.PesaLinkSendToAccountCallbackRequest) *PesaLinkSendToAccountRequestBody {
	body := &PesaLinkSendToAccountRequestBody{
		MessageReference:   p.MessageReference,
		MessageDateTime:    p.MessageDateTime,
		MessageCode:        p.MessageCode,
		MessageDescription: p.MessageDescription,
		CallBackURL:        p.CallBackURL,
	}
	if p.Source != nil {
		body.Source = marshalConnectSourceAccountCallbackRequestToSourceAccountCallbackRequestRequestBody(p.Source)
	}
	if p.Destinations != nil {
		body.Destinations = marshalConnectDestinationAccountCallbackRequestToDestinationAccountCallbackRequestRequestBody(p.Destinations)
	}
	return body
}

// NewPesaLinkSendToPhoneRequestBody builds the HTTP request body from the
// payload of the "PesaLinkSendToPhone" endpoint of the "connect" service.
func NewPesaLinkSendToPhoneRequestBody(p *connect.PesaLinkSendToPhoneCallbackRequest) *PesaLinkSendToPhoneRequestBody {
	body := &PesaLinkSendToPhoneRequestBody{
		MessageReference:   p.MessageReference,
		MessageDateTime:    p.MessageDateTime,
		MessageCode:        p.MessageCode,
		MessageDescription: p.MessageDescription,
	}
	if p.Source != nil {
		body.Source = marshalConnectSourceAccountCallbackRequestToSourceAccountCallbackRequestRequestBody(p.Source)
	}
	if p.Destinations != nil {
		body.Destinations = make([]*DestinationAccountCallbackRequestRequestBody, len(p.Destinations))
		for i, val := range p.Destinations {
			body.Destinations[i] = marshalConnectDestinationAccountCallbackRequestToDestinationAccountCallbackRequestRequestBody(val)
		}
	}
	return body
}

// NewSendToMPesaRequestBody builds the HTTP request body from the payload of
// the "SendToMPesa" endpoint of the "connect" service.
func NewSendToMPesaRequestBody(p *connect.SendToMpesaCallbackRequest) *SendToMPesaRequestBody {
	body := &SendToMPesaRequestBody{
		MessageReference:   p.MessageReference,
		MessageDateTime:    p.MessageDateTime,
		MessageCode:        p.MessageCode,
		MessageDescription: p.MessageDescription,
		CallBackURL:        p.CallBackURL,
	}
	if p.Source != nil {
		body.Source = marshalConnectSourceAccountCallbackRequestToSourceAccountCallbackRequestRequestBody(p.Source)
	}
	if p.Destinations != nil {
		body.Destinations = marshalConnectDestinationAccountCallbackRequestToDestinationAccountCallbackRequestRequestBody(p.Destinations)
	}
	return body
}

// NewIFTAccountToAccountPartialSuccess builds a connect service
// IFTAccountToAccount endpoint partial_success error.
func NewIFTAccountToAccountPartialSuccess(body *IFTAccountToAccountPartialSuccessResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewIFTAccountToAccountFullFailure builds a connect service
// IFTAccountToAccount endpoint full_failure error.
func NewIFTAccountToAccountFullFailure(body *IFTAccountToAccountFullFailureResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewPesaLinkSendToAccountPartialSuccess builds a connect service
// PesaLinkSendToAccount endpoint partial_success error.
func NewPesaLinkSendToAccountPartialSuccess(body *PesaLinkSendToAccountPartialSuccessResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewPesaLinkSendToAccountFullFailure builds a connect service
// PesaLinkSendToAccount endpoint full_failure error.
func NewPesaLinkSendToAccountFullFailure(body *PesaLinkSendToAccountFullFailureResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewPesaLinkSendToPhonePartialSuccess builds a connect service
// PesaLinkSendToPhone endpoint partial_success error.
func NewPesaLinkSendToPhonePartialSuccess(body *PesaLinkSendToPhonePartialSuccessResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewPesaLinkSendToPhoneFullFailure builds a connect service
// PesaLinkSendToPhone endpoint full_failure error.
func NewPesaLinkSendToPhoneFullFailure(body *PesaLinkSendToPhoneFullFailureResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewSendToMPesaPartialSuccess builds a connect service SendToMPesa endpoint
// partial_success error.
func NewSendToMPesaPartialSuccess(body *SendToMPesaPartialSuccessResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewSendToMPesaFullFailure builds a connect service SendToMPesa endpoint
// full_failure error.
func NewSendToMPesaFullFailure(body *SendToMPesaFullFailureResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateIFTAccountToAccountPartialSuccessResponseBody runs the validations
// defined on IFTAccountToAccount_partial_success_Response_Body
func ValidateIFTAccountToAccountPartialSuccessResponseBody(body *IFTAccountToAccountPartialSuccessResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateIFTAccountToAccountFullFailureResponseBody runs the validations
// defined on IFTAccountToAccount_full_failure_Response_Body
func ValidateIFTAccountToAccountFullFailureResponseBody(body *IFTAccountToAccountFullFailureResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidatePesaLinkSendToAccountPartialSuccessResponseBody runs the validations
// defined on PesaLinkSendToAccount_partial_success_Response_Body
func ValidatePesaLinkSendToAccountPartialSuccessResponseBody(body *PesaLinkSendToAccountPartialSuccessResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidatePesaLinkSendToAccountFullFailureResponseBody runs the validations
// defined on PesaLinkSendToAccount_full_failure_Response_Body
func ValidatePesaLinkSendToAccountFullFailureResponseBody(body *PesaLinkSendToAccountFullFailureResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidatePesaLinkSendToPhonePartialSuccessResponseBody runs the validations
// defined on PesaLinkSendToPhone_partial_success_Response_Body
func ValidatePesaLinkSendToPhonePartialSuccessResponseBody(body *PesaLinkSendToPhonePartialSuccessResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidatePesaLinkSendToPhoneFullFailureResponseBody runs the validations
// defined on PesaLinkSendToPhone_full_failure_Response_Body
func ValidatePesaLinkSendToPhoneFullFailureResponseBody(body *PesaLinkSendToPhoneFullFailureResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSendToMPesaPartialSuccessResponseBody runs the validations defined
// on SendToMPesa_partial_success_Response_Body
func ValidateSendToMPesaPartialSuccessResponseBody(body *SendToMPesaPartialSuccessResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSendToMPesaFullFailureResponseBody runs the validations defined on
// SendToMPesa_full_failure_Response_Body
func ValidateSendToMPesaFullFailureResponseBody(body *SendToMPesaFullFailureResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSourceAccountCallbackRequestRequestBody runs the validations defined
// on SourceAccountCallbackRequestRequestBody
func ValidateSourceAccountCallbackRequestRequestBody(body *SourceAccountCallbackRequestRequestBody) (err error) {
	if utf8.RuneCountInString(body.AccountNumber) < 14 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, true))
	}
	if utf8.RuneCountInString(body.AccountNumber) > 14 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, false))
	}
	if body.Amount < 0.01 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("body.Amount", body.Amount, 0.01, true))
	}
	if body.Amount > 999999.99 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("body.Amount", body.Amount, 999999.99, false))
	}
	if !(body.TransactionCurrency == "USD" || body.TransactionCurrency == "KES" || body.TransactionCurrency == "EUR" || body.TransactionCurrency == "GBP" || body.TransactionCurrency == "AUD" || body.TransactionCurrency == "CHF" || body.TransactionCurrency == "CAD" || body.TransactionCurrency == "ZAR") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.TransactionCurrency", body.TransactionCurrency, []interface{}{"USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR"}))
	}
	if utf8.RuneCountInString(body.Narration) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.Narration", body.Narration, utf8.RuneCountInString(body.Narration), 1, true))
	}
	if utf8.RuneCountInString(body.Narration) > 25 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.Narration", body.Narration, utf8.RuneCountInString(body.Narration), 25, false))
	}
	return
}

// ValidateDestinationAccountCallbackRequestRequestBody runs the validations
// defined on DestinationAccountCallbackRequestRequestBody
func ValidateDestinationAccountCallbackRequestRequestBody(body *DestinationAccountCallbackRequestRequestBody) (err error) {
	if utf8.RuneCountInString(body.ReferenceNumber) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.ReferenceNumber", body.ReferenceNumber, utf8.RuneCountInString(body.ReferenceNumber), 1, true))
	}
	if utf8.RuneCountInString(body.ReferenceNumber) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.ReferenceNumber", body.ReferenceNumber, utf8.RuneCountInString(body.ReferenceNumber), 30, false))
	}
	if utf8.RuneCountInString(body.MobileNumber) < 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.MobileNumber", body.MobileNumber, utf8.RuneCountInString(body.MobileNumber), 10, true))
	}
	if utf8.RuneCountInString(body.MobileNumber) > 13 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.MobileNumber", body.MobileNumber, utf8.RuneCountInString(body.MobileNumber), 13, false))
	}
	if body.Amount < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("body.Amount", body.Amount, 0, true))
	}
	if body.Amount > 999999 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("body.Amount", body.Amount, 999999, false))
	}
	if utf8.RuneCountInString(body.Narration) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.Narration", body.Narration, utf8.RuneCountInString(body.Narration), 1, true))
	}
	if utf8.RuneCountInString(body.Narration) > 25 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.Narration", body.Narration, utf8.RuneCountInString(body.Narration), 25, false))
	}
	return
}
