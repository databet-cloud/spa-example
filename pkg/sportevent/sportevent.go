//go:generate go run github.com/mailru/easyjson/easyjson sportevent.go
package sportevent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"golang.org/x/exp/maps"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/market"
)

//easyjson:json
type SportEvent struct {
	ID         string                 `json:"id"`
	Meta       map[string]interface{} `json:"meta"`
	Fixture    fixture.Fixture        `json:"fixture"`
	MarketIter *market.Iterator
	Markets    market.Markets `json:"markets"`
	BetStop    bool           `json:"bet_stop"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func (se *SportEvent) ApplyPatchesV1(rawPatches json.RawMessage) error {
	var patches map[string]json.RawMessage

	err := json.Unmarshal(rawPatches, &patches)
	if err != nil {
		return fmt.Errorf("unmarshal patches: %w", err)
	}

	for path, patch := range patches {
		err := se.applyPatch(path, patch)
		if err != nil {
			return err
		}
	}

	return nil
}

func (se *SportEvent) applyPatch(path string, value json.RawMessage) error {
	var (
		unmarshaller     any
		key, rest, found = strings.Cut(path, "/")
		partialPatch     = found
	)

	switch key {
	case "bet_stop":
		unmarshaller = &se.BetStop
	case "updated_at":
		unmarshaller = &se.UpdatedAt
	case "fixture":
		if partialPatch {
			return se.Fixture.ApplyPatch(rest, value)
		}

		unmarshaller = &se.Fixture
	case "markets":
		if partialPatch {
			return se.Markets.ApplyPatch(rest, value)
		}

		unmarshaller = &se.Markets
	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}

func (se *SportEvent) Clone() *SportEvent {
	return &SportEvent{
		ID:         se.ID,
		Meta:       maps.Clone(se.Meta),
		Fixture:    se.Fixture.Clone(),
		MarketIter: se.MarketIter,
		Markets:    se.Markets.Clone(),
		BetStop:    se.BetStop,
		UpdatedAt:  se.UpdatedAt,
	}
}
