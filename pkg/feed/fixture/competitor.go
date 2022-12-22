package fixture

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
	Name     string `json:"name"`
	MasterID string `json:"master_id"`
	// CountryCode ISO 3166-1 alpha-2
	CountryCode string `json:"country_code"`
}

func (c Competitor) Clone() Competitor {
	c.Scores = c.Scores.Clone()

	return c
}

type Competitors map[string]Competitor

func (c Competitors) Clone() Competitors {
	result := make(Competitors, len(c))
	for k, v := range c {
		result[k] = v.Clone()
	}

	return result
}
