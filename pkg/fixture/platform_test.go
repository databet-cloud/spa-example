package fixture_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

func TestPlatforms_WithPatch(t *testing.T) {
	testCases := []struct {
		name     string
		input    patch.Patch
		prev     fixture.Platforms
		expected fixture.Platforms
	}{
		{
			name: "patch existent platform",
			prev: fixture.Platforms{
				"platform1": fixture.Platform{
					Type:             "platform1",
					AllowedCountries: []string{"country1", "country2"},
					Enabled:          false,
				},
			},
			input: patch.Patch{
				"platform1/type":              "platform1",
				"platform1/allowed_countries": []interface{}{"country3", "country4"},
				"platform1/enabled":           true,
			},
			expected: fixture.Platforms{
				"platform1": fixture.Platform{
					Type:             "platform1",
					AllowedCountries: []string{"country3", "country4"},
					Enabled:          true,
				},
			},
		},
		{
			name: "patch new platforms",
			prev: fixture.Platforms{
				"platform1": fixture.Platform{
					Type:             "platform1",
					AllowedCountries: []string{"country1"},
					Enabled:          false,
				},
			},
			input: patch.Patch{
				"platform2/type":              "platform2",
				"platform2/allowed_countries": []interface{}{"country2"},
				"platform2/enabled":           true,
				//
				"platform3": patch.Patch{
					"type":              "platform3",
					"allowed_countries": []string{"country3"},
					"enabled":           false,
				},
			},
			expected: fixture.Platforms{
				"platform1": fixture.Platform{
					Type:             "platform1",
					AllowedCountries: []string{"country1"},
					Enabled:          false,
				},
				"platform2": fixture.Platform{
					Type:             "platform2",
					AllowedCountries: []string{"country2"},
					Enabled:          true,
				},
				"platform3": fixture.Platform{
					Type:             "platform3",
					AllowedCountries: []string{"country3"},
					Enabled:          false,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := fixture.Platforms(patch.MapPatchable(tc.prev, patch.NewTree(tc.input, "/")))
			assert.Equal(t, tc.expected, actual)
		})
	}
}
