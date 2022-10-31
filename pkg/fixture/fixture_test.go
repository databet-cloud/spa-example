package fixture_test

import (
	_ "embed"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
)

var fixtureTestCases = []struct {
	name       string
	rawPatches json.RawMessage
	patches    map[string]json.RawMessage
	expected   fixture.Fixture
}{
	{
		name: "patch all fields",
		rawPatches: json.RawMessage(`
		{
		  "status": 2,
		  "type": 2,
		  "start_time": "2022-01-02T15:04:05Z",
		  "live_coverage": true,
		  "venue": {
			"id": "venue2"
		  },
		  "tournament": {
			"id": "id2",
			"name": "name2",
			"master_id": "masterID2",
			"country_code": "country2"
		  },
		  "competitors": {
			"competitor1": {
			  "id": "competitor1",
			  "meta": {},
			  "name": "team1",
			  "type": 2,
			  "scores": {
				"total": {
				  "id": "total",
				  "type": "total",
				  "number": 1,
				  "points": "1"
				}
			  },
			  "home_away": 1,
			  "master_id": "master_id2",
			  "country_code": "UA",
			  "template_position": 1
			}
		  },
		  "streams": {
			"stream1": {
			  "id": "stream1",
			  "locale": "locale1",
			  "url": "url1",
			  "platforms": {
				"platform1": {
				  "type": "type1",
				  "allowed_countries": ["ua", "eu"],
				  "enabled": true
				}
			  },
			  "priority": 2
			}
		  }
		}`),
		expected: fixture.Fixture{
			Status: 2,
			Type:   2,
			Tournament: fixture.Tournament{
				ID:          "id2",
				Name:        "name2",
				MasterID:    "masterID2",
				CountryCode: "country2",
			},
			Venue:        fixture.Venue{ID: "venue2"},
			LiveCoverage: true,
			StartTime:    mustParseTime("2022-01-02T15:04:05Z"),
			Competitors: fixture.Competitors{
				"competitor1": {
					ID:               "competitor1",
					Type:             2,
					HomeAway:         1,
					TemplatePosition: 1,
					Scores: fixture.Scores{
						"total": {ID: "total", Type: "total", Number: 1, Points: "1"},
					},
					Name:        "team1",
					MasterID:    "master_id2",
					CountryCode: "UA",
				},
			},
			Streams: fixture.Streams{
				"stream1": fixture.Stream{
					ID:     "stream1",
					Locale: "locale1",
					URL:    "url1",
					Platforms: fixture.Platforms{
						"platform1": {Type: "type1", AllowedCountries: []string{"ua", "eu"}, Enabled: true},
					},
					Priority: 2,
				},
			},
		},
	},
}

func TestFixture_ApplyPatch(t *testing.T) {
	for _, tc := range fixtureTestCases {
		t.Run(tc.name, func(t *testing.T) {
			var f fixture.Fixture

			if tc.rawPatches != nil {
				err := json.Unmarshal(tc.rawPatches, &tc.patches)
				require.NoError(t, err)
			}

			for path, value := range tc.patches {
				err := f.ApplyPatch(path, value)
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expected, f)
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
