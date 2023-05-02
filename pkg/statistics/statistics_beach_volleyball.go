package statistics

func (s BeachVolleyballStatistic) GetType() Type {
	return s.Type
}

type BeachVolleyballStatistic struct {
	Type Type                 `json:"type"`
	Sets []BeachVolleyballSet `json:"sets"`
}

type BeachVolleyballSet struct {
	Number        int                    `json:"number"`
	Type          SetType                `json:"type"`
	SetServer     Team                   `json:"set_server"`
	CurrentServer Team                   `json:"current_server"`
	Winner        Team                   `json:"winner"`
	BallSide      Team                   `json:"ball_side"`
	Points        []BeachVolleyballPoint `json:"points"`
}

type BeachVolleyballPoint struct {
	Number int  `json:"number"`
	Winner Team `json:"winner"`
	Server Team `json:"server"`
	Value  int  `json:"value"`
}
