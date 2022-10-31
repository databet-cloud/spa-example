package simdutil

import (
	"errors"
	"fmt"
	"time"
	"unsafe"

	"github.com/minio/simdjson-go"
)

func RewindIterToPath(iter *simdjson.Iter, keys ...string) (*simdjson.Iter, error) {
	var err error

	obj, elem := &simdjson.Object{}, simdjson.Element{}

	for _, key := range keys {
		if obj, err = iter.Object(obj); err != nil {
			return nil, err
		}

		if e := obj.FindKey(key, &elem); e == nil {
			return nil, errors.New("not found")
		}

		iter = &elem.Iter
	}

	return iter, nil
}

func IntFromIter(iter *simdjson.Iter) (int, error) {
	i64, err := iter.Int()

	return int(i64), err
}

func UnsafeStrFromIter(iter *simdjson.Iter) (string, error) {
	b, err := iter.StringBytes()
	if err != nil {
		return "", err
	}

	return *(*string)(unsafe.Pointer(&b)), nil
}

func UnsafeStrFromBytes(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func MapStrStrFromIter(iter *simdjson.Iter) (map[string]string, error) {
	o, err := iter.Object(nil)
	if err != nil {
		return nil, err
	}

	dst := map[string]string{}

	for {
		name, t, err := o.NextElementBytes(iter)
		if err != nil {
			return nil, err
		}

		if t == simdjson.TypeNone {
			// Done
			break
		}

		dst[*(*string)(unsafe.Pointer(&name))], err = UnsafeStrFromIter(iter)
		if err != nil {
			return nil, fmt.Errorf("parsing element %q: %w", name, err)
		}
	}

	if dst == nil {
		return map[string]string{}, nil
	}

	return dst, nil
}

func MapStrAnyFromIter(iter *simdjson.Iter) (map[string]any, error) {
	o, err := iter.Object(nil)
	if err != nil {
		return nil, err
	}

	dst := make(map[string]any)

	for {
		name, t, err := o.NextElement(iter)
		if err != nil {
			return nil, err
		}

		if t == simdjson.TypeNone {
			// Done
			break
		}

		dst[name], err = iter.Interface()
		if err != nil {
			return nil, fmt.Errorf("parsing element %q: %w", name, err)
		}
	}

	if dst == nil {
		return map[string]any{}, nil
	}

	return dst, nil
}

func TimeFromIter(iter *simdjson.Iter) (time.Time, error) {
	rawTime, err := UnsafeStrFromIter(iter)
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.Parse(time.RFC3339, rawTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse time: %w", err)
	}

	return t, nil
}

func CreateRootIter(parsedJson *simdjson.ParsedJson) (*simdjson.Iter, error) {
	var err error

	parsedIter := parsedJson.Iter()
	if simdjson.TypeRoot != parsedIter.Advance() {
		return nil, errors.New("invalid obj")
	}

	rootIter := &simdjson.Iter{}
	if _, rootIter, err = parsedIter.Root(rootIter); err != nil {
		return nil, err
	}

	return rootIter, nil
}
