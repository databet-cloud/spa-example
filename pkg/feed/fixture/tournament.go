package fixture

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

type Tournament struct {
	// ID of the tournament
	ID string `json:"id"`

	// Name of the tournament
	Name string `json:"name"`

	// CountryCode ISO 3166-1 alpha-2
	CountryCode string `json:"country_code"`
}

func (t *Tournament) UnmarshalSimdJSON(obj *simdjson.Object, reuseIter *simdjson.Iter) error {
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
			t.ID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "name":
			t.Name, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "country_code":
			t.CountryCode, err = simdutil.UnsafeStrFromIter(reuseIter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", string(name), err)
		}
	}

	return nil
}
