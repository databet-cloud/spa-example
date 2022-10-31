package fixture

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

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

func (s *Score) UnmarshalSimdJSON(obj *simdjson.Object, reuseIter *simdjson.Iter) error {
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
		case "type":
			s.Type, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "points":
			s.Points, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "number":
			s.Number, err = simdutil.IntFromIter(reuseIter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}

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

func (s Scores) UnmarshalSimdJSON(obj *simdjson.Object, reuseIter *simdjson.Iter, reuseScoreObj *simdjson.Object, reuseScore *Score) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
	}

	if reuseScoreObj == nil {
		reuseScoreObj = new(simdjson.Object)
	}

	if reuseScore == nil {
		reuseScore = new(Score)
	}

	for {
		name, elementType, err := obj.NextElement(reuseIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		scoreObj, err := reuseIter.Object(reuseScoreObj)
		if err != nil {
			return fmt.Errorf("create %q object: %w", name, err)
		}

		err = reuseScore.UnmarshalSimdJSON(scoreObj, nil)
		if err != nil {
			return fmt.Errorf("unmarshal %q odd: %w", name, err)
		}

		s[name] = *reuseScore
	}

	return nil
}
