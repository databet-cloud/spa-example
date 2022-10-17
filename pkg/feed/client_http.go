package feed

import (
	"context"
	"encoding/json"
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

func (c *ClientHTTP) GetAll(
	ctx context.Context,
	bookmakerID string,
	receiveCh chan<- json.RawMessage,
) (lastVersion string, err error) {
	defer close(receiveCh)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/bookmaker/%s/all", c.feedURL, bookmakerID),
		nil, // empty body
	)
	if err != nil {
		return "", fmt.Errorf("create http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do http request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusConflict {
		return "", ErrVersionNotFound
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code %q isn't ok", resp.Status)
	}

	lastVersion = resp.Header.Get(HeaderLastVersion)

	decoder := json.NewDecoder(resp.Body)
	for decoder.More() {
		var msg json.RawMessage

		err := decoder.Decode(&msg)
		if err != nil {
			return "", fmt.Errorf("decode message: %w", err)
		}

		select {
		case receiveCh <- msg:
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}

	return lastVersion, nil
}

func (c *ClientHTTP) GetFeedVersion(ctx context.Context, bookmakerID string) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/bookmaker/%s/logVersion", c.feedURL, bookmakerID),
		nil, // empty body
	)
	if err != nil {
		return "", fmt.Errorf("create http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do http request: %w", err)
	}

	defer resp.Body.Close()

	return resp.Header.Get(HeaderLastVersion), nil
}

func (c *ClientHTTP) GetLogsFromVersion(
	ctx context.Context,
	bookmakerID string,
	version string,
	receiveCh chan<- json.RawMessage,
) error {
	defer close(receiveCh)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/bookmaker/%s/log", c.feedURL, bookmakerID),
		nil, // empty body
	)
	if err != nil {
		return fmt.Errorf("create http request: %w", err)
	}

	req.Header.Add("Last-Version", version)
	req.URL.Query().Add("longPolling", "true")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do http request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusConflict {
		return ErrVersionNotFound
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %q isn't ok", resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	for decoder.More() {
		var msg json.RawMessage

		err := decoder.Decode(&msg)
		if err != nil {
			return fmt.Errorf("decode message: %w", err)
		}

		select {
		case receiveCh <- msg:
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
