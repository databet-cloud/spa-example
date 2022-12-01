package sportevent

import (
	"time"

	"golang.org/x/exp/maps"

	"github.com/databet-cloud/databet-go-sdk/pkg/feed/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/feed/market"
)

type SportEvent struct {
	// ID is the sport event's unique id
	ID string `json:"id"`
	// Meta contains additional information, such as prize pool, best of, half number, ...
	Meta map[string]any `json:"meta"`
	// Fixture contains all the necessary information about competitors/tournament/score/status/etc...
	Fixture fixture.Fixture `json:"fixture"`
	// Markets contain collection of all markets with their outcomes/odds/etc...
	Markets market.Markets `json:"markets"`
	// BetStop indicates whether you should stop accepting new bets or not
	BetStop   bool      `json:"bet_stop"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Clone sport event and return it's deep copy
func (se *SportEvent) Clone() *SportEvent {
	return &SportEvent{
		ID:        se.ID,
		Meta:      maps.Clone(se.Meta),
		Fixture:   se.Fixture.Clone(),
		Markets:   se.Markets.Clone(),
		BetStop:   se.BetStop,
		UpdatedAt: se.UpdatedAt,
	}
}
