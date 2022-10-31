//go:generate go run github.com/mailru/easyjson/easyjson -no_std_marshalers odd.go
package market

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

const (
	OddStatusNotResulted OddStatus = iota
	OddStatusWin
	OddStatusLoss
	OddStatusHalfWin
	OddStatusHalfLoss
	OddStatusRefunded
	OddStatusCancelled
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

func (c Odds) UnmarshalSimdJSON(obj *simdjson.Object, reuseIter *simdjson.Iter, reuseOddObj *simdjson.Object, reuseOdd *Odd) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
	}

	if reuseOddObj == nil {
		reuseOddObj = new(simdjson.Object)
	}

	if reuseOdd == nil {
		reuseOdd = new(Odd)
	}

	for {
		name, elementType, err := obj.NextElement(reuseIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		oddObj, err := reuseIter.Object(reuseOddObj)
		if err != nil {
			return fmt.Errorf("create %q object: %w", name, err)
		}

		err = reuseOdd.UnmarshalSimdJSON(oddObj, reuseIter)
		if err != nil {
			return fmt.Errorf("unmarshal %q odd: %w", name, err)
		}

		c[name] = *reuseOdd
	}

	return nil
}

func (c Odds) ApplyPatch(path string, value json.RawMessage) error {
	key, rest, found := strings.Cut(path, "/")
	odd, ok := c[key]

	if !found {
		err := json.Unmarshal(value, &odd)
		if err != nil {
			return fmt.Errorf("odd %q unmarshal: %w", key, err)
		}

		c[key] = odd
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent odd: %q", key)
	}

	err := odd.ApplyPatch(rest, value)
	if err != nil {
		return fmt.Errorf("apply odd patch: %w", err)
	}

	c[key] = odd
	return nil
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

func (o *Odd) ApplyPatch(path string, value json.RawMessage) error {
	var unmarshaller any

	switch path {
	case "name":
		unmarshaller = &o.Template
	case "value":
		unmarshaller = &o.Value
	case "is_active":
		unmarshaller = &o.IsActive
	case "status":
		unmarshaller = &o.Status
	case "status_reason":
		unmarshaller = &o.StatusReason
	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", path, err)
	}

	return nil
}

func (o *Odd) UnmarshalSimdJSON(obj *simdjson.Object, reuseIter *simdjson.Iter) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
	}

	for {
		name, elementType, err := obj.NextElementBytes(reuseIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		switch string(name) {
		case "id":
			o.ID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "template":
			o.Template, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "is_active":
			o.IsActive, err = reuseIter.Bool()
		case "status":
			var value int64
			value, err = reuseIter.Int()
			o.Status = OddStatus(value)
		case "value":
			o.Value, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "status_reason":
			o.StatusReason, err = reuseIter.String()
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
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
