package fixture

import (
	"encoding/json"
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

func FromIter(rootIter *simdjson.Iter, path ...string) (*Fixture, error) {
	var err error

	rootIter, err = simdutil.RewindIterToPath(rootIter, path...)
	if err != nil {
		return nil, err
	}

	sportEvent := new(Fixture)

	obj, err := rootIter.Object(nil)
	if err != nil {
		return nil, err
	}

	err = sportEvent.UnmarshalSimdJSON(obj)
	if err != nil {
		return nil, fmt.Errorf("unmarshal simdjson: %w", err)
	}

	return sportEvent, nil
}

func FromJson(eventData json.RawMessage) (*Fixture, error) {
	rootIter, err := simdutil.JSONToRootIter(eventData)
	if err != nil {
		return nil, err
	}

	return FromIter(rootIter, "sport_event")
}
