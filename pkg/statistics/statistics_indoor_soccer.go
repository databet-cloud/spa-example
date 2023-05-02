package statistics

type IndoorSoccerPeriodType string

const (
	IndoorSoccerPeriodUnknown    IndoorSoccerPeriodType = "unknown"
	IndoorSoccerPeriodFirstHalf  IndoorSoccerPeriodType = "1st_half"
	IndoorSoccerPeriodSecondHalf IndoorSoccerPeriodType = "2nd_half"
)

type IndoorSoccerTimeFormat string

const (
	IndoorSoccerTimeFormatUnknown       IndoorSoccerTimeFormat = "unknown"
	IndoorSoccerTimeFormatTenMinutes    IndoorSoccerTimeFormat = "10_min"
	IndoorSoccerTimeFormatTwelveMinutes IndoorSoccerTimeFormat = "12_min"
)

func (s IndoorSoccerStatistic) GetType() Type {
	return s.Type
}

type IndoorSoccerStatistic struct {
	Type       Type                   `json:"type"`
	TimeFormat IndoorSoccerTimeFormat `json:"time_format"`
	Periods    []IndoorSoccerPeriod   `json:"periods"`
}

type IndoorSoccerPeriod struct {
	Number int                    `json:"number"`
	Type   IndoorSoccerPeriodType `json:"type"`
	Ended  bool                   `json:"ended"`
	Goals  []IndoorSoccerGoal     `json:"goals"`
	Cards  []IndoorSoccerCard     `json:"cards"`
	Timer  Timer                  `json:"timer"`
}

type IndoorSoccerGoal struct {
	Number    int  `json:"number"`
	Team      Team `json:"team"`
	Cancelled bool `json:"cancelled"`
}

type IndoorSoccerCard struct {
	Number    int      `json:"number"`
	MatchTime int      `json:"match_time"`
	Type      CardType `json:"type"`
	Cancelled bool     `json:"cancelled"`
	Team      string   `json:"team"`
}
