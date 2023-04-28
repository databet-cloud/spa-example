package statistics

const (
	IndoorSoccerPeriodUnknown           = "unknown"
	IndoorSoccerPeriodFirstHalf         = "1st_half"
	IndoorSoccerPeriodSecondHalf        = "2nd_half"
	IndoorSoccerTimeFormatUnknown       = "unknown"
	IndoorSoccerTimeFormatTenMinutes    = "10_min"
	IndoorSoccerTimeFormatTwelveMinutes = "12_min"
)

func (s IndoorSoccerStatistic) Typ() string {
	return s.Type
}

type IndoorSoccerStatistic struct {
	Type       string               `json:"type"`
	TimeFormat string               `json:"time_format"`
	Periods    []IndoorSoccerPeriod `json:"periods"`
}

type IndoorSoccerPeriod struct {
	Number int                `json:"number"`
	Type   string             `json:"type"`
	Ended  bool               `json:"ended"`
	Goals  []IndoorSoccerGoal `json:"goals"`
	Cards  []IndoorSoccerCard `json:"cards"`
	Timer  Timer              `json:"timer"`
}

type IndoorSoccerGoal struct {
	Number    int  `json:"number"`
	Team      Team `json:"team"`
	Cancelled bool `json:"cancelled"`
}

type IndoorSoccerCard struct {
	Number    int    `json:"number"`
	MatchTime int    `json:"match_time"`
	Type      string `json:"type"`
	Cancelled bool   `json:"cancelled"`
	Team      string `json:"team"`
}
