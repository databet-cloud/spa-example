package sportevent_test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/databet-cloud/databet-go-sdk/pkg/feed"
	"github.com/databet-cloud/databet-go-sdk/pkg/sportevent"
)

//go:embed testdata/benchmark/sport_event.json
var rawSportEvent []byte

//go:embed testdata/benchmark/feed_logs.json
var rawLogs []byte

func BenchmarkSportEventApplyPatch(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	var sportEvent sportevent.SportEvent

	err := json.Unmarshal(rawSportEvent, &sportEvent)
	require.NoError(b, err)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var log feed.LogEntry

		decoder := json.NewDecoder(bytes.NewReader(rawLogs))
		for decoder.More() {
			err := decoder.Decode(&log)
			require.NoError(b, err)

			for path, value := range log.Patches {
				err = sportEvent.ApplyPatch(path, value)
				require.NoError(b, err)
			}
		}
	}
}
