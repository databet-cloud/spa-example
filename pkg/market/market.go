//go:generate go run github.com/mailru/easyjson/easyjson market.go
package market

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/minio/simdjson-go"
	"golang.org/x/exp/maps"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
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

//easyjson:json
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

func (c Markets) ApplyPatch(path string, value json.RawMessage) error {
	key, rest, found := strings.Cut(path, "/")
	market, ok := c[key]

	if !found {
		err := json.Unmarshal(value, &market)
		if err != nil {
			return fmt.Errorf("market %q unmarshal: %w", key, err)
		}

		c[key] = market
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent market: %q", key)
	}

	err := market.ApplyPatch(rest, value)
	if err != nil {
		return fmt.Errorf("apply market patch: %w", err)
	}

	c[key] = market
	return nil
}

func (c Markets) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	market, ok := c[key]

	if !partialPatch {
		obj, err := iter.Object(nil)
		if err != nil {
			return err
		}

		err = market.UnmarshalSimdJSON(obj)
		if err != nil {
			return fmt.Errorf("market %q unmarshal simdjson: %w", key, err)
		}

		c[key] = market
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent market: %q", key)
	}

	err := market.ApplyPatchSimdJSON(rest, iter)
	if err != nil {
		return fmt.Errorf("apply market patch: %w", err)
	}

	c[key] = market
	return nil
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

func (m *Market) UnmarshalSimdJSON(obj *simdjson.Object) error {
	tmpIter := new(simdjson.Iter)
	tmpObj := new(simdjson.Object)

	for {
		name, elementType, err := obj.NextElement(tmpIter)
		if err != nil {
			return fmt.Errorf("next element: %w", err)
		}

		if elementType == simdjson.TypeNone {
			break
		}

		switch name {
		case "id":
			m.ID, err = tmpIter.String()
		case "status":
			var value int64

			value, err = tmpIter.Int()
			m.Status = Status(value)
		case "type_id":
			m.TypeID, err = simdutil.IntFromIter(tmpIter)
		case "template":
			m.Template, err = tmpIter.String()
		case "flags":
			m.Flags, err = simdutil.IntFromIter(tmpIter)
		case "is_defective":
			m.IsDefective, err = tmpIter.Bool()
		case "specifiers":
			m.Specifiers, err = simdutil.MapStrStrFromIter(tmpIter)
		case "odds":
			m.Odds = make(Odds, 4)
			oddsObj, err := tmpIter.Object(tmpObj)
			if err != nil {
				return fmt.Errorf("create %q object: %w", name, err)
			}

			err = m.Odds.UnmarshalSimdJSON(oddsObj)
		case "meta":
			m.Meta, err = simdutil.MapStrAnyFromIter(tmpIter)
		default:
			continue
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}

func (m *Market) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	var (
		err                     error
		key, rest, partialPatch = strings.Cut(path, "/")
	)

	switch key {
	case "name":
		m.Template, err = iter.String()
	case "status":
		var value int64

		value, err = iter.Int()
		m.Status = Status(value)
	case "type_id":
		m.TypeID, err = simdutil.IntFromIter(iter)
	case "odds":
		if !partialPatch {
			obj, err := iter.Object(nil)
			if err != nil {
				return fmt.Errorf("create %q object: %w", key, err)
			}

			m.Odds = make(Odds)
			return m.Odds.UnmarshalSimdJSON(obj)
		}

		if m.Odds == nil {
			return fmt.Errorf("patch nil odds")
		}

		return m.Odds.ApplyPatchSimdJSON(rest, iter)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}

func (m *Market) ApplyPatch(path string, value json.RawMessage) error {
	var (
		unmarshaller     any
		key, rest, found = strings.Cut(path, "/")
		partialPatch     = found
	)

	switch key {
	case "name":
		unmarshaller = &m.Template
	case "status":
		unmarshaller = &m.Status
	case "type_id":
		unmarshaller = &m.TypeID
	case "odds":
		if partialPatch {
			if m.Odds == nil {
				return fmt.Errorf("patch nil odds")
			}

			return m.Odds.ApplyPatch(rest, value)
		}

		unmarshaller = &m.Odds
	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
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
