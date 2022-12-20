package api

import (
	"context"
)

type Client interface {
	FindMarketByID(ctx context.Context, marketID int) (*Market, error)
	FindMarketsByIDs(ctx context.Context, marketIDs []int) ([]Market, error)
	FindLocalizedMarketsByIDs(ctx context.Context, locale Locale, marketIDs []int) ([]MarketLocalized, error)
	FindMarketsByFilters(ctx context.Context, filters *MarketFilters) ([]Market, error)
	FindLocalizedMarketsByFilters(ctx context.Context, locale Locale, filters *MarketFilters) ([]MarketLocalized, error)

	FindSportByID(ctx context.Context, sportID string) (*Sport, error)
	FindSportsByFilters(ctx context.Context, filters *SportFilters) ([]Sport, error)
	GetAllLocalizedSports(ctx context.Context, locale Locale, ids ...string) ([]SportLocalized, error)

	FindTournamentByID(ctx context.Context, tournamentID string) (*Tournament, error)
	SearchTournaments(ctx context.Context, req *SearchTournamentsRequest) (*SearchTournamentsResponse, error)
	FindLocalizedTournamentByID(ctx context.Context, locale Locale, tournamentID string) (*TournamentLocalized, error)
	SearchLocalizedTournaments(ctx context.Context, locale Locale, req *SearchTournamentsRequest) (*SearchLocalizedTournamentsResponse, error)
}
