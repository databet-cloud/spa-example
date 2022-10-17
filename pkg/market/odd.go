//go:generate go run github.com/mailru/easyjson/easyjson odd.go
package market

import (
	"encoding/json"

	"github.com/mailru/easyjson/jwriter"
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
type OddCollection map[string]Odd

func (c OddCollection) Equals(other OddCollection) bool {
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

func (c OddCollection) Clone() OddCollection {
	newC := make(OddCollection, len(c))
	for id, odd := range c {
		newC[id] = odd.Clone()
	}

	return newC
}

type Odd struct {
	ID           string    `json:"id" bson:"id"`
	Template     string    `json:"template" bson:"template"`
	IsActive     bool      `json:"is_active" bson:"is_active"`
	Status       OddStatus `json:"status" bson:"status"`
	Value        float64   `json:"value,string" bson:"value"`
	Marge        float64   `json:"marge,string" bson:"marge"`
	StatusReason string    `json:"status_reason" bson:"status_reason"`
}

func (o Odd) Equals(other Odd) bool {
	return o.ID == other.ID &&
		o.Template == other.Template &&
		o.IsActive == other.IsActive &&
		o.Status == other.Status &&
		o.Value == other.Value &&
		o.Marge == other.Marge &&
		o.StatusReason == other.StatusReason
}

func (o Odd) Clone() Odd {
	return Odd{
		ID:           o.ID,
		Template:     o.Template,
		IsActive:     o.IsActive,
		Status:       o.Status,
		Value:        o.Value,
		Marge:        o.Marge,
		StatusReason: o.StatusReason,
	}
}

//easyjson:json
type oddJSON struct {
	ID           string          `json:"id"`
	Template     string          `json:"template"`
	IsActive     bool            `json:"is_active"`
	Status       OddStatus       `json:"status"`
	Value        float64         `json:"value,string"`
	Marge        float64         `json:"marge,string"`
	Meta         json.RawMessage `json:"meta"` // todo: Deprecated field (saved for BC), should be removed
	StatusReason string          `json:"status_reason"`
}

//nolint:gochecknoglobals // shared constant meta for BC
var emptyMeta = []byte{'{', '}'}

func (o Odd) makeJSONOdd() oddJSON {
	return oddJSON{
		ID:           o.ID,
		Template:     o.Template,
		IsActive:     o.IsActive,
		Status:       o.Status,
		Value:        o.Value,
		Marge:        o.Marge,
		Meta:         emptyMeta,
		StatusReason: o.StatusReason,
	}
}

func (o Odd) MarshalEasyJSON(w *jwriter.Writer) {
	o.makeJSONOdd().MarshalEasyJSON(w)
}

func (o Odd) MarshalJSON() ([]byte, error) {
	return o.makeJSONOdd().MarshalJSON()
}
