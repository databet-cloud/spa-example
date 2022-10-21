package fixture

import (
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

func (c Competitor) WithPatch(tree patch.Tree) Competitor {
	if v, ok := patch.GetFromTree[string](tree, "id"); ok {
		c.ID = v
	}

	if v, ok := patch.GetFromTree[int](tree, "type"); ok {
		c.Type = v
	} else if v, ok := patch.GetFromTree[float64](tree, "type"); ok {
		c.Type = int(v)
	}

	if v, ok := patch.GetFromTree[int](tree, "home_away"); ok {
		c.HomeAway = v
	} else if v, ok := patch.GetFromTree[float64](tree, "home_away"); ok {
		c.HomeAway = int(v)
	}

	if v, ok := patch.GetFromTree[int](tree, "template_position"); ok {
		c.TemplatePosition = v
	} else if v, ok := patch.GetFromTree[float64](tree, "template_position"); ok {
		c.TemplatePosition = int(v)
	}

	if v, ok := patch.GetFromTree[string](tree, "name"); ok {
		c.Name = v
	}

	if v, ok := patch.GetFromTree[string](tree, "master_id"); ok {
		c.MasterID = v
	}

	if v, ok := patch.GetFromTree[string](tree, "country_code"); ok {
		c.CountryCode = v
	}

	if subTree := tree.SubTree("score"); !subTree.Empty() {
		c.Scores = patch.MapPatchable(c.Scores, subTree)
	}

	return c
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
