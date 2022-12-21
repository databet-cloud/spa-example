package api

import (
	"context"
)

type Client interface {
	MarketClient
	SportClient
	TournamentClient
	PlayerClient
}

type MarketClient interface {
	FindMarketByID(ctx context.Context, marketID int) (*Market, error)
	FindMarketsByIDs(ctx context.Context, marketIDs []int) ([]Market, error)
	FindLocalizedMarketsByIDs(ctx context.Context, locale Locale, marketIDs []int) ([]MarketLocalized, error)
	FindMarketsByFilters(ctx context.Context, filters *MarketFilters) ([]Market, error)
	FindLocalizedMarketsByFilters(ctx context.Context, locale Locale, filters *MarketFilters) ([]MarketLocalized, error)
}

type SportClient interface {
	FindSportByID(ctx context.Context, sportID string) (*Sport, error)
	FindSportsByFilters(ctx context.Context, filters *SportFilters) ([]Sport, error)
	GetAllLocalizedSports(ctx context.Context, locale Locale, ids ...string) ([]SportLocalized, error)
}

type TournamentClient interface {
	FindTournamentByID(ctx context.Context, tournamentID string) (*Tournament, error)
	FindTournamentsByIDs(ctx context.Context, tournamentIDs []string) ([]Tournament, error)
	SearchTournaments(ctx context.Context, req *SearchTournamentsRequest) (*SearchTournamentsResponse, error)
	FindLocalizedTournamentByID(ctx context.Context, locale Locale, tournamentID string) (*TournamentLocalized, error)
	FindLocalizedTournamentsByIDs(ctx context.Context, locale Locale, tournamentIDs []string) ([]TournamentLocalized, error)
	SearchLocalizedTournaments(ctx context.Context, locale Locale, req *SearchTournamentsRequest) (*SearchLocalizedTournamentsResponse, error)
}

type PlayerClient interface {
	FindPlayerByID(ctx context.Context, playerID string) (*Player, error)
	FindPlayersByIDs(ctx context.Context, playerIDs []string) ([]Player, error)
	SearchPlayers(ctx context.Context, req *SearchPlayersRequest) (*SearchPlayersResponse, error)
	FindLocalizedPlayerByID(ctx context.Context, locale Locale, playerID string) (*PlayerLocalized, error)
	FindLocalizedPlayersByIDs(ctx context.Context, locale Locale, playerIDs []string) ([]PlayerLocalized, error)
}
