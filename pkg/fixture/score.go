//go:generate go run github.com/mailru/easyjson/easyjson score.go
package fixture

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

//easyjson:json
type Score struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Points string `json:"points"`
	Number int    `json:"number"`
}

func (s *Score) ApplyPatch(key string, value json.RawMessage) error {
	var unmarshaller any

	switch key {
	case "id":
		unmarshaller = &s.ID
	case "type":
		unmarshaller = &s.Type
	case "points":
		unmarshaller = &s.Points
	case "number":
		unmarshaller = &s.Number
	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("unmarshal score value: %w, field: %q", err, key)
	}

	return nil
}

func (s *Score) UnmarshalSimdJSON(obj *simdjson.Object) error {
	iter := new(simdjson.Iter)

	for {
		name, elementType, err := obj.NextElement(iter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		switch name {
		case "id":
			s.ID, err = iter.String()
		case "type":
			s.Type, err = iter.String()
		case "points":
			s.Points, err = iter.String()
		case "number":
			s.Number, err = simdutil.IntFromIter(iter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}

func (s *Score) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	var err error

	switch path {
	case "id":
		s.ID, err = iter.String()
	case "type":
		s.Type, err = iter.String()
	case "points":
		s.Points, err = iter.String()
	case "number":
		s.Number, err = simdutil.IntFromIter(iter)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", path, err)
	}

	return nil
}

//easyjson:json
type Scores map[string]Score

func (s Scores) Clone() Scores {
	result := make(Scores, len(s))
	for k, v := range s {
		result[k] = v
	}

	return result
}

func (s Scores) ApplyPatch(path string, value json.RawMessage) error {
	key, rest, found := strings.Cut(path, "/")
	score, ok := s[key]

	if !found {
		err := json.Unmarshal(value, &score)
		if err != nil {
			return fmt.Errorf("unmarshal competitor: %w, key: %q", err, key)
		}

		s[key] = score
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent score: %q", key)
	}

	err := score.ApplyPatch(rest, value)
	if err != nil {
		return fmt.Errorf("patch competitor: %w, path: %q", err, rest)
	}

	s[key] = score
	return nil
}

func (s Scores) UnmarshalSimdJSON(obj *simdjson.Object) error {
	tmpIter := new(simdjson.Iter)
	scoreObj := new(simdjson.Object)
	tmpScore := new(Score)

	for {
		name, elementType, err := obj.NextElement(tmpIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		scoreObj, err = tmpIter.Object(scoreObj)
		if err != nil {
			return fmt.Errorf("create %q object: %w", name, err)
		}

		err = tmpScore.UnmarshalSimdJSON(scoreObj)
		if err != nil {
			return fmt.Errorf("unmarshal %q odd: %w", name, err)
		}

		s[name] = *tmpScore
	}

	return nil
}

func (c Scores) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	score, ok := c[key]

	if !partialPatch {
		obj, err := iter.Object(nil)
		if err != nil {
			return err
		}

		err = score.UnmarshalSimdJSON(obj)
		if err != nil {
			return fmt.Errorf("score %q unmarshal simdjson: %w", key, err)
		}

		c[key] = score
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent score: %q", key)
	}

	err := score.ApplyPatchSimdJSON(rest, iter)
	if err != nil {
		return fmt.Errorf("apply score patch: %w", err)
	}

	c[key] = score
	return nil
}
