//go:generate go run github.com/mailru/easyjson/easyjson fixture.go
package fixture

import (
	"fmt"
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

type FixturePatch struct {
	Status       *int      `mapstructure:"status"`
	Type         *int      `mapstructure:"type"`
	StartTime    time.Time `mapstructure:"start_time"`
	LiveCoverage *bool     `mapstructure:"live_coverage"`
}

func (f Fixture) WithPatch(patchTree patch.Tree) (Fixture, error) {
	var fixturePatch FixturePatch

	err := patchTree.UnmarshalPatch(&fixturePatch)
	if err != nil {
		return Fixture{}, fmt.Errorf("unmarshal fixture patch: %w", err)
	}

	if fixturePatch.Status != nil {
		f.Status = *fixturePatch.Status
	}

	if fixturePatch.Type != nil {
		f.Type = *fixturePatch.Type
	}

	if !fixturePatch.StartTime.IsZero() {
		f.StartTime = fixturePatch.StartTime
	}

	if fixturePatch.LiveCoverage != nil {
		f.LiveCoverage = *fixturePatch.LiveCoverage
	}

	f.Tournament, err = f.Tournament.WithPatch(patchTree.SubTree("tournament"))
	if err != nil {
		return Fixture{}, fmt.Errorf("patch tournament: %w", err)
	}

	f.Venue = f.Venue.WithPatch(patchTree.SubTree("venue"))

	if subTree := patchTree.SubTree("streams"); !subTree.Empty() {
		f.Streams, err = patch.MapPatchable(f.Streams, subTree)
		if err != nil {
			return Fixture{}, fmt.Errorf("patch streams: %w", err)
		}
	}

	if subTree := patchTree.SubTree("competitors"); !subTree.Empty() {
		f.Competitors, err = patch.MapPatchable(f.Competitors, subTree)
		if err != nil {
			return Fixture{}, fmt.Errorf("patch competitors: %w", err)
		}
	}

	return f, nil
}

func (f Fixture) Clone() Fixture {
	result := f
	result.Competitors = f.Competitors.Clone()

	return result
}
