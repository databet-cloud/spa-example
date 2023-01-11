package sportevent_test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/minio/simdjson-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
	"github.com/databet-cloud/databet-go-sdk/pkg/feed"
	"github.com/databet-cloud/databet-go-sdk/pkg/feed/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/feed/market"
	"github.com/databet-cloud/databet-go-sdk/pkg/feed/sportevent"
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
	b.ReportAllocs()

	b.Run("with copy strings", func(b *testing.B) {
		b.StopTimer()

		var sportEvent sportevent.SportEvent

		err := json.Unmarshal(rawSportEvent, &sportEvent)
		require.NoError(b, err)

		patcher := sportevent.NewPatcherSimdJSON()

		b.StartTimer()
		b.ReportAllocs()

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
	})

	b.Run("without copy strings", func(b *testing.B) {
		b.StopTimer()

		var sportEvent sportevent.SportEvent

		err := json.Unmarshal(rawSportEvent, &sportEvent)
		require.NoError(b, err)

		patcher := sportevent.NewPatcherSimdJSON(simdjson.WithCopyStrings(false))

		b.StartTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			decoder := json.NewDecoder(bytes.NewReader(rawLogs))
			for decoder.More() {
				var log feed.LogEntry

				err := decoder.Decode(&log)
				require.NoError(b, err)

				err = patcher.ApplyPatches(&sportEvent, log.Patches)
				require.NoError(b, err)
			}
		}
	})
}

func BenchmarkSportEvent_Unmarshal(b *testing.B) {
	b.Run("simdjson", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var sportEvent sportevent.SportEventLazy

			err := sportEvent.UnmarshalJSON(rawSportEventWithMarkets)
			require.NoError(b, err)

			markets, err := market.MarketsFromMarketIter(sportEvent.MarketIter)
			require.NoError(b, err)

			_ = markets
		}
	})

	b.Run("simdjson with reuse", func(b *testing.B) {
		b.ReportAllocs()

		reuseParsedJSON := new(simdjson.ParsedJson)
		rootObj := new(simdjson.Object)
		reuseIter := new(simdjson.Iter)
		reuseObj := new(simdjson.Object)
		fixtureObj := new(simdjson.Object)
		competitorObj := new(simdjson.Object)
		competitor := new(fixture.Competitor)
		scoresObj := new(simdjson.Object)
		scoreObj := new(simdjson.Object)
		score := new(fixture.Score)

		for i := 0; i < b.N; i++ {
			var sportEvent sportevent.SportEventLazy

			parsedJson, err := simdjson.Parse(rawSportEventWithMarkets, reuseParsedJSON, simdjson.WithCopyStrings(false))
			require.NoError(b, err)

			rootIter, err := simdutil.CreateRootIter(parsedJson)
			require.NoError(b, err)

			rootObj, err := rootIter.Object(rootObj)
			require.NoError(b, err)

			err = sportEvent.UnmarshalSimdJSON(
				rootObj,
				reuseIter,
				reuseObj,
				fixtureObj,
				competitorObj,
				competitor,
				scoresObj,
				scoreObj,
				score,
			)
			require.NoError(b, err)

			markets, err := market.MarketsFromMarketIter(sportEvent.MarketIter)
			require.NoError(b, err)

			_ = markets
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

	patcher := sportevent.NewPatcherSimdJSON(simdjson.WithCopyStrings(false))

	decoder := json.NewDecoder(bytes.NewReader(rawLogs))
	for decoder.More() {
		var log feed.LogEntry

		err := decoder.Decode(&log)
		require.NoError(t, err)

		err = patcher.ApplyPatches(&sportEvent, log.Patches)
		assert.NoError(t, err)
	}

	// Convert ints to floats, because various libs unmarshal num to any different
	for k, v := range sportEvent.Meta {
		if num, ok := v.(int64); ok {
			sportEvent.Meta[k] = float64(num)
		}
	}

	assert.Equal(t, expectedSportEvent, sportEvent)
}

func TestSportEventLazy_UnmarshalJSON(t *testing.T) {
	var sportEvent sportevent.SportEventLazy
	var expected sportevent.SportEvent

	err := json.Unmarshal(rawSportEventWithMarkets, &expected)
	require.NoError(t, err)

	err = sportEvent.UnmarshalJSON(rawSportEventWithMarkets)
	require.NoError(t, err)

	markets, err := market.MarketsFromMarketIter(sportEvent.MarketIter)
	require.NoError(t, err)

	actual := sportevent.SportEvent{
		ID:        sportEvent.ID,
		Meta:      sportEvent.Meta,
		Fixture:   sportEvent.Fixture,
		Markets:   markets,
		BetStop:   sportEvent.BetStop,
		UpdatedAt: sportEvent.UpdatedAt,
	}

	// Convert ints to floats, because various libs unmarshal num to any different
	for k, v := range actual.Meta {
		if num, ok := v.(int64); ok {
			actual.Meta[k] = float64(num)
		}
	}

	assert.Equal(t, expected, actual)
}
