package sportevent

import (
	"fmt"
	"time"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/market"
	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

// SportEventLazy could be used to lazy unmarshal and handle markets
type SportEventLazy struct {
	MarketIter *market.Iterator `json:"-"`

	ID        string          `json:"id"`
	Meta      map[string]any  `json:"meta"`
	Fixture   fixture.Fixture `json:"fixture"`
	BetStop   bool            `json:"bet_stop"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (se *SportEventLazy) UnmarshalJSON(data []byte) error {
	parsedJson, err := simdjson.Parse(data, nil, simdjson.WithCopyStrings(false))
	if err != nil {
		return fmt.Errorf("simdjson parse: %w", err)
	}

	rootIter, err := simdutil.CreateRootIter(parsedJson)
	if err != nil {
		return fmt.Errorf("simdjson create root iter: %w", err)
	}

	rootObj, err := rootIter.Object(nil)
	if err != nil {
		return fmt.Errorf("create sport event obj: %w", err)
	}

	// Reuse fields are nil, because we can't reuse them while implementing json.Unmarshaller interface
	return se.UnmarshalSimdJSON(rootObj, nil, nil, nil, nil, nil, nil, nil, nil)
}

func (se *SportEventLazy) UnmarshalSimdJSON(
	obj *simdjson.Object,
	reuseIter *simdjson.Iter,
	reuseObj *simdjson.Object,
	fixtureObj *simdjson.Object,
	reuseCompetitorObj *simdjson.Object,
	reuseCompetitor *fixture.Competitor,
	reuseScoresObj *simdjson.Object,
	reuseScoreObj *simdjson.Object,
	reuseScore *fixture.Score,
) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
	}

	if reuseObj == nil {
		reuseObj = new(simdjson.Object)
	}

	if fixtureObj == nil {
		fixtureObj = new(simdjson.Object)
	}

	if reuseCompetitorObj == nil {
		reuseCompetitorObj = new(simdjson.Object)
	}

	if reuseCompetitor == nil {
		reuseCompetitor = new(fixture.Competitor)
	}

	if reuseScoresObj == nil {
		reuseScoresObj = new(simdjson.Object)
	}

	if reuseScoreObj == nil {
		reuseScoreObj = new(simdjson.Object)
	}

	if reuseScore == nil {
		reuseScore = new(fixture.Score)
	}

	for {
		name, t, err := obj.NextElementBytes(reuseIter)
		if err != nil {
			return err
		}

		if t == simdjson.TypeNone {
			// Done
			break
		}

		switch string(name) {
		case "id":
			se.ID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "markets":
			tmpIter := *reuseIter // cloning iter
			se.MarketIter, err = market.NewIterator(&tmpIter)
		case "meta":
			se.Meta, err = simdutil.MapStrAnyFromIter(reuseIter)
		case "fixture":
			obj, err := reuseIter.Object(fixtureObj)
			if err != nil {
				return err
			}

			err = se.Fixture.UnmarshalSimdJSON(
				obj,
				reuseIter,
				reuseObj,
				reuseCompetitorObj,
				reuseCompetitor,
				reuseScoresObj,
				reuseScoreObj,
				reuseScore,
			)
		case "bet_stop":
			se.BetStop, err = reuseIter.Bool()
		case "updated_at":
			se.UpdatedAt, err = simdutil.TimeFromIter(reuseIter)
		default:
			continue
		}

		if err != nil {
			return fmt.Errorf("parsing element %q: %w", name, err)
		}
	}

	return nil
}
