package statistics

const (
	Dota2SideDire    = "dire"
	Dota2SideRadiant = "radiant"
)

func (s Dota2Statistic) Typ() string {
	return s.Type
}

type Dota2Statistic struct {
	Type string     `json:"type"`
	Maps []Dota2Map `json:"maps"`
}

type Dota2Map struct {
	Number       int    `json:"number"`
	Winner       string `json:"winner"`
	HomeTeamSide string `json:"home_team_side"`
	AwayTeamSide string `json:"away_team_side"`
	HomeKills    int    `json:"home_kills"`
	AwayKills    int    `json:"away_kills"`
	GoldLead     int    `json:"gold_lead"`
	Timer        Timer  `json:"timer"`
}
