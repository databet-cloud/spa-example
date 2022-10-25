//go:generate go run github.com/mailru/easyjson/easyjson fixture.go
package fixture

import (
	"encoding/json"
	"fmt"
	"strings"
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

func (f *Fixture) ApplyPatch(path string, value json.RawMessage) error {
	var (
		unmarshaller     any
		key, rest, found = strings.Cut(path, "/")
		partialPatch     = found
	)

	switch key {
	case "status":
		unmarshaller = &f.Status
	case "type":
		unmarshaller = &f.Type
	case "start_time":
		unmarshaller = &f.StartTime
	case "live_coverage":
		unmarshaller = &f.LiveCoverage
	case "competitors":
		if partialPatch {
			if f.Competitors == nil {
				return fmt.Errorf("patch nil competitors")
			}

			return f.Competitors.ApplyPatch(rest, value)
		}

		unmarshaller = &f.Competitors
	case "tournament":
		if partialPatch {
			return f.Tournament.ApplyPatch(rest, value)
		}

		unmarshaller = &f.Tournament
	case "streams":
		if partialPatch {
			if f.Streams == nil {
				return fmt.Errorf("patch nil streams")
			}

			return f.Streams.ApplyPatch(rest, value)
		}

		unmarshaller = &f.Streams
	case "venue":
		if partialPatch {
			return f.Venue.ApplyPatch(rest, value)
		}

		unmarshaller = &f.Venue
	default:
		return nil
	}

	err := json.Unmarshal(value, &unmarshaller)
	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", path, err)
	}

	return nil
}

func (f Fixture) Clone() Fixture {
	result := f
	result.Competitors = f.Competitors.Clone()

	return result
}
