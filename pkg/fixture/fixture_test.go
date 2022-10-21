package fixture_test

import (
	_ "embed"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

//go:embed testdata/fixture_patch.json
var fixturePatch []byte

func TestFixtureDiff(t *testing.T) {
	testCases := []struct {
		name     string
		rawInput []byte
		input    patch.Patch
		expected fixture.Fixture
	}{
		{
			name: "patch scalar types",
			input: patch.Patch{
				"status":        1,
				"type":          1,
				"start_time":    time.Unix(100, 0).UTC(),
				"live_coverage": true,
			},
			expected: fixture.Fixture{
				Status:       1,
				Type:         1,
				StartTime:    time.Unix(100, 0).UTC(),
				LiveCoverage: true,
			},
		},
		{
			name: "patch tournament",
			input: patch.Patch{
				"tournament/id":           "id1",
				"tournament/name":         "name1",
				"tournament/master_id":    "master_id1",
				"tournament/country_code": "UA",
			},
			expected: fixture.Fixture{
				Tournament: fixture.Tournament{
					ID:          "id1",
					Name:        "name1",
					MasterID:    "master_id1",
					CountryCode: "UA",
				},
			},
		},
		{
			name:     "patch venue",
			input:    patch.Patch{"venue/id": "id1"},
			expected: fixture.Fixture{Venue: fixture.Venue{ID: "id1"}},
		},
		{
			name: "patch streams",
			input: patch.Patch{
				"streams/stream1/id": "stream1",
			},
			expected: fixture.Fixture{Streams: map[string]fixture.Stream{
				"stream1": {ID: "stream1"},
			}},
		},
		{
			name: "patch competitors",
			input: patch.Patch{
				"competitors/competitor1/id": "competitor1",
			},
			expected: fixture.Fixture{Competitors: map[string]fixture.Competitor{
				"competitor1": {ID: "competitor1"},
			}},
		},
		{
			name:     "json patch",
			rawInput: fixturePatch,
			expected: fixture.Fixture{
				ID:       "", // isn't patched, can't change
				Version:  0,  // isn't patched, can't change
				OwnerID:  "", // isn't patched, can't change
				Template: "", // isn't patched, can't change
				Status:   8,
				Type:     1,
				SportID:  "", // isn't patched, can't change
				Tournament: fixture.Tournament{
					ID:          "betting:0:dsd-esports_counter_strike-esports_counter_strike",
					Name:        "tournament1",
					MasterID:    "betting:0:dsd-esports_counter_strike-esports_counter_strike",
					CountryCode: "UA",
				},
				Venue: fixture.Venue{ID: "id1"},
				Competitors: fixture.Competitors{
					"betting:0:10-esports_counter_strike": fixture.Competitor{
						ID:               "betting:0:10-esports_counter_strike",
						Type:             2,
						HomeAway:         1,
						TemplatePosition: 1,
						Scores: fixture.Scores{
							"total": fixture.Score{
								ID:     "total",
								Type:   "total",
								Points: "1",
								Number: 1,
							},
						},
						Name:        "team1",
						MasterID:    "betting:0:10-esports_counter_strike",
						CountryCode: "UA",
					},
					"betting:0:16-esports_counter_strike": fixture.Competitor{
						ID:               "betting:0:16-esports_counter_strike",
						Type:             2,
						HomeAway:         2,
						TemplatePosition: 2,
						Scores: fixture.Scores{
							"total": fixture.Score{
								ID:     "total",
								Type:   "total",
								Points: "2",
								Number: 2,
							},
						},
						Name:        "hmm",
						MasterID:    "betting:0:16-esports_counter_strike",
						CountryCode: "UA",
					},
				},
				Streams: fixture.Streams{
					"stream1": {
						ID:     "stream1",
						Locale: "locale1",
						URL:    "url1",
						Platforms: fixture.Platforms{
							"platform1": {Type: "platform1", AllowedCountries: []string{"ua", "eu"}, Enabled: true},
						},
						Priority: 2,
					},
				},
				LiveCoverage: true,
				StartTime:    mustParseTime("2022-10-14T17:23:00Z"),
				Flags:        0,                                                   // isn't patched, can't change
				CreatedAt:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), // isn't patched, can't change
				UpdatedAt:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), // isn't patched, can't change
				PublishedAt:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), // isn't patched, can't change
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var f fixture.Fixture

			if tc.rawInput != nil {
				err := json.Unmarshal(tc.rawInput, &tc.input)
				require.NoError(t, err)
			}

			actual := f.WithPatch(patch.NewTree(tc.input, "/"))
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func mustParseTime(t string) time.Time {
	res, err := time.Parse(time.RFC3339, t)
	if err != nil {
		panic(err)
	}

	return res
}
