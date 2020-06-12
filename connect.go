package listener

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	connect "github.com/wondenge/coop-notifiers/gen/connect"
)

// connect service example implementation.
// The example methods log the requests and return zero values.
type connectsrvc struct {
	logger log.Logger
}

// NewConnect returns the connect service implementation.
func NewConnect(logger log.Logger) connect.Service {
	return &connectsrvc{logger}
}

// IFTAccountToAccount Early Hints Callback Request
func (s *connectsrvc) IFTAccountToAccount(ctx context.Context, p *connect.IFTAccountToAccountCallbackRequest) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("connect.IFTAccountToAccount"))
	return
}

// PesaLinkSendToAccount Early Hints Callback Request
func (s *connectsrvc) PesaLinkSendToAccount(ctx context.Context, p *connect.PesaLinkSendToAccountCallbackRequest) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("connect.PesaLinkSendToAccount"))
	return
}

// PesaLinkSendToPhone Early Hints Callback Request
func (s *connectsrvc) PesaLinkSendToPhone(ctx context.Context, p *connect.PesaLinkSendToPhoneCallbackRequest) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("connect.PesaLinkSendToPhone"))
	return
}

// SendToMPesa Early Hints Callback Request
func (s *connectsrvc) SendToMPesa(ctx context.Context, p *connect.SendToMpesaCallbackRequest) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("connect.SendToMPesa"))
	return
}
