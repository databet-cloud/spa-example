package feed

import (
	"context"
	"encoding/json"
)

type Client interface {
	GetAll(ctx context.Context, bookmakerID string, receiveCh chan<- json.RawMessage) (version string, err error)
	GetFeedVersion(ctx context.Context, bookmakerID string) (string, error)
	GetLogsFromVersion(ctx context.Context, bookmakerID string, version string, receiveCh chan<- json.RawMessage) error
}
