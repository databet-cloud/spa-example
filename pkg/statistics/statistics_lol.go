package statistics

func (s LOLStatistic) GetType() Type {
	return s.Type
}

type LOLPlayerStatistic struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

type LOLStatistic struct {
	Type             Type                 `json:"type"`
	Maps             []LOLMap             `json:"maps"`
	PlayersStatistic []LOLPlayerStatistic `json:"players"`
}

type LOLMap struct {
	Number    int    `json:"number"`
	Winner    string `json:"winner"`
	HomeKills int    `json:"home_kills"`
	AwayKills int    `json:"away_kills"`
	GoldLead  int    `json:"gold_lead"`
	Timer     Timer  `json:"timer"`
}
