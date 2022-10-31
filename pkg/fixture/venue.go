//go:generate go run github.com/mailru/easyjson/easyjson venue.go
package fixture

import (
	"encoding/json"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
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

	id, err := simdutil.UnsafeStrFromIter(&element.Iter)
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
