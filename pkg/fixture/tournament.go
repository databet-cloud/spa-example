package fixture

import (
	"encoding/json"
	"fmt"
	"unsafe"

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

func (t *Tournament) UnmarshalSimdJSON(obj *simdjson.Object) error {
	iter := new(simdjson.Iter)

	for {
		name, elementType, err := obj.NextElementBytes(iter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		err = t.unmarshalFieldSimdJSON(*(*string)(unsafe.Pointer(&name)), iter)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Tournament) FromIter(iter *simdjson.Iter, dst *simdjson.Object) error {
	obj, err := iter.Object(dst)
	if err != nil {
		return err
	}

	return t.UnmarshalSimdJSON(obj)
}

func (t *Tournament) ApplyPatchSimdJSON(key string, iter *simdjson.Iter) error {
	return t.unmarshalFieldSimdJSON(key, iter)
}

func (t *Tournament) unmarshalFieldSimdJSON(key string, iter *simdjson.Iter) error {
	var err error

	switch key {
	case "id":
		t.ID, err = simdutil.UnsafeStrFromIter(iter)
	case "name":
		t.Name, err = simdutil.UnsafeStrFromIter(iter)
	case "master_id":
		t.MasterID, err = simdutil.UnsafeStrFromIter(iter)
	case "country_code":
		t.CountryCode, err = simdutil.UnsafeStrFromIter(iter)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}
