package statistics

func (s LOLStatistic) Typ() string {
	return s.Type
}

type LOLStatistic struct {
	Type string   `json:"type"`
	Maps []LOLMap `json:"maps"`
}

type LOLMap struct {
	Number    int    `json:"number"`
	Winner    string `json:"winner"`
	HomeKills int    `json:"home_kills"`
	AwayKills int    `json:"away_kills"`
	GoldLead  int    `json:"gold_lead"`
	Timer     Timer  `json:"timer"`
}
