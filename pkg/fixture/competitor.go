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

func (c *Competitor) UnmarshalSimdJSON(obj *simdjson.Object) error {
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
			c.ID, err = iter.String()
		case "type":
			c.Type, err = simdutil.IntFromIter(iter)
		case "home_away":
			c.HomeAway, err = simdutil.IntFromIter(iter)
		case "template_position":
			c.TemplatePosition, err = simdutil.IntFromIter(iter)
		case "scores":
			obj, err := iter.Object(nil)
			if err != nil {
				return err
			}

			c.Scores = make(Scores)
			err = c.Scores.UnmarshalSimdJSON(obj)
		case "name":
			c.Name, err = iter.String()
		case "master_id":
			c.MasterID, err = iter.String()
		case "country_code":
			c.CountryCode, err = iter.String()
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}

func (c *Competitor) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	var (
		err                     error
		key, rest, partialPatch = strings.Cut(path, "/")
	)

	switch key {
	case "id":
		c.ID, err = iter.String()
	case "type":
		c.Type, err = simdutil.IntFromIter(iter)
	case "home_away":
		c.HomeAway, err = simdutil.IntFromIter(iter)
	case "template_position":
		c.TemplatePosition, err = simdutil.IntFromIter(iter)
	case "name":
		c.Name, err = iter.String()
	case "master_id":
		c.MasterID, err = iter.String()
	case "country_code":
		c.CountryCode, err = iter.String()
	case "scores":
		if partialPatch {
			if c.Scores == nil {
				return fmt.Errorf("partial patch nil scores")
			}

			return c.Scores.ApplyPatchSimdJSON(rest, iter)
		}

		obj, err := iter.Object(nil)
		if err != nil {
			return err
		}

		c.Scores = make(Scores)
		return c.Scores.UnmarshalSimdJSON(obj)

	case "score":
		if c.Scores == nil {
			return fmt.Errorf("partial patch nil scores")
		}
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
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

func (c Competitors) UnmarshalSimdJSON(obj *simdjson.Object) error {
	tmpIter := new(simdjson.Iter)
	competitorObj := new(simdjson.Object)
	tmpCompetitor := new(Competitor)

	for {
		name, elementType, err := obj.NextElement(tmpIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		competitorObj, err = tmpIter.Object(competitorObj)
		if err != nil {
			return fmt.Errorf("create %q object: %w", name, err)
		}

		err = tmpCompetitor.UnmarshalSimdJSON(competitorObj)
		if err != nil {
			return fmt.Errorf("unmarshal %q odd: %w", name, err)
		}

		c[name] = *tmpCompetitor
	}

	return nil
}

func (c Competitors) FromIter(iter *simdjson.Iter, dst *simdjson.Object) error {
	obj, err := iter.Object(dst)
	if err != nil {
		return err
	}

	return c.UnmarshalSimdJSON(obj)
}

func (c Competitors) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	competitor, ok := c[key]

	if !partialPatch {
		obj, err := iter.Object(nil)
		if err != nil {
			return err
		}

		err = competitor.UnmarshalSimdJSON(obj)
		if err != nil {
			return fmt.Errorf("competitor %q unmarshal simdjson: %w", key, err)
		}

		c[key] = competitor
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent competitor: %q", key)
	}

	err := competitor.ApplyPatchSimdJSON(rest, iter)
	if err != nil {
		return fmt.Errorf("apply competitor patch: %w", err)
	}

	c[key] = competitor
	return nil
}
