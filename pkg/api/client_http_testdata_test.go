package api_test

import "github.com/databet-cloud/databet-go-sdk/pkg/api"

var DefaultTournament = api.Tournament{
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
	Localizations: []api.TournamentLocalization{
		{
			Locale:      "en",
			Name:        "CSAAA",
			Description: "123",
			Keywords:    []string{"csaaa"},
		},
	},
}

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
