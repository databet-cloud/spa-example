package fixture

import (
	"fmt"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
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

type CompetitorPatch struct {
	ID               *string `mapstructure:"id"`
	Type             *int    `mapstructure:"type"`
	HomeAway         *int    `mapstructure:"home_away"`
	TemplatePosition *int    `mapstructure:"template_position"`
	Name             *string `mapstructure:"name"`
	MasterID         *string `mapstructure:"master_id"`
	CountryCode      *string `mapstructure:"country_code"`
}

func (c Competitor) WithPatch(tree patch.Tree) (Competitor, error) {
	var competitorPatch CompetitorPatch

	err := tree.UnmarshalPatch(&competitorPatch)
	if err != nil {
		return Competitor{}, fmt.Errorf("unmarshal competitor patch: %w", err)
	}

	c.applyCompetitorPatch(competitorPatch)

	if subTree := tree.SubTree("score"); !subTree.Empty() {
		c.Scores, err = patch.MapPatchable(c.Scores, subTree)
		if err != nil {
			return Competitor{}, fmt.Errorf("patch scores: %w", err)
		}
	}

	return c, nil
}

func (c *Competitor) ApplyPatch(tree patch.Tree) error {
	var competitorPatch CompetitorPatch

	err := tree.UnmarshalPatch(&competitorPatch)
	if err != nil {
		return fmt.Errorf("unmarshal competitor patch: %w", err)
	}

	c.applyCompetitorPatch(competitorPatch)

	if subTree := tree.SubTree("score"); !subTree.Empty() {
		if c.Scores == nil {
			c.Scores = map[string]Score{}
		}

		for id, subTree := range subTree.SubTrees() {
			v := c.Scores[id]

			err := v.ApplyPatch(subTree)
			if err != nil {
				return err
			}

			c.Scores[id] = v
		}
	}

	return nil
}

func (c *Competitor) applyCompetitorPatch(patch CompetitorPatch) {
	if patch.ID != nil {
		c.ID = *patch.ID
	}

	if patch.Type != nil {
		c.Type = *patch.Type
	}

	if patch.HomeAway != nil {
		c.HomeAway = *patch.HomeAway
	}

	if patch.TemplatePosition != nil {
		c.TemplatePosition = *patch.TemplatePosition
	}

	if patch.Name != nil {
		c.Name = *patch.Name
	}

	if patch.MasterID != nil {
		c.MasterID = *patch.MasterID
	}

	if patch.CountryCode != nil {
		c.CountryCode = *patch.CountryCode
	}
}

func (c Competitor) Clone() Competitor {
	result := c
	c.Scores = c.Scores.Clone()

	return result
}

type Competitors map[string]Competitor

func (c Competitors) Clone() Competitors {
	result := make(Competitors, len(c))
	for k, v := range c {
		result[k] = v.Clone()
	}

	return result
}
