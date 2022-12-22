package market

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

func (m *Market) UnmarshalSimdJSON(
	obj *simdjson.Object,
	reuseIter *simdjson.Iter,
	reuseOddsObj *simdjson.Object,
	reuseOddObj *simdjson.Object,
	reuseOdd *Odd,
) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
	}

	if reuseOddsObj == nil {
		reuseOddsObj = new(simdjson.Object)
	}

	if reuseOddObj == nil {
		reuseOddObj = new(simdjson.Object)
	}

	if reuseOdd == nil {
		reuseOdd = new(Odd)
	}

	for {
		name, elementType, err := obj.NextElementBytes(reuseIter)
		if err != nil {
			return fmt.Errorf("next element: %w", err)
		}

		if elementType == simdjson.TypeNone {
			break
		}

		switch string(name) {
		case "id":
			m.ID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "status":
			var value int64

			value, err = reuseIter.Int()
			m.Status = Status(value)
		case "type_id":
			m.TypeID, err = simdutil.IntFromIter(reuseIter)
		case "template":
			m.Template, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "flags":
			m.Flags, err = simdutil.IntFromIter(reuseIter)
		case "is_defective":
			m.IsDefective, err = reuseIter.Bool()
		case "specifiers":
			m.Specifiers, err = simdutil.MapStrStrFromIter(reuseIter)
		case "odds":
			m.Odds = make(Odds, 4)

			oddsObj, err := reuseIter.Object(reuseOddsObj)
			if err != nil {
				return fmt.Errorf("create %q object: %w", name, err)
			}

			err = m.Odds.UnmarshalSimdJSON(oddsObj, reuseIter, reuseOddObj, reuseOdd)
		case "meta":
			m.Meta, err = simdutil.MapStrAnyFromIter(reuseIter)
		default:
			continue
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}
