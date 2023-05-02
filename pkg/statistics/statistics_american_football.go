package statistics

type AmericanFootballPeriodTimeFormat string

const (
	AmericanFootballPeriodTimeFormatUnknown        AmericanFootballPeriodTimeFormat = "unknown"
	AmericanFootballPeriodTimeFormatTwelveMinutes  AmericanFootballPeriodTimeFormat = "12_min"
	AmericanFootballPeriodTimeFormatFifteenMinutes AmericanFootballPeriodTimeFormat = "15_min"
)

func (s AmericanFootballStatistic) GetType() Type {
	return s.Type
}

type AmericanFootballStatistic struct {
	Type       Type                             `json:"type"`
	TimeFormat AmericanFootballPeriodTimeFormat `json:"time_format"`
	Periods    []AmericanFootballPeriod         `json:"periods"`
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
