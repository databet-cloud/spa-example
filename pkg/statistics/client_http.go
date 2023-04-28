package statistics

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

type ClientHTTP struct {
	httpClient *http.Client
	apiURL     string
}

var _ Client = (*ClientHTTP)(nil)

func NewClientHTTP(httpClient *http.Client, apiURL string) *ClientHTTP {
	return &ClientHTTP{
		httpClient: httpClient,
		apiURL:     apiURL,
	}
}

func (c *ClientHTTP) FindFixtureStatisticsByID(ctx context.Context, fixtureID string, version string) (Statistics, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/statistics/%s/%s", c.apiURL, fixtureID, version),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Statistics Statistics `json:"statistics"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Statistics, nil
}

func (c *ClientHTTP) doAPIRequest(httpReq *http.Request) ([]byte, error) {
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	rawBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	switch httpResp.StatusCode {
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%w, response body: %q", ErrInvalidCertificate, string(rawBody))
	case http.StatusForbidden:
		return nil, ErrForbidden
	case http.StatusNotFound:
		return nil, ErrNotFound
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w, extra data: %v", ErrUnknown, err)
	}

	return rawBody, nil
}
