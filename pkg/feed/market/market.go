package market

import (
	"strconv"

	"golang.org/x/exp/maps"
)

const (
	StatusActive Status = iota
	StatusSuspended
	StatusDeactivated
	StatusResulted
	StatusCancelled

	IsDefective     = 1 << 0
	IsRobotDetached = 1 << 1

	MetadataManual = "manual"
)

type Status int

func (s Status) String() string {
	return strconv.Itoa(int(s))
}

func (s Status) IsValid() bool {
	switch s {
	case StatusActive,
		StatusSuspended,
		StatusDeactivated,
		StatusResulted,
		StatusCancelled:
		return true
	default:
		return false
	}
}

// Markets is a map with market id in its key and market as a value
type Markets map[string]Market

func (c Markets) Clone() Markets {
	res := make(Markets, len(c))

	for id, market := range c {
		res[id] = market.Clone()
	}

	return res
}

func (c Markets) Suspend() Markets {
	res := make(Markets, len(c))

	for mID, m := range c {
		if m.Status == StatusActive {
			m.Status = StatusSuspended
		}

		res[mID] = m
	}

	return res
}

func (c Markets) Has(id string) bool {
	_, ok := c[id]
	return ok
}

type Market struct {
	// ID of the market
	ID string `json:"id"`

	// Template is a name with specifiers, that could be replaced to format market's name
	Template string `json:"template"`
	Status   Status `json:"status"`

	// Odds is the collection of all available odds for the current market
	Odds Odds `json:"odds"`

	// TypeID is an original id for the market, without odds in it
	TypeID int `json:"type_id"`

	// Specifiers is the collection of all specifiers, such as
	Specifiers  map[string]string `json:"specifiers"`
	IsDefective bool              `json:"is_defective"`
	Meta        map[string]any    `json:"meta"`

	// Flags could be checked, using & operator, e.g.: flags & IsDefective
	Flags int `json:"flags"`
}

func (m Market) Clone() Market {
	m.Specifiers = maps.Clone(m.Specifiers)
	m.Meta = maps.Clone(m.Meta)
	m.Odds = m.Odds.Clone()

	return m
}

func (m Market) WithRobotDetached(robotDetached bool) Market {
	newM := m.Clone()
	if robotDetached {
		newM.Flags |= IsRobotDetached
	} else {
		newM.Flags &= ^IsRobotDetached
	}

	return newM
}

func (m Market) IsRobotDetached() bool {
	return (m.Flags & IsRobotDetached) != 0
}
