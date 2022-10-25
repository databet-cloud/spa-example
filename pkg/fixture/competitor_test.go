package fixture_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
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

func TestCompetitors_ApplyPatch(t *testing.T) {
	testCases := []struct {
		name     string
		prev     fixture.Competitor
		patches  map[string]json.RawMessage
		expected fixture.Competitor
	}{
		{
			name: "patch all fields",
			prev: testCompetitor,
			patches: map[string]json.RawMessage{
				"id":                json.RawMessage(`"competitor2"`),
				"type":              json.RawMessage(`2`),
				"home_away":         json.RawMessage(`2`),
				"template_position": json.RawMessage(`2`),
				"name":              json.RawMessage(`"name2"`),
				"master_id":         json.RawMessage(`"masterID2"`),
				"country_code":      json.RawMessage(`"country2"`),
				"scores":            json.RawMessage(`{"total": { "id": "total", "type": "total", "number": 1, "points": "1" }}`),
			},
			expected: fixture.Competitor{
				ID:               "competitor2",
				Type:             2,
				HomeAway:         2,
				TemplatePosition: 2,
				Name:             "name2",
				MasterID:         "masterID2",
				CountryCode:      "country2",
				Scores:           fixture.Scores{"total": {ID: "total", Type: "total", Points: "1", Number: 1}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for path, value := range tc.patches {
				err := tc.prev.ApplyPatch(path, value)
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expected, tc.prev)
		})
	}
}
