package api

import "time"

type MarketType string

func (mt MarketType) String() string {
	return string(mt)
}

const (
	MarketTypeMatch    MarketType = "match"
	MarketTypeOutright MarketType = "outright"
)

type Market struct {
	ID            int                 `json:"id"`
	Version       string              `json:"version"`
	Priority      int                 `json:"priority"`
	Tags          []string            `json:"tags"`
	Localizations MarketLocalizations `json:"localizations"`
	Outcomes      map[string]string   `json:"outcomes"`
	Specifiers    map[string]string   `json:"specifiers"`
	// SportIDs where market can be used
	SportIDs   []string   `json:"sport_ids"`
	Deprecated bool       `json:"deprecated"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Type       MarketType `json:"type"`
}

type MarketLocalizations map[Locale]MarketLocalization

type MarketLocalization struct {
	Locale Locale `json:"locale"`
	// Template is a market name template
	Template string `json:"template"`
}

type MarketLocalized struct {
	ID       int      `json:"id"`
	Version  string   `json:"version"`
	Priority int      `json:"priority"`
	Tags     []string `json:"tags"`
	Locale   Locale   `json:"locale"`
	// Template is a market name template
	Template   string            `json:"template"`
	Outcomes   map[string]string `json:"outcomes"`
	Specifiers map[string]string `json:"specifiers"`
	// SportIDs where market can be used
	SportIDs   []string   `json:"sport_ids"`
	Deprecated bool       `json:"deprecated"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Type       MarketType `json:"type"`
}
