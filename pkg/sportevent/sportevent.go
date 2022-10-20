//go:generate go run github.com/mailru/easyjson/easyjson sportevent.go
package sportevent

import (
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

func (se SportEvent) WithPatch(tree patch.Tree) (SportEvent, error) {
	if v, ok := patch.GetFromTree[time.Time](tree, "updated_at"); ok {
		se.UpdatedAt = v
		se.Fixture.UpdatedAt = v
	}

	if v, ok := patch.GetFromTree[bool](tree, "bet_stop"); ok {
		se.BetStop = v
	}

	if subTree := tree.SubTree("markets"); !subTree.Empty() {
		se.Markets = patch.MapPatchable(se.Markets, subTree)
	}

	if subTree := tree.SubTree("meta"); !subTree.Empty() {
		se.Meta = patch.PatchMap(se.Meta, subTree)
	}

	se.Fixture = se.Fixture.WithPatch(tree.SubTree("fixture"))

	return se, nil
}
