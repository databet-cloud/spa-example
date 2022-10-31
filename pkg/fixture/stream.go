package fixture

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

type Stream struct {
	ID        string    `json:"id"`
	Locale    string    `json:"locale"`
	URL       string    `json:"url"`
	Platforms Platforms `json:"platforms"`
	Priority  int       `json:"priority"`
}

func (s *Stream) UnmarshalSimdJSON(obj *simdjson.Object, reuseIter *simdjson.Iter) error {
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
			s.ID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "locale":
			s.Locale, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "url":
			s.URL, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "platforms":
			obj, err := reuseIter.Object(nil)
			if err != nil {
				return err
			}

			s.Platforms = make(Platforms)

			err = s.Platforms.UnmarshalSimdJSON(obj, reuseIter)

		case "priority":
			s.Priority, err = simdutil.IntFromIter(reuseIter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}

type Streams map[string]Stream

func (s Streams) UnmarshalSimdJSON(obj *simdjson.Object) error {
	tmpIter := new(simdjson.Iter)
	streamObj := new(simdjson.Object)
	tmpStream := new(Stream)

	for {
		name, elementType, err := obj.NextElement(tmpIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		streamObj, err = tmpIter.Object(streamObj)
		if err != nil {
			return fmt.Errorf("create %q object: %w", name, err)
		}

		err = tmpStream.UnmarshalSimdJSON(streamObj, tmpIter)
		if err != nil {
			return fmt.Errorf("unmarshal %q odd: %w", name, err)
		}

		s[name] = *tmpStream
	}

	return nil
}

func (s Streams) FromIter(iter *simdjson.Iter, dst *simdjson.Object) error {
	obj, err := iter.Object(dst)
	if err != nil {
		return err
	}

	return s.UnmarshalSimdJSON(obj)
}
