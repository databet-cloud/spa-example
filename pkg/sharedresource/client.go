package sharedresource

import "context"

type Client interface {
	FindMarketByID(ctx context.Context, marketID int) (*Market, error)
	FindMarketsByIDs(ctx context.Context, marketIDs []int) ([]Market, error)
	FindLocalizedMarketsByIDs(ctx context.Context, locale Locale, marketIDs []int) ([]MarketLocalized, error)
	FindMarketsByFilters(ctx context.Context, filters *MarketFilters) ([]Market, error)
	FindLocalizedMarketsByFilters(ctx context.Context, locale Locale, filters *MarketFilters) ([]MarketLocalized, error)

	FindSportByID(ctx context.Context, sportID string) (*Sport, error)
	GetAllSports(ctx context.Context, tags ...string) ([]Sport, error)
	GetAllLocalizedSports(ctx context.Context, locale Locale, ids ...string) ([]SportLocalized, error)
}
