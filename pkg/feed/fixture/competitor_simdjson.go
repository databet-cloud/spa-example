package fixture

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

func (c Competitors) UnmarshalSimdJSON(
	obj *simdjson.Object,
	reuseIter *simdjson.Iter,
	reuseCompetitorObj *simdjson.Object,
	reuseCompetitor *Competitor,
	reuseScoresObj *simdjson.Object,
	reuseScoreObj *simdjson.Object,
	reuseScore *Score,
) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
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
		name, elementType, err := obj.NextElement(reuseIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		competitorObj, err := reuseIter.Object(reuseCompetitorObj)
		if err != nil {
			return fmt.Errorf("create %q object: %w", name, err)
		}

		err = reuseCompetitor.UnmarshalSimdJSON(competitorObj, reuseIter, reuseScoresObj, reuseScoreObj, reuseScore)
		if err != nil {
			return fmt.Errorf("unmarshal %q odd: %w", name, err)
		}

		c[name] = *reuseCompetitor
	}

	return nil
}

func (c *Competitor) UnmarshalSimdJSON(
	obj *simdjson.Object,
	reuseIter *simdjson.Iter,
	reuseScoresObj *simdjson.Object,
	reuseScoreObj *simdjson.Object,
	reuseScore *Score,
) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
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
		name, elementType, err := obj.NextElement(reuseIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		switch name {
		case "id":
			c.ID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "type":
			var v int

			v, err = simdutil.IntFromIter(reuseIter)
			c.Type = CompetitorType(v)
		case "home_away":
			var v int

			v, err = simdutil.IntFromIter(reuseIter)
			c.HomeAway = CompetitorSide(v)
		case "template_position":
			c.TemplatePosition, err = simdutil.IntFromIter(reuseIter)
		case "scores":
			scoresObj, err := reuseIter.Object(reuseScoresObj)
			if err != nil {
				return err
			}

			c.Scores = make(Scores, 4)
			err = c.Scores.UnmarshalSimdJSON(scoresObj, reuseIter, reuseScoreObj, reuseScore)
		case "name":
			c.Name, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "country_code":
			c.CountryCode, err = simdutil.UnsafeStrFromIter(reuseIter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}
