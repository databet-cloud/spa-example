package statistics

import "fmt"

var (
	ErrNotFound           = fmt.Errorf("not found")
	ErrUnknown            = fmt.Errorf("unknown error")
	ErrInvalidCertificate = fmt.Errorf("invalid certificate")
	ErrForbidden          = fmt.Errorf("forbidden")
)

type apiError struct {
	Code string         `json:"code"`
	Data map[string]any `json:"data"`
}
