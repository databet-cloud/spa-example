//go:generate go run github.com/mailru/easyjson/easyjson fixture.go
package fixture

import (
	"time"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

func NewFixture() *Fixture {
	return &Fixture{
		Competitors: Competitors{},
		Streams:     Streams{},
		Venue:       Venue{},
		Tournament:  Tournament{},
	}
}

// easyjson:json
type Fixture struct {
	ID           string      `json:"id"`
	Version      int         `json:"version"`
	OwnerID      string      `json:"owner_id"`
	Template     string      `json:"template"`
	Status       int         `json:"status"`
	Type         int         `json:"type"`
	SportID      string      `json:"sport_id"`
	Tournament   Tournament  `json:"tournament"`
	Venue        Venue       `json:"venue"`
	Competitors  Competitors `json:"competitors"`
	Streams      Streams     `json:"streams"`
	LiveCoverage bool        `json:"live_coverage"`
	StartTime    time.Time   `json:"start_time"`
	Flags        int         `json:"flags"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	PublishedAt  time.Time   `json:"published_at"`
}

func (f Fixture) WithPatch(patchTree patch.Tree) Fixture {
	if v, ok := patch.GetFromTree[int](patchTree, "status"); ok {
		f.Status = v
	}

	if v, ok := patch.GetFromTree[int](patchTree, "type"); ok {
		f.Type = v
	} else if v, ok := patch.GetFromTree[float64](patchTree, "type"); ok {
		f.Type = int(v)
	}

	if v, ok := patch.GetFromTree[string](patchTree, "start_time"); ok {
		startTime, err := time.Parse(time.RFC3339, v)
		if err == nil {
			f.StartTime = startTime
		}
	}

	if v, ok := patch.GetFromTree[bool](patchTree, "live_coverage"); ok {
		f.LiveCoverage = v
	}

	f.Tournament = f.Tournament.WithPatch(patchTree.SubTree("tournament"))
	f.Venue = f.Venue.WithPatch(patchTree.SubTree("venue"))

	if subTree := patchTree.SubTree("streams"); !subTree.Empty() {
		f.Streams = patch.MapPatchable(f.Streams, subTree)
	}

	if subTree := patchTree.SubTree("competitors"); !subTree.Empty() {
		f.Competitors = patch.MapPatchable(f.Competitors, subTree)
	}

	return f
}

func (f Fixture) Clone() Fixture {
	result := f
	result.Competitors = f.Competitors.Clone()

	return result
}
