package feed_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/feed"
)

var (
	//go:embed testdata/log_entry_match_new.json
	logEntryMatchNew []byte

	//go:embed testdata/log_entry_match_update.json
	logEntryMatchUpdate []byte

	//go:embed testdata/log_entry_bet_rollback.json
	logEntryBetRollback []byte
)

func TestLogEntryUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name        string
		rawLogEntry []byte
	}{
		{
			name:        "match_new",
			rawLogEntry: logEntryMatchNew,
		},
		{
			name:        "match_update",
			rawLogEntry: logEntryMatchUpdate,
		},
		{
			name:        "bet_rollback",
			rawLogEntry: logEntryBetRollback,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var logEntry feed.LogEntry

			err := json.Unmarshal(tc.rawLogEntry, &logEntry)
			assert.NoError(t, err)
		})
	}
}
