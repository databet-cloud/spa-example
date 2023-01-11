package api

import (
	"context"
)

type Client interface {
	MarketClient
	SportClient
	TournamentClient
	PlayerClient
	TeamClient
	OrganizationClient
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

	// GetAllLocalizedSports finds sports by ids (optional argument, by default all sports are returned)
	// and translates them to a given locale.
	GetAllLocalizedSports(ctx context.Context, locale Locale, ids ...string) ([]SportLocalized, error)
}

type TournamentClient interface {
	FindLocalizedTournamentByID(ctx context.Context, locale Locale, tournamentID string) (*TournamentLocalized, error)
	FindLocalizedTournamentsByIDs(ctx context.Context, locale Locale, tournamentIDs []string) ([]TournamentLocalized, error)
}

type PlayerClient interface {
	FindLocalizedPlayerByID(ctx context.Context, locale Locale, playerID string) (*PlayerLocalized, error)
	FindLocalizedPlayersByIDs(ctx context.Context, locale Locale, playerIDs []string) ([]PlayerLocalized, error)
}

type TeamClient interface {
	FindLocalizedTeamByID(ctx context.Context, locale Locale, teamID string) (*TeamLocalized, error)
	FindLocalizedTeamsByIDs(ctx context.Context, locale Locale, teamIDs []string) ([]TeamLocalized, error)
}

type OrganizationClient interface {
	FindLocalizedOrganizationsByIDs(ctx context.Context, locale Locale, organizationIDs []string) ([]OrganizationLocalized, error)
}
