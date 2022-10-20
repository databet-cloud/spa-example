package market_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databet-cloud/databet-go-sdk/pkg/market"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

//go:embed testdata/odd_patch.json
var oddPatch []byte

func TestOdd_WithPatch(t *testing.T) {
	testCases := []struct {
		name     string
		prev     market.Odd
		rawInput []byte
		input    patch.Patch
		expected market.Odd
	}{
		{
			name: "simple patch",
			prev: market.Odd{
				ID:           "id1",
				Template:     "template1",
				IsActive:     false,
				Status:       1,
				Value:        "1.42",
				StatusReason: "reason1",
			},
			input: patch.Patch{
				"id":            "id2",
				"template":      "template2",
				"is_active":     true,
				"status":        2,
				"value":         "2.42",
				"status_reason": "reason2",
			},
			expected: market.Odd{
				ID:           "id2",
				Template:     "template2",
				IsActive:     true,
				Status:       2,
				Value:        "2.42",
				StatusReason: "reason2",
			},
		},
		{
			name: "raw patch",
			prev: market.Odd{
				ID:           "1",
				Template:     "template1",
				IsActive:     false,
				Status:       1,
				Value:        "1.42",
				StatusReason: "reason1",
			},
			rawInput: oddPatch,
			expected: market.Odd{
				ID:           "2",
				Template:     "even",
				IsActive:     true,
				Status:       2,
				Value:        "99999",
				StatusReason: "reason2",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.rawInput != nil {
				err := json.Unmarshal(tc.rawInput, &tc.input)
				require.NoError(t, err)
			}

			actual := tc.prev.WithPatch(patch.NewTree(tc.input, "/"))
			assert.Equal(t, tc.expected, actual)
		})
	}
}
