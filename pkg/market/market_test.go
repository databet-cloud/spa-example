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

//go:embed testdata/market_patch.json
var marketPatch []byte

func TestMarket_WithPatch(t *testing.T) {
	testCases := []struct {
		name     string
		prev     market.Market
		rawInput []byte
		input    patch.Patch
		expected market.Market
	}{
		{
			name: "simple patch",
			prev: market.Market{
				ID:         "id1",
				TypeID:     1,
				Template:   "template1",
				Status:     1,
				Specifiers: map[string]string{"1": "1", "10": "10"},
				Meta:       map[string]interface{}{"1": "1", "10": "10"},
				Flags:      1,
			},
			input: patch.Patch{
				"id":           "id2",
				"type_id":      2,
				"template":     "template2",
				"status":       2,
				"specifiers/1": "2",
				"meta/1":       "2",
				"flags":        2,
			},
			expected: market.Market{
				ID:         "id2",
				TypeID:     2,
				Template:   "template2",
				Status:     2,
				Specifiers: map[string]string{"1": "2", "10": "10"},
				Meta:       map[string]interface{}{"1": "2", "10": "10"},
				Flags:      2,
			},
		},
		{
			name:     "json patch full market",
			prev:     market.Market{},
			rawInput: marketPatch,
			expected: market.Market{
				ID:       "1",
				TypeID:   1,
				Template: "Winner",
				Status:   2,
				Odds: market.Odds{
					"1": market.Odd{
						ID:           "1",
						Template:     "{$competitor1}",
						IsActive:     true,
						Status:       1,
						Value:        "1.88",
						StatusReason: "reason1",
					},
					"2": market.Odd{
						ID:           "2",
						Template:     "{$competitor2}",
						Status:       1,
						IsActive:     true,
						Value:        "2.09",
						StatusReason: "reason2",
					},
				},
				Specifiers:  map[string]string{"1": "2"},
				Meta:        nil,
				Flags:       2,
				IsDefective: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.rawInput != nil {
				err := json.Unmarshal(tc.rawInput, &tc.input)
				require.NoError(t, err)
			}

			actual, err := tc.prev.WithPatch(patch.NewTree(tc.input, "/"))
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
