package market_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databet-cloud/databet-go-sdk/pkg/market"
	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

var (
	//go:embed testdata/markets1.json
	testMarkets1 []byte
)

func TestMarketIterator(t *testing.T) {
	testCases := []struct {
		name       string
		rawMarkets json.RawMessage
	}{
		{
			name:       "few markets",
			rawMarkets: testMarkets1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var expected market.Markets

			err := json.Unmarshal(tc.rawMarkets, &expected)
			require.NoError(t, err)
			//
			rootIter, err := simdutil.JSONToRootIter(tc.rawMarkets)
			require.NoError(t, err)

			marketsIter, err := market.NewIterator(rootIter)
			require.NoError(t, err)
			//
			actual, err := market.MarketsFromMarketIter(marketsIter)
			assert.NoError(t, err)

			// Ignore field Meta, because of different types in values (42 float / 42 int)
			assert.Empty(t, cmp.Diff(expected, actual, cmpopts.IgnoreFields(market.Market{}, "Meta")))
			//
			encoder := json.NewEncoder(os.Stdout)
			encoder.SetIndent("", "  ")
			err = encoder.Encode(actual)
			require.NoError(t, err)
		})
	}
}
