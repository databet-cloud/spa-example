package statistics

func (s EVolleyballStatistic) GetType() Type {
	return s.Type
}

type EVolleyballStatistic struct {
	Type Type            `json:"type"`
	Sets []VolleyballSet `json:"sets"`
}

type EVolleyballSet struct {
	Number        int                `json:"number"`
	Type          SetType            `json:"type"`
	SetServer     Team               `json:"set_server"`
	CurrentServer Team               `json:"current_server"`
	Winner        Team               `json:"winner"`
	BallSide      Team               `json:"ball_side"`
	Points        []EVolleyballPoint `json:"points"`
}

type EVolleyballPoint struct {
	Number int  `json:"number"`
	Winner Team `json:"winner"`
	Server Team `json:"server"`
	Value  int  `json:"value"`
}
