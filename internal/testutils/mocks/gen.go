//go:generate go run github.com/golang/mock/mockgen -destination=round_tripper.go -package=mocks -mock_names=RoundTripper=RoundTripper net/http RoundTripper
package mocks
