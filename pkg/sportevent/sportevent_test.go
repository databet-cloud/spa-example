package sportevent_test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databet-cloud/databet-go-sdk/pkg/feed"
	"github.com/databet-cloud/databet-go-sdk/pkg/market"
	"github.com/databet-cloud/databet-go-sdk/pkg/sportevent"
)

//go:embed testdata/benchmark/sport_event.json
var rawSportEvent []byte

//go:embed testdata/benchmark/sport_event_with_markets.json
var rawSportEventWithMarkets []byte

//go:embed testdata/benchmark/feed_logs.json
var rawLogs []byte

//go:embed testdata/patched_sport_event.json
var patchedSportEvent []byte

func BenchmarkSportEventApplyPatchSimdJSON(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	var sportEvent sportevent.SportEvent

	err := json.Unmarshal(rawSportEvent, &sportEvent)
	require.NoError(b, err)

	patcher := sportevent.NewPatcherSimdJSON()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var log feed.LogEntry

		decoder := json.NewDecoder(bytes.NewReader(rawLogs))
		for decoder.More() {
			err := decoder.Decode(&log)
			require.NoError(b, err)

			err = patcher.ApplyPatches(&sportEvent, log.Patches)
			require.NoError(b, err)
		}
	}
}

func BenchmarkSportEvent_Unmarshal(b *testing.B) {
	b.Run("simdjson", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var sportEvent sportevent.SportEventLazy

			err := sportEvent.UnmarshalJSON(rawSportEventWithMarkets)
			require.NoError(b, err)

			sportEvent.Markets, err = market.MarketsFromMarketIter(sportEvent.MarketIter)
			require.NoError(b, err)
		}
	})

	b.Run("sonic", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var sportEvent sportevent.SportEvent

			err := sonic.Unmarshal(rawSportEventWithMarkets, &sportEvent)
			require.NoError(b, err)
		}
	})
}

func TestSportEventPatcher_ApplyPatches(t *testing.T) {
	var sportEvent sportevent.SportEvent
	var expectedSportEvent sportevent.SportEvent

	err := sonic.Unmarshal(rawSportEvent, &sportEvent)
	require.NoError(t, err)

	err = sonic.Unmarshal(patchedSportEvent, &expectedSportEvent)
	require.NoError(t, err)

	var log feed.LogEntry

	patcher := sportevent.NewPatcherSimdJSON()

	decoder := json.NewDecoder(bytes.NewReader(rawLogs))
	for decoder.More() {
		err := decoder.Decode(&log)
		require.NoError(t, err)

		err = patcher.ApplyPatches(&sportEvent, log.Patches)
		assert.NoError(t, err)
	}

	assert.Equal(t, expectedSportEvent, sportEvent)
}
