package statistics

const (
	BeachVolleyballSetRegular BeachVolleyballSetType = "regular"
	BeachVolleyballSetGolden  BeachVolleyballSetType = "golden"
)

type BeachVolleyballSetType string

func (s BeachVolleyballStatistic) Typ() string {
	return s.Type
}

type BeachVolleyballStatistic struct {
	Type string               `json:"type"`
	Sets []BeachVolleyballSet `json:"sets"`
}

type BeachVolleyballSet struct {
	Number        int                    `json:"number"`
	Type          BeachVolleyballSetType `json:"type"`
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
