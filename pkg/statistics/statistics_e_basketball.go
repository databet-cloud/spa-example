package statistics

const (
	EbasketballQuarterFormatUnknown       = "unknown"
	EbasketballQuarterFormatFourMinutes   = "4_min"
	EbasketballQuarterFormatFiveMinutes   = "5_min"
	EbasketballQuarterFormatSixMinutes    = "6_min"
	EbasketballQuarterFormatTenMinutes    = "10_min"
	EbasketballQuarterFormatTwelveMinutes = "12_min"
)

func (s EBasketballStatistic) Typ() string {
	return s.Type
}

type EBasketballStatistic struct {
	Type       string               `json:"type"`
	TimeFormat string               `json:"time_format"`
	Quarters   []EBasketballQuarter `json:"quarters"`
}

type EBasketballQuarter struct {
	Number             int                    `json:"number"`
	Ended              bool                   `json:"ended"`
	BallPossessionTeam string                 `json:"ball_possession_team"`
	Timer              Timer                  `json:"timer"`
	Score              EBasketballScore       `json:"score"`
	FoulThrows         []EBasketballFoulThrow `json:"foul_throws"`
}

type EBasketballFoulThrow struct {
	Number   int  `json:"number"`
	Team     Team `json:"team"`
	Canceled bool `json:"canceled"`
}

type EBasketballScore struct {
	Home int `json:"home"`
	Away int `json:"away"`
}
