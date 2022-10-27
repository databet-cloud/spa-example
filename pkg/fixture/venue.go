//go:generate go run github.com/mailru/easyjson/easyjson venue.go
package fixture

import (
	"encoding/json"

	"github.com/minio/simdjson-go"
)

//easyjson:json
type Venue struct {
	ID string `json:"id"`
}

func (v *Venue) ApplyPatch(path string, value json.RawMessage) error {
	if path == "id" {
		return json.Unmarshal(value, &v.ID)
	}

	return nil
}

func (v *Venue) UnmarshalSimdJSON(obj *simdjson.Object) error {
	element := obj.FindKey("id", nil)
	if element == nil {
		return nil
	}

	id, err := element.Iter.String()
	if err != nil {
		return err
	}

	v.ID = id
	return nil
}

func (v *Venue) FromIter(iter *simdjson.Iter, dst *simdjson.Object) error {
	obj, err := iter.Object(dst)
	if err != nil {
		return err
	}

	return v.UnmarshalSimdJSON(obj)
}

func (v *Venue) ApplyPatchSimdJSON(key string, iter *simdjson.Iter) error {
	if key != "id" {
		return nil
	}

	id, err := iter.String()
	if err != nil {
		return err
	}

	v.ID = id
	return nil
}
