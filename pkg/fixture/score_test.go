package fixture_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

var scoreTestCases = []struct {
	name     string
	input    patch.Patch
	prev     fixture.Score
	expected fixture.Score
}{
	{
		name: "patch existent scores",
		prev: fixture.Score{ID: "map:1", Type: "map1", Points: "1", Number: 1},
		input: patch.Patch{
			"id":     "map:1",
			"type":   "changed",
			"points": "changed",
			"number": 1000,
		},
		expected: fixture.Score{
			ID:     "map:1",
			Type:   "changed",
			Points: "changed",
			Number: 1000,
		},
	},
}

func TestScores_WithPatch(t *testing.T) {
	for _, tc := range scoreTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.prev.WithPatch(patch.NewMapTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestScores_ApplyPatch(t *testing.T) {
	for _, tc := range scoreTestCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.prev.ApplyPatch(patch.NewMapTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, tc.prev)
		})
	}
}
