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

	if competitorPatch.ID != nil {
		c.ID = *competitorPatch.ID
	}

	if competitorPatch.Type != nil {
		c.Type = *competitorPatch.Type
	}

	if competitorPatch.HomeAway != nil {
		c.HomeAway = *competitorPatch.HomeAway
	}

	if competitorPatch.TemplatePosition != nil {
		c.TemplatePosition = *competitorPatch.TemplatePosition
	}

	if competitorPatch.Name != nil {
		c.Name = *competitorPatch.Name
	}

	if competitorPatch.MasterID != nil {
		c.MasterID = *competitorPatch.MasterID
	}

	if competitorPatch.CountryCode != nil {
		c.CountryCode = *competitorPatch.CountryCode
	}

	if subTree := tree.SubTree("score"); !subTree.Empty() {
		c.Scores, err = patch.MapPatchable(c.Scores, subTree)
		if err != nil {
			return Competitor{}, fmt.Errorf("patch scores: %w", err)
		}
	}

	return c, nil
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
