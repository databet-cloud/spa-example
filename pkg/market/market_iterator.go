package market

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

type Iterator struct {
	tmpIter simdjson.Iter
	tmpObj  simdjson.Object

	embeddedObjIter *simdjson.Iter
	embeddedObjRoot *simdjson.Object
}

func NewIterator(rootIter *simdjson.Iter, marketsPath ...string) (*Iterator, error) {
	var err error
	result := &Iterator{
		embeddedObjRoot: &simdjson.Object{},
	}

	if result.embeddedObjIter, err = simdutil.RewindIterToPath(rootIter, marketsPath...); err != nil {
		return nil, err
	}

	return result, nil
}

func (i *Iterator) Rewind() error {
	var err error
	i.embeddedObjRoot, err = i.embeddedObjIter.Object(i.embeddedObjRoot)

	return err
}

func (i *Iterator) Next(dst *Market) (*Market, error) {
	if dst == nil {
		dst = new(Market)
	}

	_, t, err := i.embeddedObjRoot.NextElementBytes(&i.tmpIter)
	if err != nil {
		return nil, fmt.Errorf("next element bytes: %w", err)
	}

	if t == simdjson.TypeNone {
		// Done
		return nil, nil
	}

	obj, err := i.tmpIter.Object(&i.tmpObj)
	if err != nil {
		return nil, err
	}

	err = dst.UnmarshalSimdJSON(obj)
	if err != nil {
		return nil, fmt.Errorf("unmarshal market simdjson: %w", err)
	}

	return dst, nil
}
