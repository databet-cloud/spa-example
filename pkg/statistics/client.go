package statistics

import "context"

type Client interface {
	// FindFixtureStatisticsByID find specified version  of fixture statistics.
	FindFixtureStatisticsByID(ctx context.Context, fixtureID string, version string) (Statistics, error)
}
