//go:generate go run github.com/mailru/easyjson/easyjson odd.go
package market

import (
	"fmt"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

const (
	OddStatusNotResulted OddStatus = 0
	OddStatusWin         OddStatus = 1
	OddStatusLoss        OddStatus = 2
	OddStatusHalfWin     OddStatus = 3
	OddStatusHalfLoss    OddStatus = 4
	OddStatusRefunded    OddStatus = 5
	OddStatusCancelled   OddStatus = 6
)

type OddStatus int

// nolint:gochecknoglobals // package dictionary of odd statuses
var oddStatuses = []OddStatus{
	OddStatusNotResulted,
	OddStatusWin,
	OddStatusLoss,
	OddStatusHalfWin,
	OddStatusHalfLoss,
	OddStatusRefunded,
	OddStatusCancelled,
}

func (os OddStatus) IsValid() bool {
	for _, status := range oddStatuses {
		if os == status {
			return true
		}
	}

	return false
}

func (os OddStatus) IsResulted() bool {
	return os != OddStatusNotResulted
}

//easyjson:json
type Odds map[string]Odd

func (c Odds) Equals(other Odds) bool {
	if len(c) != len(other) {
		return false
	}

	for id, odd := range c {
		otherOdd, ok := other[id]
		if !ok {
			return false
		}

		if !odd.Equals(otherOdd) {
			return false
		}
	}

	return true
}

func (c Odds) Clone() Odds {
	newC := make(Odds, len(c))
	for id, odd := range c {
		newC[id] = odd.Clone()
	}

	return newC
}

//easyjson:json
type Odd struct {
	ID           string    `json:"id"`
	Template     string    `json:"template"`
	IsActive     bool      `json:"is_active"`
	Status       OddStatus `json:"status"`
	Value        string    `json:"value"`
	StatusReason string    `json:"status_reason"`
}

type OddPatch struct {
	ID           *string    `mapstructure:"id"`
	Template     *string    `mapstructure:"template"`
	IsActive     *bool      `mapstructure:"is_active"`
	Status       *OddStatus `mapstructure:"status"`
	Value        *string    `mapstructure:"value"`
	StatusReason *string    `mapstructure:"status_reason"`
}

func (o Odd) WithPatch(tree patch.Tree) (Odd, error) {
	var oddPatch OddPatch

	err := tree.UnmarshalPatch(&oddPatch)
	if err != nil {
		return Odd{}, fmt.Errorf("unmarshal odd patch: %w", err)
	}

	if oddPatch.ID != nil {
		o.ID = *oddPatch.ID
	}

	if oddPatch.Template != nil {
		o.Template = *oddPatch.Template
	}

	if oddPatch.IsActive != nil {
		o.IsActive = *oddPatch.IsActive
	}

	if oddPatch.Status != nil {
		o.Status = *oddPatch.Status
	}

	if oddPatch.Value != nil {
		o.Value = *oddPatch.Value
	}

	if oddPatch.StatusReason != nil {
		o.StatusReason = *oddPatch.StatusReason
	}

	return o, nil
}

func (o Odd) Equals(other Odd) bool {
	return o.ID == other.ID &&
		o.Template == other.Template &&
		o.IsActive == other.IsActive &&
		o.Status == other.Status &&
		o.Value == other.Value &&
		o.StatusReason == other.StatusReason
}

func (o Odd) Clone() Odd {
	return Odd{
		ID:           o.ID,
		Template:     o.Template,
		IsActive:     o.IsActive,
		Status:       o.Status,
		Value:        o.Value,
		StatusReason: o.StatusReason,
	}
}
