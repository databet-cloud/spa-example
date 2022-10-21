package fixture_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

//nolint:gochecknoglobals // testdata
var testScores = fixture.Scores{
	"map:1": fixture.Score{
		ID:     "map:1",
		Type:   "map1",
		Points: "1",
		Number: 1,
	},
	"map:2": fixture.Score{
		ID:     "map:2",
		Type:   "map2",
		Points: "2",
		Number: 2,
	},
}

func TestScores_WithPatch(t *testing.T) {
	testCases := []struct {
		name     string
		input    patch.Patch
		prev     fixture.Scores
		expected fixture.Scores
	}{
		{
			name: "patch existent scores",
			prev: testScores,
			input: patch.Patch{
				"map:1/id":     "map:1",
				"map:1/type":   "changed",
				"map:1/points": "changed",
				"map:1/number": 1000,
			},
			expected: fixture.Scores{
				"map:1": fixture.Score{
					ID:     "map:1",
					Type:   "changed",
					Points: "changed",
					Number: 1000,
				},
				"map:2": fixture.Score{
					ID:     "map:2",
					Type:   "map2",
					Points: "2",
					Number: 2,
				},
			},
		},
		{
			name: "patch new scores",
			prev: testScores,
			input: patch.Patch{
				"map:3/id":     "map:3",
				"map:3/type":   "map3",
				"map:3/number": 3,
				"map:3/points": "3",
				//
				"map:4": patch.Patch{
					"id":     "map:4",
					"type":   "map4",
					"number": 4,
					"points": "4",
				},
			},
			expected: fixture.Scores{
				"map:1": fixture.Score{
					ID:     "map:1",
					Type:   "map1",
					Points: "1",
					Number: 1,
				},
				"map:2": fixture.Score{
					ID:     "map:2",
					Type:   "map2",
					Points: "2",
					Number: 2,
				},
				"map:3": fixture.Score{
					ID:     "map:3",
					Type:   "map3",
					Points: "3",
					Number: 3,
				},
				"map:4": fixture.Score{
					ID:     "map:4",
					Type:   "map4",
					Points: "4",
					Number: 4,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := patch.MapPatchable(tc.prev, patch.NewTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, fixture.Scores(actual))
		})
	}
}
