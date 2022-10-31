//go:generate go run github.com/mailru/easyjson/easyjson stream.go
package fixture

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

//easyjson:json
type Stream struct {
	ID        string    `json:"id"`
	Locale    string    `json:"locale"`
	URL       string    `json:"url"`
	Platforms Platforms `json:"platforms"`
	Priority  int       `json:"priority"`
}

func (s *Stream) ApplyPatch(path string, value json.RawMessage) error {
	var (
		unmarshaller     any
		key, rest, found = strings.Cut(path, "/")
		partialPatch     = found
	)

	switch key {
	case "id":
		unmarshaller = &s.ID
	case "locale":
		unmarshaller = &s.Locale
	case "url":
		unmarshaller = &s.URL
	case "priority":
		unmarshaller = &s.Priority
	case "platforms":
		if partialPatch {
			if s.Platforms == nil {
				return fmt.Errorf("patch nil platfroms")
			}

			return s.Platforms.ApplyPatch(rest, value)
		}

		unmarshaller = &s.Platforms
	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("unmarshal value: %w, path: %q", err, path)
	}

	return nil
}

func (s *Stream) UnmarshalSimdJSON(obj *simdjson.Object) error {
	iter := new(simdjson.Iter)

	for {
		name, elementType, err := obj.NextElementBytes(iter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		switch string(name) {
		case "id":
			s.ID, err = simdutil.UnsafeStrFromIter(iter)
		case "locale":
			s.Locale, err = simdutil.UnsafeStrFromIter(iter)
		case "url":
			s.URL, err = simdutil.UnsafeStrFromIter(iter)
		case "platforms":
			obj, err := iter.Object(nil)
			if err != nil {
				return err
			}

			s.Platforms = make(Platforms)

			err = s.Platforms.UnmarshalSimdJSON(obj)

		case "priority":
			s.Priority, err = simdutil.IntFromIter(iter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}

func (s *Stream) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	var (
		err                     error
		key, rest, partialPatch = strings.Cut(path, "/")
	)

	switch path {
	case "id":
		s.ID, err = simdutil.UnsafeStrFromIter(iter)
	case "locale":
		s.Locale, err = simdutil.UnsafeStrFromIter(iter)
	case "url":
		s.URL, err = simdutil.UnsafeStrFromIter(iter)
	case "platforms":
		if partialPatch {
			if s.Platforms == nil {
				return fmt.Errorf("partial patch nil platforms")
			}

			s.Platforms.ApplyPatchSimdJSON(rest, iter)
		}

		s.Platforms.FromIter(iter)
	case "priority":
		s.Priority, err = simdutil.IntFromIter(iter)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}

type Streams map[string]Stream

func (s Streams) ApplyPatch(path string, value json.RawMessage) error {
	key, rest, found := strings.Cut(path, "/")
	stream, ok := s[key]

	if !found {
		err := json.Unmarshal(value, &stream)
		if err != nil {
			return fmt.Errorf("unmarshal stream: %w, path: %q", err, path)
		}

		s[key] = stream
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent stream: %q", path)
	}

	err := stream.ApplyPatch(rest, value)
	if err != nil {
		return fmt.Errorf("patch stream: %w, path: %q", err, rest)
	}

	s[key] = stream
	return nil
}

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

		err = tmpStream.UnmarshalSimdJSON(streamObj)
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

func (s Streams) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	stream, ok := s[key]

	if !partialPatch {
		obj, err := iter.Object(nil)
		if err != nil {
			return err
		}

		err = stream.UnmarshalSimdJSON(obj)
		if err != nil {
			return fmt.Errorf("stream %q unmarshal simdjson: %w", key, err)
		}

		s[key] = stream
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent stream: %q", key)
	}

	err := stream.ApplyPatchSimdJSON(rest, iter)
	if err != nil {
		return fmt.Errorf("apply stream patch: %w", err)
	}

	s[key] = stream
	return nil
}
