package fixture_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

var platformTestCases = []struct {
	name     string
	input    patch.Patch
	prev     fixture.Platform
	expected fixture.Platform
}{
	{
		name: "patch existent platform",
		prev: fixture.Platform{
			Type:             "platform1",
			AllowedCountries: []string{"country1", "country2"},
			Enabled:          false,
		},
		input: patch.Patch{
			"type":              "platform1",
			"allowed_countries": []interface{}{"country3", "country4"},
			"enabled":           true,
		},
		expected: fixture.Platform{
			Type:             "platform1",
			AllowedCountries: []string{"country3", "country4"},
			Enabled:          true,
		},
	},
}

func TestPlatforms_WithPatch(t *testing.T) {
	for _, tc := range platformTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.prev.WithPatch(patch.NewMapTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPlatforms_ApplyPatch(t *testing.T) {
	for _, tc := range platformTestCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.prev.ApplyPatch(patch.NewMapTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, tc.prev)
		})
	}
}
