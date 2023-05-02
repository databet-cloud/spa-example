package statistics

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
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

	rc, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	defer rc.Close()

	var resp struct {
		Statistics Statistics `json:"statistics"`
	}

	if err = sonic.ConfigDefault.NewDecoder(rc).Decode(&resp); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Statistics, nil
}

func (c *ClientHTTP) doAPIRequest(httpReq *http.Request) (rc io.ReadCloser, err error) {
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer func() {
		if err != nil {
			httpResp.Body.Close()
		}
	}()

	switch httpResp.StatusCode {
	case http.StatusUnauthorized:
		return nil, ErrInvalidCertificate
	case http.StatusForbidden:
		return nil, ErrForbidden
	case http.StatusNotFound:
		return nil, ErrNotFound
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w, extra data: %v", ErrUnknown, err)
	}

	return httpResp.Body, nil
}
