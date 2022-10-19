package fixture_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

func TestFixtureDiff(t *testing.T) {
	testCases := []struct {
		name     string
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var f fixture.Fixture

			actual := f.WithPatch(patch.NewTree(tc.input, "/"))
			assert.Equal(t, tc.expected, actual)
		})
	}
}
