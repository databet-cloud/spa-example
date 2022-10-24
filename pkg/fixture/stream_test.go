package fixture_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

var streamTestCases = []struct {
	name     string
	input    patch.Patch
	prev     fixture.Stream
	expected fixture.Stream
}{
	{
		name: "patch all fields",
		prev: fixture.Stream{
			ID:       "id1",
			Locale:   "locale1",
			URL:      "url1",
			Priority: 1,
		},
		input: patch.Patch{
			"id":       "id2",
			"locale":   "locale2",
			"url":      "url2",
			"priority": 2,
		},
		expected: fixture.Stream{
			ID:       "id2",
			Locale:   "locale2",
			URL:      "url2",
			Priority: 2,
		},
	},
	{
		name: "patch platform",
		prev: fixture.Stream{
			ID: "id1",
			Platforms: map[string]fixture.Platform{
				"platform1": {Type: "platform1"},
			},
		},
		input: patch.Patch{
			"platforms/platform1/enabled": true,
		},
		expected: fixture.Stream{
			ID: "id1",
			Platforms: map[string]fixture.Platform{
				"platform1": {Type: "platform1", Enabled: true},
			},
		},
	},
}

func TestStream_WithPatch(t *testing.T) {
	for _, tc := range streamTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.prev.WithPatch(patch.NewMapTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestStream_ApplyPatch(t *testing.T) {
	for _, tc := range streamTestCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.prev.ApplyPatch(patch.NewMapTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, tc.prev)
		})
	}
}
