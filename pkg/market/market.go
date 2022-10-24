//go:generate go run github.com/mailru/easyjson/easyjson market.go
package market

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/maps"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
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
type Markets map[string]Market

func (c Markets) Clone() Markets {
	res := make(Markets, len(c))

	for id, market := range c {
		res[id] = market.Clone()
	}

	return res
}

func (c Markets) Suspended() Markets {
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

func (c Markets) ToSlice() []Market {
	res := make([]Market, 0, len(c))

	for _, m := range c {
		res = append(res, m)
	}

	return res
}

//easyjson:json
type Market struct {
	ID          string                 `json:"id"`
	Template    string                 `json:"template"`
	Status      Status                 `json:"status"`
	Odds        Odds                   `json:"odds"`
	TypeID      int                    `json:"type_id"`
	Specifiers  map[string]string      `json:"specifiers"`
	IsDefective bool                   `json:"is_defective"`
	Meta        map[string]interface{} `json:"meta"`
	Flags       int                    `json:"flags"`
}

type MarketPatch struct {
	ID          *string `mapstructure:"id"`
	Template    *string `mapstructure:"template"`
	Status      *Status `mapstructure:"status"`
	TypeID      *int    `mapstructure:"type_id"`
	IsDefective *bool   `mapstructure:"is_defective"`
	Flags       *int    `mapstructure:"flags"`
}

func (m Market) WithPatch(tree patch.Tree) (Market, error) {
	var marketPatch MarketPatch

	err := tree.UnmarshalPatch(&marketPatch)
	if err != nil {
		return Market{}, fmt.Errorf("unmarshal patch: %w", err)
	}

	m.applyMarketPatch(marketPatch)

	if subTree := tree.SubTree("odds"); !subTree.Empty() {
		m.Odds, err = patch.MapPatchable(m.Odds, subTree)
		if err != nil {
			return Market{}, fmt.Errorf("patch odd: %w", err)
		}
	}

	if subTree := tree.SubTree("specifiers"); !subTree.Empty() {
		m.Specifiers = patch.PatchMap(m.Specifiers, subTree)
	}

	if subTree := tree.SubTree("meta"); !subTree.Empty() {
		m.Meta = patch.PatchMap(m.Meta, subTree)
	}

	return m, nil
}

func (m Market) applyMarketPatch(patch MarketPatch) {
	if patch.ID != nil {
		m.ID = *patch.ID
	}

	if patch.TypeID != nil {
		m.TypeID = *patch.TypeID
	}

	if patch.Template != nil {
		m.Template = *patch.Template
	}

	if patch.Status != nil {
		m.Status = *patch.Status
	}

	if patch.Flags != nil {
		m.Flags = *patch.Flags
	}

	if patch.IsDefective != nil {
		m.IsDefective = *patch.IsDefective
	}
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
