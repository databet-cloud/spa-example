package mts

import (
	"fmt"
)

const (
	errCodeUnknownSportEvent      = "UNKNOWN_SPORT_EVENT"
	errCodeUnknownMarket          = "UNKNOWN_MARKET"
	errCodeInvalidRequest         = "INVALID_REQUEST"
	errCodeInternalError          = "INTERNAL_ERROR"
	errCodeAuthInvalidCertificate = "AUTH_INVALID_CERTIFICATE"
)

var (
	ErrAccessForIPDenied  = fmt.Errorf("access for your ip denied")
	ErrInternal           = fmt.Errorf("internal server error")
	ErrBadRequest         = fmt.Errorf("bad request")
	ErrInvalidCertificate = fmt.Errorf("invalid certificate")
	ErrUnknownSportEvent  = fmt.Errorf("unknown sport event")
	ErrUnknownMarket      = fmt.Errorf("unknown market")
	ErrUnknown            = fmt.Errorf("unknown error")
)

type apiError struct {
	Code string         `json:"code"`
	Data map[string]any `json:"data"`
}

func (e *apiError) Error() string {
	return fmt.Sprintf("code: %s, data: %v", e.Code, e.Data)
}
