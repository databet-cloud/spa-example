package fixture_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

var testCompetitor = fixture.Competitor{
	ID:               "competitor1",
	Type:             1,
	HomeAway:         1,
	TemplatePosition: 1,
	Name:             "name1",
	MasterID:         "masterID1",
	CountryCode:      "country1",
}

var competitorTestCases = []struct {
	name     string
	input    patch.Patch
	prev     fixture.Competitor
	expected fixture.Competitor
}{
	{
		name: "patch existent competitors",
		prev: testCompetitor,
		input: patch.Patch{
			"type":              2,
			"home_away":         2,
			"template_position": 2,
			"name":              "name2",
			"master_id":         "masterID2",
			"country_code":      "country2",
		},
		expected: fixture.Competitor{
			ID:               "competitor1",
			Type:             2,
			HomeAway:         2,
			TemplatePosition: 2,
			Name:             "name2",
			MasterID:         "masterID2",
			CountryCode:      "country2",
		},
	},
	{
		name: "patch score",
		prev: testCompetitor,
		input: patch.Patch{
			"score/map:1": patch.Patch{
				"id":     "map:1",
				"type":   "map1",
				"number": 1,
				"points": "1",
			},
		},
		expected: fixture.Competitor{
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
}

func TestCompetitors_WithPatch(t *testing.T) {
	for _, tc := range competitorTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.prev.WithPatch(patch.NewMapTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestCompetitors_ApplyPatch(t *testing.T) {
	for _, tc := range competitorTestCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.prev.ApplyPatch(patch.NewMapTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, tc.prev)
		})
	}
}
