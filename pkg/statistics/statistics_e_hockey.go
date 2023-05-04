package statistics

type EHockeyPeriodType string

const (
	EHockeyPeriodRegular  EHockeyPeriodType = "regular"
	EHockeyPeriodOvertime EHockeyPeriodType = "overtime"
	EHockeyPeriodBullets  EHockeyPeriodType = "bullets"
)

type EHockeyPeriodDuration string

const (
	EHockeyPeriodDurationTwentyMin EHockeyPeriodDuration = "20_min"
)

type EHockeyPeriodOvertimeDuration string

const (
	EHockeyPeriodOvertimeDurationFiveMin EHockeyPeriodOvertimeDuration = "5_min"
)

func (s EHockeyStatistic) GetType() Type {
	return s.Type
}

type EHockeyStatistic struct {
	Type                   Type                          `json:"type"`
	OvertimePeriodDuration EHockeyPeriodOvertimeDuration `json:"overtime_period_duration"`
	PeriodDuration         EHockeyPeriodDuration         `json:"period_duration"`
	Periods                []EHockeyPeriod               `json:"periods"`
}

type EHockeyPeriod struct {
	Number       int                  `json:"number"`
	Type         EHockeyPeriodType    `json:"type"`
	Ended        bool                 `json:"ended"`
	Timer        Timer                `json:"timer"`
	Points       []EHockeyPoint       `json:"points"`
	BulletThrows []EHockeyBulletThrow `json:"bullet_throws"`
}

type EHockeyPoint struct {
	Number   int    `json:"number"`
	Team     string `json:"team"`
	Canceled bool   `json:"canceled"`
}

type EHockeyBulletThrow struct {
	Number   int  `json:"number"`
	Team     Team `json:"team"`
	Canceled bool `json:"canceled"`
	Active   bool `json:"active"`
	Scored   bool `json:"scored"`
}