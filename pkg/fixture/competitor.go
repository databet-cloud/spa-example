//go:generate go run github.com/mailru/easyjson/easyjson competitor.go
package fixture

import (
	"encoding/json"
	"fmt"
	"strings"
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
		return fmt.Errorf("partial patch non-existent competitor: %q", key)
	}

	err := competitor.ApplyPatch(rest, value)
	if err != nil {
		return fmt.Errorf("patch competitor: %w, path: %q", err, rest)
	}

	c[key] = competitor
	return nil
}
