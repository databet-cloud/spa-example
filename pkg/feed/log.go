package feed

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type LogType string

func (l LogType) String() string {
	return string(l)
}

const (
	LogTypeMatchNew    LogType = "match_new"
	LogTypeMatchUpdate LogType = "match_update"
	LogTypeBetRollback LogType = "bet_rollback"
)

type LogEntry struct {
	Version      string    `json:"v"`
	SportEventID string    `json:"sport_event_id"`
	Type         LogType   `json:"type"`
	Timestamp    timestamp `json:"timestamp"`

	// Patches filled for LogTypeMatchUpdate only
	Patches json.RawMessage `json:"changes,omitempty"`

	// SportEvent filled for LogTypeMatchNew only
	SportEvent json.RawMessage `json:"sport_event,omitempty"`

	// The following fields are filled for LogTypeBetRollback only
	MatchID   string    `json:"match_id,omitempty"`
	MarketIDs []string  `json:"market_ids,omitempty"`
	DateStart time.Time `json:"dt_start,omitempty"`
	DateEnd   time.Time `json:"dt_end,omitempty"`
}

type timestamp time.Time

func (t *timestamp) UnmarshalJSON(bytes []byte) error {
	raw := strings.Trim(string(bytes), "\"")

	seconds, err := strconv.Atoi(raw)
	if err != nil {
		return fmt.Errorf("atoi: %w, raw: %q", err, raw)
	}

	*t = timestamp(time.Unix(int64(seconds), 0))

	return nil
}