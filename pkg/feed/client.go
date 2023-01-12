package feed

import (
	"context"
)

type Client interface {
	// GetAll returns cursor, containing raw json (sport events), last feed version and error.
	// It can be used to synchronize your system with feed's data, you should use it only once,
	// during the first synchronization.
	GetAll(ctx context.Context) (cur *RawMessageCursor, version string, err error)
	// GetFeedVersion returns an actual version from feed service.
	GetFeedVersion(ctx context.Context) (string, error)
	// GetLogsFromVersion returns cursor, containing raw json (log entries) and error.
	GetLogsFromVersion(ctx context.Context, version string) (*RawMessageCursor, error)
}
