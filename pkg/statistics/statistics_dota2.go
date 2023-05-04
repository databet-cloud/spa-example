package statistics

type Dota2Side string

const (
	Dota2SideDire    Dota2Side = "dire"
	Dota2SideRadiant Dota2Side = "radiant"
)

func (s Dota2Statistic) GetType() Type {
	return s.Type
}

type Dota2Statistic struct {
	Type Type       `json:"type"`
	Maps []Dota2Map `json:"maps"`
}

type Dota2Map struct {
	Number       int       `json:"number"`
	Winner       Team      `json:"winner"`
	HomeTeamSide Dota2Side `json:"home_team_side"`
	AwayTeamSide Dota2Side `json:"away_team_side"`
	HomeKills    int       `json:"home_kills"`
	AwayKills    int       `json:"away_kills"`
	GoldLead     int       `json:"gold_lead"`
	Timer        Timer     `json:"timer"`
}
