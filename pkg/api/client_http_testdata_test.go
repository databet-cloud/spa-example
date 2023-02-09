package api_test

import (
	"github.com/databet-cloud/databet-go-sdk/pkg/api"
)

var DefaultLocalizedTournament = api.TournamentLocalized{
	ID:             "betting:0:csaaa-esports_counter_strike-esports_counter_strike",
	SportID:        "esports_counter_strike",
	CountryCode:    "WW-AFR",
	OrganizationID: "organizationID",
	Logo: api.Logo{
		URL: "binary-storage-stage-betting.ginsp.net/tournaments/564px-ESL_Pro_League5a8d7358ad21b9668747145ae1bd8fd847d441215577.png",
	},
	Meta: map[string]any{
		"prize_pool": "100 000$",
	},
	Locale:      "en",
	Name:        "CSAAA",
	Description: "123",
	Keywords:    []string{"csaaa"},
}

var DefaultLocalizedPlayer = api.PlayerLocalized{
	ID:          "betting:1:sr:competitor:100221",
	SportID:     "tennis",
	CountryCode: "",
	Locale:      "en",
	Nickname:    "Romero De Avila Senise, Alberto",
	Description: "",
	Keywords:    []string{"romero-de-avila-senise-alberto"},
}

var DefaultLocalizedTeam = api.TeamLocalized{
	ID:          "betting:0:1-esports_dota_2",
	CountryCode: "AX",
	Name:        "3NaVi123123",
	SportID:     "esports_dota_2",
	Logo:        api.Logo{URL: "test-storage/team/surprised-cat5aa905ff3a61a819205529.webp"},
	Description: "1",
	Keywords:    []string{"3navi123123"},
}

var DefaultLocalizedOrganization = api.OrganizationLocalized{
	ID:       "0754ec51-964b-4d00-9a5b-dbeab376eac3",
	Name:     "ITF. Men",
	SportIDs: []string{"tennis"},
	Locale:   "en",
}

var DefaultSportEventLimit = api.SportEventLimit{
	BookmakerID:  "test",
	SportEventID: "035896b9-cfcf-4447-a28d-bac27aa00471",
	FixtureState: api.LimitFixtureStateInProgress,
	BetLimit: api.BetLimit{
		BetDelay: 10,
	},
	Risks: api.Risks{
		MaxBetRisk: "100.00",
	},
	MarketsLimits: []api.MarketLimit{
		{MarketID: "123", MaxBetRiskMultiplier: "1.42"},
	},
	LimitUpdatedAt: mustParseTime("2023-02-06T09:19:16.919Z"),
	Version:        "cfgcc52mg7redge4kn6g",
	UpdatedAt:      mustParseTime("2023-02-06T09:19:16.926Z"),
}
