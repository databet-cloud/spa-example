package fixture

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

func (f *Fixture) UnmarshalSimdJSON(
	obj *simdjson.Object,
	reuseIter *simdjson.Iter,
	reuseObj *simdjson.Object,
	reuseCompetitorObj *simdjson.Object,
	reuseCompetitor *Competitor,
	reuseScoresObj *simdjson.Object,
	reuseScoreObj *simdjson.Object,
	reuseScore *Score,
) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
	}

	if reuseObj == nil {
		reuseObj = new(simdjson.Object)
	}

	if reuseCompetitorObj == nil {
		reuseCompetitorObj = new(simdjson.Object)
	}

	if reuseCompetitor == nil {
		reuseCompetitor = new(Competitor)
	}

	if reuseScoresObj == nil {
		reuseScoresObj = new(simdjson.Object)
	}

	if reuseScoreObj == nil {
		reuseScoreObj = new(simdjson.Object)
	}

	if reuseScore == nil {
		reuseScore = new(Score)
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
			f.ID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "version":
			f.Version, err = simdutil.IntFromIter(reuseIter)
		case "template":
			f.Template, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "status":
			var v int
			v, err = simdutil.IntFromIter(reuseIter)

			f.Status = Status(v)
		case "type":
			var v int
			v, err = simdutil.IntFromIter(reuseIter)

			f.Type = Type(v)
		case "sport_id":
			f.SportID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "tournament":
			tournamentObj, err := reuseIter.Object(reuseObj)
			if err != nil {
				return fmt.Errorf("create tournament obj: %w", err)
			}

			err = f.Tournament.UnmarshalSimdJSON(tournamentObj, reuseIter)
		case "venue":
			err = f.Venue.FromIter(reuseIter, reuseObj)
		case "competitors":
			f.Competitors = make(Competitors, 4)

			obj, err := reuseIter.Object(reuseObj)
			if err != nil {
				return fmt.Errorf("create competitors obj: %w", err)
			}

			err = f.Competitors.UnmarshalSimdJSON(obj, reuseIter, reuseCompetitorObj, reuseCompetitor, reuseScoresObj, reuseScoreObj, reuseScore)
		case "streams":
			f.Streams = make(Streams)
			err = f.Streams.FromIter(reuseIter, reuseObj)
		case "live_coverage":
			f.LiveCoverage, err = reuseIter.Bool()
		case "start_time":
			f.StartTime, err = simdutil.TimeFromIter(reuseIter)
		case "flags":
			f.Flags, err = simdutil.IntFromIter(reuseIter)
		case "created_at":
			f.CreatedAt, err = simdutil.TimeFromIter(reuseIter)
		case "updated_at":
			f.UpdatedAt, err = simdutil.TimeFromIter(reuseIter)
		case "published_at":
			f.PublishedAt, err = simdutil.TimeFromIter(reuseIter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}
