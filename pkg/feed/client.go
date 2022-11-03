package feed

import (
	"context"
)

type Client interface {
	GetAll(ctx context.Context, bookmakerID string) (cur *RawMessageCursor, version string, err error)
	GetFeedVersion(ctx context.Context, bookmakerID string) (string, error)
	GetLogsFromVersion(ctx context.Context, bookmakerID string, version string) (*RawMessageCursor, error)
}
