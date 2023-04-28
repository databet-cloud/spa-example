package statistics

const (
	TableTennisSetRegular = "regular"
	TableTennisSetGolden  = "golden"
)

func (s TableTennisStatistic) Typ() string {
	return s.Type
}

type TableTennisStatistic struct {
	Type string           `json:"type"`
	Sets []TableTennisSet `json:"sets"`
}

type TableTennisSet struct {
	Number          int                `json:"number"`
	Type            string             `json:"type"`
	SetServer       Team               `json:"set_server"`
	CurrentServer   Team               `json:"current_server"`
	Winner          Team               `json:"winner"`
	Points          []TableTennisPoint `json:"points"`
	TotalPointsHome int                `json:"total_points_home"`
	TotalPointsAway int                `json:"total_points_away"`
}

type TableTennisPoint struct {
	Number int  `json:"number"`
	Winner Team `json:"winner"`
	Server Team `json:"server"`
	Value  int  `json:"value"`
}
