package statistics

const (
	AmericanFootballPeriodTimeFormatUnknown        = "unknown"
	AmericanFootballPeriodTimeFormatTwelveMinutes  = "12_min"
	AmericanFootballPeriodTimeFormatFifteenMinutes = "15_min"
)

func (s AmericanFootballStatistic) Typ() string {
	return s.Type
}

type AmericanFootballStatistic struct {
	Type       string                   `json:"type"`
	TimeFormat string                   `json:"time_format"`
	Periods    []AmericanFootballPeriod `json:"periods"`
}

type AmericanFootballPeriod struct {
	Number     int                   `json:"number"`
	Ended      bool                  `json:"ended"`
	AttackSide Team                  `json:"attack_side"`
	Score      AmericanFootballScore `json:"score"`
	Timer      Timer                 `json:"timer"`
}

type AmericanFootballScore struct {
	Home   int                     `json:"home"`
	Away   int                     `json:"away"`
	Points []AmericanFootballPoint `json:"points"`
}

type AmericanFootballPoint struct {
	Number int  `json:"number"`
	Team   Team `json:"team"`
	Value  int  `json:"value"`
	Time   int  `json:"time"`
}
