package fixture

type CompetitorType int

const (
	CompetitorTypeOther CompetitorType = iota
	CompetitorTypePerson
	CompetitorTypeTeam
)

type CompetitorSide int

const (
	CompetitorSideUnknown CompetitorSide = iota
	CompetitorSideHome
	CompetitorSideAway
)

type Competitor struct {
	// ID of the competitor
	ID string `json:"id"`
	// Type of the competitor
	Type CompetitorType `json:"type"`
	// HomeAway indicates side of the competitor in the current sport event
	HomeAway CompetitorSide `json:"home_away"`
	// TemplatePosition indicates index of the competitor, to replace variables in fixture name with it
	TemplatePosition int    `json:"template_position"`
	Scores           Scores `json:"scores"`
	// Name of the competitor in default locale
	Name string `json:"name"`
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
