package statistics

const (
	BasketballQuarterFormatUnknown       = "unknown"
	BasketballQuarterFormatTenMinutes    = "10_min"
	BasketballQuarterFormatTwelveMinutes = "12_min"
	BasketballQuarterFormatTwentyMinutes = "20_min"
)

func (s BasketballStatistic) Typ() string {
	return s.Type
}

type BasketballStatistic struct {
	Type       string              `json:"type"`
	TimeFormat string              `json:"time_format"`
	Quarters   []BasketballQuarter `json:"quarters"`
}

type BasketballQuarter struct {
	Number             int                   `json:"number"`
	Ended              bool                  `json:"ended"`
	BallPossessionTeam string                `json:"ball_possession_team"`
	Timer              Timer                 `json:"timer"`
	Score              BasketballScore       `json:"score"`
	FoulThrows         []BasketballFoulThrow `json:"foul_throws"`
}

type BasketballFoulThrow struct {
	Number   int  `json:"number"`
	Team     Team `json:"team"`
	Canceled bool `json:"canceled"`
}

type BasketballScore struct {
	Home int `json:"home"`
	Away int `json:"away"`
}
