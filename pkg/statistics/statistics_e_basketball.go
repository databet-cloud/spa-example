package statistics

type EbasketballQuarterFormat string

const (
	EbasketballQuarterFormatUnknown       EbasketballQuarterFormat = "unknown"
	EbasketballQuarterFormatFourMinutes   EbasketballQuarterFormat = "4_min"
	EbasketballQuarterFormatFiveMinutes   EbasketballQuarterFormat = "5_min"
	EbasketballQuarterFormatSixMinutes    EbasketballQuarterFormat = "6_min"
	EbasketballQuarterFormatTenMinutes    EbasketballQuarterFormat = "10_min"
	EbasketballQuarterFormatTwelveMinutes EbasketballQuarterFormat = "12_min"
)

func (s EBasketballStatistic) GetType() Type {
	return s.Type
}

type EBasketballStatistic struct {
	Type       Type                     `json:"type"`
	TimeFormat EbasketballQuarterFormat `json:"time_format"`
	Quarters   []EBasketballQuarter     `json:"quarters"`
}

type EBasketballQuarter struct {
	Number             int                    `json:"number"`
	Ended              bool                   `json:"ended"`
	BallPossessionTeam Team                   `json:"ball_possession_team"`
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
