package fixture

import (
	"time"
)

func NewFixture() *Fixture {
	return &Fixture{
		Competitors: Competitors{},
		Streams:     Streams{},
		Venue:       Venue{},
		Tournament:  Tournament{},
	}
}

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

func (f Fixture) Clone() Fixture {
	f.Competitors = f.Competitors.Clone()
	f.Streams = f.Streams.Clone()

	return f
}
