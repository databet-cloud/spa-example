package sportevent

import (
	"time"

	"golang.org/x/exp/maps"

	"github.com/databet-cloud/databet-go-sdk/pkg/feed/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/feed/market"
)

type SportEvent struct {
	ID        string          `json:"id"`
	Meta      map[string]any  `json:"meta"`
	Fixture   fixture.Fixture `json:"fixture"`
	Markets   market.Markets  `json:"markets"`
	BetStop   bool            `json:"bet_stop"`
	UpdatedAt time.Time       `json:"updated_at"`
}

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
