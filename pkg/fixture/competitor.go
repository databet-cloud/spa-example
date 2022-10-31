//go:generate go run github.com/mailru/easyjson/easyjson competitor.go
package fixture

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

const (
	CompetitorTypeOther = iota
	CompetitorTypePerson
	CompetitorTypeTeam
)

const (
	CompetitorUnknown = iota
	CompetitorHome
	CompetitorAway
)

//easyjson:json
type Competitor struct {
	ID               string `json:"id"`
	Type             int    `json:"type"`
	HomeAway         int    `json:"home_away"`
	TemplatePosition int    `json:"template_position"`
	Scores           Scores `json:"scores"`
	//
	Name        string `json:"name"`
	MasterID    string `json:"master_id"`
	CountryCode string `json:"country_code"`
}

func (c *Competitor) ApplyPatch(path string, value json.RawMessage) error {
	var (
		unmarshaller     any
		key, rest, found = strings.Cut(path, "/")
		partialPatch     = found
	)

	switch key {
	case "id":
		unmarshaller = &c.ID
	case "type":
		unmarshaller = &c.Type
	case "home_away":
		unmarshaller = &c.HomeAway
	case "template_position":
		unmarshaller = &c.TemplatePosition
	case "name":
		unmarshaller = &c.Name
	case "master_id":
		unmarshaller = &c.MasterID
	case "country_code":
		unmarshaller = &c.CountryCode
	case "scores":
		if partialPatch {
			if c.Scores == nil {
				return fmt.Errorf("patch nil scores")
			}

			return c.Scores.ApplyPatch(rest, value)
		}

		unmarshaller = &c.Scores
	case "score":
		if c.Scores == nil {
			return fmt.Errorf("patch nil scores")
		}

		return c.Scores.ApplyPatch(rest, value)

	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("unmarshal value: %w, path: %q", err, path)
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
			c.Type, err = simdutil.IntFromIter(reuseIter)
		case "home_away":
			c.HomeAway, err = simdutil.IntFromIter(reuseIter)
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
		case "master_id":
			c.MasterID, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "country_code":
			c.CountryCode, err = simdutil.UnsafeStrFromIter(reuseIter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}

func (c Competitor) Clone() Competitor {
	result := c
	c.Scores = c.Scores.Clone()

	return result
}

//easyjson:json
type Competitors map[string]Competitor

func (c Competitors) Clone() Competitors {
	result := make(Competitors, len(c))
	for k, v := range c {
		result[k] = v.Clone()
	}

	return result
}

func (c Competitors) ApplyPatch(path string, value json.RawMessage) error {
	key, rest, found := strings.Cut(path, "/")
	competitor, ok := c[key]

	if !found {
		err := json.Unmarshal(value, &competitor)
		if err != nil {
			return fmt.Errorf("unmarshal competitor: %w, path: %q", err, path)
		}

		c[key] = competitor
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent competitor: %q", path)
	}

	err := competitor.ApplyPatch(rest, value)
	if err != nil {
		return fmt.Errorf("patch competitor: %w, path: %q", err, rest)
	}

	c[key] = competitor
	return nil
}

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
