package fixture

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

type Score struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Points string `json:"points"`
	Number int    `json:"number"`
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
