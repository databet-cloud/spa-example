package fixture

import (
	"time"
)

func NewFixture() *Fixture {
	return &Fixture{
		Competitors: Competitors{},
		Streams:     Streams{},
		Tournament:  Tournament{},
	}
}

type Type int

const (
	TypeMatch Type = iota
	TpyeOutright
)

type Fixture struct {
	ID      string `json:"id"`
	Version int    `json:"version"`

	// Template is the name of the current sport event with variables in it.
	// To replace them, you could use TemplatePosition in competitor
	Template string `json:"template"`
	Status   Status `json:"status"`
	Type     Type   `json:"type"`

	// SportID to which the sport event belongs
	SportID string `json:"sport_id"`

	// Tournament of the current sport event
	Tournament  Tournament  `json:"tournament"`
	Competitors Competitors `json:"competitors"`

	// Streams is the collection of available streams, binded to the current sport event
	Streams Streams `json:"streams"`

	// LiveCoverage indicates whether the sport event could cover live mode, or only prematch
	LiveCoverage bool      `json:"live_coverage"`
	StartTime    time.Time `json:"start_time"`
	Flags        int       `json:"flags"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	PublishedAt  time.Time `json:"published_at"`
}

func (f Fixture) Clone() Fixture {
	f.Competitors = f.Competitors.Clone()
	f.Streams = f.Streams.Clone()

	return f
}
