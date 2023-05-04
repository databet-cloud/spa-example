package statistics

func (s VolleyballStatistic) GetType() Type {
	return s.Type
}

type VolleyballStatistic struct {
	Type Type            `json:"type"`
	Sets []VolleyballSet `json:"sets"`
}

type VolleyballSet struct {
	Number        int               `json:"number"`
	Type          SetType           `json:"type"`
	SetServer     Team              `json:"set_server"`
	CurrentServer Team              `json:"current_server"`
	Winner        Team              `json:"winner"`
	BallSide      Team              `json:"ball_side"`
	Points        []VolleyballPoint `json:"points"`
}

type VolleyballPoint struct {
	Number int  `json:"number"`
	Winner Team `json:"winner"`
	Server Team `json:"server"`
	Value  int  `json:"value"`
}
