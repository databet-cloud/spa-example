//go:generate go run github.com/mailru/easyjson/easyjson market.go
package market

import (
	"strconv"

	"github.com/mailru/easyjson/jwriter"
	"golang.org/x/exp/maps"
)

const (
	StatusActive      Status = 0
	StatusSuspended   Status = 1
	StatusDeactivated Status = 2
	StatusResulted    Status = 3
	StatusCancelled   Status = 4

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

//easyjson:json
type Collection map[string]Market

func (c Collection) Suspended() Collection {
	res := make(Collection, len(c))

	for mID, m := range c {
		if m.Status == StatusActive {
			m.Status = StatusSuspended
		}

		res[mID] = m
	}

	return res
}

func (c Collection) Has(id string) bool {
	_, ok := c[id]

	return ok
}

func (c Collection) ToSlice() []Market {
	res := make([]Market, 0, len(c))
	for _, m := range c {
		res = append(res, m)
	}

	return res
}

func (c Collection) Clone() Collection {
	newC := make(Collection, len(c))
	for id, market := range c {
		newC[id] = market.Clone()
	}

	return newC
}

type Markets map[string]Market

type Market struct {
	ID         string                 `json:"id" bson:"id"`
	TypeID     int                    `json:"type_id" bson:"type_id"`
	Template   string                 `json:"template" bson:"template"`
	Status     Status                 `json:"status" bson:"status"`
	Odds       OddCollection          `json:"odds" bson:"odds"`
	Specifiers map[string]string      `json:"specifiers" bson:"specifiers"`
	Meta       map[string]interface{} `json:"meta" bson:"meta"`
	Flags      int                    `json:"flags" bson:"flags"`
}

//easyjson:json
type jsonMarket struct {
	ID              string                 `json:"id"`
	TypeID          int                    `json:"type_id"`
	Template        string                 `json:"template"`
	Status          Status                 `json:"status"`
	Odds            OddCollection          `json:"odds"`
	Specifiers      map[string]string      `json:"specifiers"`
	Meta            map[string]interface{} `json:"meta"`
	Flags           int                    `json:"flags"`
	IsRobotDetached bool                   `json:"is_robot_detached"`
	IsDefective     bool                   `json:"is_defective"`
}

func (m Market) makeJSONMarket() jsonMarket {
	return jsonMarket{
		ID:              m.ID,
		TypeID:          m.TypeID,
		Template:        m.Template,
		Status:          m.Status,
		Odds:            m.Odds,
		Specifiers:      m.Specifiers,
		Meta:            m.Meta,
		Flags:           m.Flags,
		IsRobotDetached: m.IsRobotDetached(),
		IsDefective:     m.IsDefective(),
	}
}

func (m Market) MarshalEasyJSON(w *jwriter.Writer) {
	initial := w.Flags
	w.Flags = jwriter.NilMapAsEmpty

	m.makeJSONMarket().MarshalEasyJSON(w)

	w.Flags = initial
}

func (m Market) MarshalJSON() ([]byte, error) {
	return m.makeJSONMarket().MarshalJSON()
}

func (m Market) Clone() Market {
	return Market{
		ID:         m.ID,
		TypeID:     m.TypeID,
		Template:   m.Template,
		Status:     m.Status,
		Odds:       m.Odds.Clone(),
		Specifiers: maps.Clone(m.Specifiers),
		Meta:       maps.Clone(m.Meta),
		Flags:      m.Flags,
	}
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

func (m Market) IsDefective() bool {
	return (m.Flags & IsDefective) != 0
}

func (m Market) IsRobotDetached() bool {
	return (m.Flags & IsRobotDetached) != 0
}
