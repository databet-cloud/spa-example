package statistics

type BasketballQuarterFormat string

const (
	BasketballQuarterFormatUnknown       BasketballQuarterFormat = "unknown"
	BasketballQuarterFormatTenMinutes    BasketballQuarterFormat = "10_min"
	BasketballQuarterFormatTwelveMinutes BasketballQuarterFormat = "12_min"
	BasketballQuarterFormatTwentyMinutes BasketballQuarterFormat = "20_min"
)

func (s BasketballStatistic) GetType() Type {
	return s.Type
}

type BasketballStatistic struct {
	Type       Type                    `json:"type"`
	TimeFormat BasketballQuarterFormat `json:"time_format"`
	Quarters   []BasketballQuarter     `json:"quarters"`
}

type BasketballQuarter struct {
	Number             int                   `json:"number"`
	Ended              bool                  `json:"ended"`
	BallPossessionTeam Team                  `json:"ball_possession_team"`
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
