package fixture

import (
	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

type Venue struct {
	ID string `json:"id"`
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
