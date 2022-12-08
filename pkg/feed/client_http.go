package feed

import (
	"context"
	"fmt"
	"net/http"
)

const (
	HeaderLastVersion = "Last-Version"
)

type ClientHTTP struct {
	httpClient *http.Client
	feedURL    string
}

var _ Client = (*ClientHTTP)(nil)

func NewClientHTTP(httpClient *http.Client, feedURL string) *ClientHTTP {
	return &ClientHTTP{
		httpClient: httpClient,
		feedURL:    feedURL,
	}
}

func (c *ClientHTTP) GetAll(ctx context.Context, bookmakerID string) (cursor *RawMessageCursor, lastVersion string, err error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/bookmaker/%s/all", c.feedURL, bookmakerID),
		http.NoBody,
	)
	if err != nil {
		return nil, "", fmt.Errorf("create http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("do http request: %w", err)
	}

	if resp.StatusCode == http.StatusConflict {
		return nil, "", ErrVersionNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("status code %q isn't ok", resp.Status)
	}

	lastVersion = resp.Header.Get(HeaderLastVersion)

	return NewRawMessageCursor(resp.Body), lastVersion, nil
}

func (c *ClientHTTP) GetFeedVersion(ctx context.Context, bookmakerID string) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/bookmaker/%s/logVersion", c.feedURL, bookmakerID),
		http.NoBody,
	)
	if err != nil {
		return "", fmt.Errorf("create http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do http request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code %q isn't ok", resp.Status)
	}

	return resp.Header.Get(HeaderLastVersion), nil
}

func (c *ClientHTTP) GetLogsFromVersion(ctx context.Context, bookmakerID string, version string) (*RawMessageCursor, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/bookmaker/%s/log", c.feedURL, bookmakerID),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	req.Header.Add("Last-Version", version)
	req.URL.Query().Add("longPolling", "true")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	if resp.StatusCode == http.StatusConflict {
		return nil, ErrVersionNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %q isn't ok", resp.Status)
	}

	return NewRawMessageCursor(resp.Body), nil
}
