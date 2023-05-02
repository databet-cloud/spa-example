package statistics

func (s TableTennisStatistic) GetType() Type {
	return s.Type
}

type TableTennisStatistic struct {
	Type Type             `json:"type"`
	Sets []TableTennisSet `json:"sets"`
}

type TableTennisSet struct {
	Number          int                `json:"number"`
	Type            SetType            `json:"type"`
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
