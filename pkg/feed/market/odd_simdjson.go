package market

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

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
