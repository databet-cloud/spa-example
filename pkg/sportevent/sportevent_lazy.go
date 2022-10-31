package sportevent

import (
	"fmt"
	"time"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/market"
	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

type SportEventLazy struct {
	MarketIter *market.Iterator

	ID        string                 `json:"id"`
	Meta      map[string]interface{} `json:"meta"`
	Fixture   fixture.Fixture        `json:"fixture"`
	Markets   market.Markets         `json:"markets"`
	BetStop   bool                   `json:"bet_stop"`
	UpdatedAt time.Time              `json:"updated_at"`
}

func (se *SportEventLazy) UnmarshalJSON(bytes []byte) error {
	rootIter, err := simdutil.JSONToRootIter(bytes)
	if err != nil {
		return err
	}

	obj, err := rootIter.Object(nil)
	if err != nil {
		return err
	}

	iter := new(simdjson.Iter)
	reuseObj := new(simdjson.Object)
	reuseCompetitorObj := new(simdjson.Object)
	reuseCompetitor := new(fixture.Competitor)
	reuseScoreObj := new(simdjson.Object)
	reuseScore := new(fixture.Score)

	for {
		name, t, err := obj.NextElementBytes(iter)
		if err != nil {
			return err
		}

		if t == simdjson.TypeNone {
			// Done
			break
		}

		switch string(name) {
		case "id":
			se.ID, err = simdutil.UnsafeStrFromIter(iter)
		case "markets":
			tmpIter := *iter // cloning iter
			se.MarketIter, err = market.NewIterator(&tmpIter)
		case "fixture":
			obj, err := iter.Object(nil)
			if err != nil {
				return err
			}

			err = se.Fixture.UnmarshalSimdJSON(obj, iter, reuseObj, reuseCompetitorObj, reuseCompetitor, nil, reuseScoreObj, reuseScore)
		case "bet_stop":
			se.BetStop, err = iter.Bool()
		case "updated_at":
			se.UpdatedAt, err = simdutil.TimeFromIter(iter)
		default:
			continue
		}

		if err != nil {
			return fmt.Errorf("parsing element %q: %w", name, err)
		}
	}

	return nil
}
