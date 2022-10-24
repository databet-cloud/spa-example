//go:generate go run github.com/mailru/easyjson/easyjson sportevent.go
package sportevent

import (
	"fmt"
	"time"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/market"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

//easyjson:json
type SportEvent struct {
	ID        string                 `json:"id"`
	Meta      map[string]interface{} `json:"meta"`
	Fixture   fixture.Fixture        `json:"fixture"`
	Markets   market.Markets         `json:"markets"`
	BetStop   bool                   `json:"bet_stop"`
	Sources   []Source               `json:"sources"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type SportEventPatch struct {
	BetStop   *bool     `mapstructure:"bet_stop"`
	UpdatedAt time.Time `mapstructure:"updated_at"`
}

func (se SportEvent) WithPatch(tree patch.Tree) (SportEvent, error) {
	var sportEventPatch SportEventPatch

	err := tree.UnmarshalPatch(&sportEventPatch)
	if err != nil {
		return SportEvent{}, fmt.Errorf("unmarshal sport event patch: %w", err)
	}

	if sportEventPatch.BetStop != nil {
		se.BetStop = *sportEventPatch.BetStop
	}

	if !sportEventPatch.UpdatedAt.IsZero() {
		se.UpdatedAt = sportEventPatch.UpdatedAt
	}

	if subTree := tree.SubTree("markets"); !subTree.Empty() {
		se.Markets, err = patch.MapPatchable(se.Markets, subTree)
		if err != nil {
			return SportEvent{}, fmt.Errorf("patch markets: %w", err)
		}
	}

	if subTree := tree.SubTree("meta"); !subTree.Empty() {
		se.Meta = patch.PatchMap(se.Meta, subTree)
	}

	se.Fixture, err = se.Fixture.WithPatch(tree.SubTree("fixture"))
	if err != nil {
		return SportEvent{}, fmt.Errorf("patch fixture: %w", err)
	}

	return se, nil
}

func (se *SportEvent) ApplyPatch(tree patch.Tree) error {
	var sportEventPatch SportEventPatch

	err := tree.UnmarshalPatch(&sportEventPatch)
	if err != nil {
		return fmt.Errorf("unmarshal sport event patch: %w", err)
	}

	if sportEventPatch.BetStop != nil {
		se.BetStop = *sportEventPatch.BetStop
	}

	if !sportEventPatch.UpdatedAt.IsZero() {
		se.UpdatedAt = sportEventPatch.UpdatedAt
	}

	if subTree := tree.SubTree("markets"); !subTree.Empty() {
		se.Markets, err = patch.MapPatchable(se.Markets, subTree)
		if err != nil {
			return fmt.Errorf("patch markets: %w", err)
		}
	}

	if subTree := tree.SubTree("meta"); !subTree.Empty() {
		se.Meta = patch.PatchMap(se.Meta, subTree)
	}

	err = se.Fixture.ApplyPatch(tree.SubTree("fixture"))
	if err != nil {
		return fmt.Errorf("patch fixture: %w", err)
	}

	return nil
}
