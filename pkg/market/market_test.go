package market_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/market"
)

//go:embed testdata/market_patch.json
var marketPatch []byte

func TestMarkets_ApplyPatch(t *testing.T) {
	testCases := []struct {
		name     string
		markets  market.Markets
		patches  map[string]json.RawMessage
		expected market.Markets
	}{
		{
			name: "patch existent market",
			markets: market.Markets{
				"m1": {
					ID:       "m1",
					Template: "template1",
					Status:   1,
					Odds: market.Odds{
						"o1": {
							ID:           "o1",
							Template:     "template1",
							IsActive:     true,
							Status:       1,
							Value:        "1",
							StatusReason: "reason1",
						},
					},
					TypeID:      1,
					Specifiers:  map[string]string{"1": "1"},
					IsDefective: false,
					Meta:        map[string]interface{}{"1": "1"},
					Flags:       1,
				},
			},
			patches: map[string]json.RawMessage{
				"m1/name":           json.RawMessage(`"template2"`),
				"m1/status":         json.RawMessage(`2`),
				"m1/odds/o2":        json.RawMessage(`{"id": "o2", "template": "template2", "is_active": true, "status": 2, "value": "2", "status_reason": "reason2"}`),
				"m1/odds/o1/status": json.RawMessage(`2`),
				"m1/type_id":        json.RawMessage(`2`),
			},
			expected: market.Markets{
				"m1": {
					ID:       "m1",
					Template: "template2",
					Status:   2,
					Odds: market.Odds{
						"o1": {
							ID:           "o1",
							Template:     "template1",
							IsActive:     true,
							Status:       2,
							Value:        "1",
							StatusReason: "reason1",
						},
						"o2": {
							ID:           "o2",
							Template:     "template2",
							IsActive:     true,
							Status:       2,
							Value:        "2",
							StatusReason: "reason2",
						},
					},
					TypeID:      2,
					Specifiers:  map[string]string{"1": "1"},
					IsDefective: false,
					Meta:        map[string]interface{}{"1": "1"},
					Flags:       1,
				},
			},
		},
		{
			name:    "patch new market",
			markets: market.Markets{},
			patches: map[string]json.RawMessage{
				"m1": marketPatch,
			},
			expected: market.Markets{
				"m1": {
					ID:       "1",
					Template: "Winner",
					Status:   2,
					Odds: market.Odds{
						"1": {
							ID:           "1",
							Template:     "{$competitor1}",
							IsActive:     true,
							Status:       1,
							Value:        "1.88",
							StatusReason: "reason1",
						},
						"2": {
							ID:           "2",
							Template:     "{$competitor2}",
							IsActive:     true,
							Status:       1,
							Value:        "2.09",
							StatusReason: "reason2",
						},
					},
					TypeID:      1,
					Specifiers:  map[string]string{"1": "2"},
					IsDefective: false,
					Meta:        map[string]interface{}{"1": "1"},
					Flags:       2,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for path, value := range tc.patches {
				err := tc.markets.ApplyPatch(path, value)
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expected, tc.markets)
		})
	}
}
