package fixture

import (
	"encoding/json"
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

type Tournament struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	MasterID    string `json:"master_id"`
	CountryCode string `json:"country_code"`
}

func (t *Tournament) ApplyPatch(key string, value json.RawMessage) error {
	var unmarshaller any

	switch key {
	case "id":
		unmarshaller = &t.ID
	case "name":
		unmarshaller = &t.Name
	case "master_id":
		unmarshaller = &t.MasterID
	case "country_code":
		unmarshaller = &t.CountryCode
	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
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
		case "master_id":
			t.MasterID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "country_code":
			t.CountryCode, err = simdutil.UnsafeStrFromIter(reuseIter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", string(name), err)
		}
	}

	return nil
}
