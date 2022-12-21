package feed

import (
	"context"
)

type Client interface {
	GetAll(ctx context.Context) (cur *RawMessageCursor, version string, err error)
	GetFeedVersion(ctx context.Context) (string, error)
	GetLogsFromVersion(ctx context.Context, version string) (*RawMessageCursor, error)
}
