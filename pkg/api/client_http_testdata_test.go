package api_test

import (
	"github.com/databet-cloud/databet-go-sdk/pkg/api"
)

var DefaultLocalizedTournament = api.TournamentLocalized{
	ID:             "betting:0:csaaa-esports_counter_strike-esports_counter_strike",
	OriginalID:     "csaaa-esports_counter_strike",
	Type:           1,
	SourceID:       "sourceID",
	Version:        "0ee56e0b-661a-4f9b-9cb7-8abf8b7fc3e7",
	ProviderID:     "betting:0",
	SportID:        "esports_counter_strike",
	CountryCode:    "WW-AFR",
	Organization:   "",
	OrganizationID: "organizationID",
	Duplicated:     false,
	Logo: struct {
		File string `json:"file"`
		URL  string `json:"url"`
	}{
		URL: "binary-storage-stage-betting.ginsp.net/tournaments/564px-ESL_Pro_League5a8d7358ad21b9668747145ae1bd8fd847d441215577.png",
	},
	LimitGroups: struct {
		PrematchID *string `json:"prematch_id"`
		LiveID     *string `json:"live_id"`
	}{
		PrematchID: pointTo("prematchID"),
		LiveID:     pointTo("liveID"),
	},
	DateStart: mustParseTime("1977-01-01T00:00:00+00:00"),
	DateEnd:   mustParseTime("1977-01-23T01:00:00+00:00"),
	UpdatedAt: mustParseTime("2022-12-16T16:38:01+00:00"),
	Meta: map[string]any{
		"show_tournament_info": true,
		"prize_pool":           "100 000$",
		"live_coverage":        true,
	},
	LocalizationOverridden: false,
	Locale:                 "en",
	Name:                   "CSAAA",
	Description:            "123",
	Keywords:               []string{"csaaa"},
}

var DefaultLocalizedPlayer = api.PlayerLocalized{
	ID:              "betting:1:sr:competitor:100221",
	OriginalID:      "sr:competitor:100221",
	Type:            1,
	SourceID:        "",
	Version:         "7db572ae-d443-43cf-8d59-0eedcb74bcfb",
	ProviderID:      "betting:1",
	ProviderVersion: "",
	SportID:         "tennis",
	LogoAutoUpdate:  true,
	CountryCode:     "",
	DateOfBirth:     mustParseTime("1970-01-01T00:00:00+00:00"),
	Statistics:      []interface{}{},
	Duplicated:      false,
	Locale:          "en",
	Name:            "",
	Nickname:        "Romero De Avila Senise, Alberto",
	Description:     "",
	Keywords:        []string{"romero-de-avila-senise-alberto"},
	UpdatedAt:       mustParseTime("2021-12-23T17:07:19+00:00"),
	Meta:            map[string]interface{}{},
}

var DefaultLocalizedTeam = api.TeamLocalized{
	ID:             "betting:0:1-esports_dota_2",
	CountryCode:    "AX",
	Name:           "3NaVi123123",
	SportID:        "esports_dota_2",
	OrganizationID: "",
	Logo:           api.Logo{URL: "test-storage/team/surprised-cat5aa905ff3a61a819205529.webp"},
	Description:    "1",
	Keywords:       []string{"3navi123123"},
}
