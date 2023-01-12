package market

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

// Iterator allows you to lazily iterate through the simdjson.Iter and read markets from it
type Iterator struct {
	tmpIter simdjson.Iter
	tmpObj  simdjson.Object

	embeddedObjIter *simdjson.Iter
	embeddedObjRoot *simdjson.Object

	reuseOddsObj *simdjson.Object
	reuseOddObj  *simdjson.Object
	reuseOdd     *Odd
}

func NewIterator(rootIter *simdjson.Iter, marketsPath ...string) (*Iterator, error) {
	var err error
	result := &Iterator{
		embeddedObjRoot: new(simdjson.Object),
		reuseOddsObj:    new(simdjson.Object),
		reuseOddObj:     new(simdjson.Object),
		reuseOdd:        new(Odd),
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

	err = dst.UnmarshalSimdJSON(obj, &i.tmpIter, i.reuseOddsObj, i.reuseOddObj, i.reuseOdd)
	if err != nil {
		return nil, fmt.Errorf("unmarshal market simdjson: %w", err)
	}

	return dst, nil
}
