//go:generate go run github.com/mailru/easyjson/easyjson log.go
package feed

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
	"github.com/databet-cloud/databet-go-sdk/pkg/sportevent"
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

//easyjson:json
type LogEntry struct {
	Version      string    `json:"version"`
	SportEventID string    `json:"sport_event_id"`
	Type         LogType   `json:"type"`
	Timestamp    timestamp `json:"timestamp"`

	// Changes filled for LogTypeMatchUpdate only
	Changes patch.Patch `json:"changes,omitempty"`

	// SportEvent filled for LogTypeMatchNew only
	SportEvent sportevent.SportEvent `json:"sport_event,omitempty"`

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
