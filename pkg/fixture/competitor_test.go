package fixture_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

//nolint:gochecknoglobals // testdata
var testCompetitors = fixture.Competitors{
	"competitor1": fixture.Competitor{
		ID:               "competitor1",
		Type:             1,
		HomeAway:         1,
		TemplatePosition: 1,
		Name:             "name1",
		MasterID:         "masterID1",
		CountryCode:      "country1",
	},
}

func TestCompetitors_WithPatch(t *testing.T) {
	testCases := []struct {
		name     string
		input    patch.Patch
		prev     fixture.Competitors
		expected fixture.Competitors
	}{
		{
			name: "patch existent competitors",
			prev: testCompetitors,
			input: patch.Patch{
				"competitor1/type":              2,
				"competitor1/home_away":         2,
				"competitor1/template_position": 2,
				"competitor1/name":              "name2",
				"competitor1/master_id":         "masterID2",
				"competitor1/country_code":      "country2",
			},
			expected: fixture.Competitors{
				"competitor1": fixture.Competitor{
					ID:               "competitor1",
					Type:             2,
					HomeAway:         2,
					TemplatePosition: 2,
					Name:             "name2",
					MasterID:         "masterID2",
					CountryCode:      "country2",
				},
			},
		},
		{
			name: "patch new competitors",
			prev: testCompetitors,
			input: patch.Patch{
				"competitor2/id":                "competitor2",
				"competitor2/type":              2,
				"competitor2/home_away":         2,
				"competitor2/template_position": 2,
				"competitor2/name":              "name2",
				"competitor2/master_id":         "masterID2",
				"competitor2/country_code":      "country2",
				//
				"competitor3": patch.Patch{
					"id":                "competitor3",
					"type":              3,
					"home_away":         3,
					"template_position": 3,
					"name":              "name3",
					"master_id":         "masterID3",
					"country_code":      "country3",
				},
			},
			expected: fixture.Competitors{
				"competitor1": fixture.Competitor{
					ID:               "competitor1",
					Type:             1,
					HomeAway:         1,
					TemplatePosition: 1,
					Name:             "name1",
					MasterID:         "masterID1",
					CountryCode:      "country1",
				},
				"competitor2": fixture.Competitor{
					ID:               "competitor2",
					Type:             2,
					HomeAway:         2,
					TemplatePosition: 2,
					Name:             "name2",
					MasterID:         "masterID2",
					CountryCode:      "country2",
				},
				"competitor3": fixture.Competitor{
					ID:               "competitor3",
					Type:             3,
					HomeAway:         3,
					TemplatePosition: 3,
					Name:             "name3",
					MasterID:         "masterID3",
					CountryCode:      "country3",
				},
			},
		},
		{
			name: "patch score",
			prev: testCompetitors,
			input: patch.Patch{
				"competitor1/score/map:1": patch.Patch{
					"id":     "map:1",
					"type":   "map1",
					"number": 1,
					"points": "1",
				},
			},
			expected: fixture.Competitors{
				"competitor1": fixture.Competitor{
					ID:               "competitor1",
					Type:             1,
					HomeAway:         1,
					TemplatePosition: 1,
					Name:             "name1",
					MasterID:         "masterID1",
					CountryCode:      "country1",
					Scores: fixture.Scores{
						"map:1": fixture.Score{
							ID:     "map:1",
							Type:   "map1",
							Points: "1",
							Number: 1,
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := patch.MapPatchable(tc.prev, patch.NewTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, fixture.Competitors(actual))
		})
	}
}
