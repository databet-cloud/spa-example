package feed

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type RawMessageCursor struct {
	decoder    *json.Decoder
	readCloser io.ReadCloser
}

func NewRawMessageCursor(readCloser io.ReadCloser) *RawMessageCursor {
	return &RawMessageCursor{
		decoder:    json.NewDecoder(readCloser),
		readCloser: readCloser,
	}
}

func (c *RawMessageCursor) HasMore() bool {
	return c.decoder.More()
}

func (c *RawMessageCursor) Next(ctx context.Context) (json.RawMessage, error) {
	if !c.HasMore() {
		if err := c.close(); err != nil {
			return nil, err
		}

		return nil, nil
	}

	doneCh := make(chan struct{})
	defer close(doneCh)

	go func() {
		select {
		case <-ctx.Done():
			// close reader to terminate json decoder
			_ = c.close()
		case <-doneCh:
		}
	}()

	var msg json.RawMessage

	if err := c.decoder.Decode(&msg); err != nil {
		if closeErr := c.close(); closeErr != nil {
			return nil, fmt.Errorf("%v after error: %w", closeErr, err)
		}

		return nil, fmt.Errorf("decode msg: %w", err)
	}

	return msg, nil
}

func (c *RawMessageCursor) close() error {
	if c.readCloser == nil {
		return nil
	}

	if err := c.readCloser.Close(); err != nil {
		return fmt.Errorf("close: %w", err)
	}

	return nil
}
