package statistics

const (
	HockeyPeriodRegular  HockeyPeriodType = "regular"
	HockeyPeriodOvertime HockeyPeriodType = "overtime"
	HockeyPeriodBullets  HockeyPeriodType = "bullets"

	HockeyPeriodOvertimeDurationUnknown        = "unknown"
	HockeyPeriodOvertimeDurationNotAllowed     = "not_allowed"
	HockeyPeriodOvertimeDurationUntilFirstGoal = "until_first_goal"
	HockeyPeriodOvertimeDurationFiveMin        = "5_min"
	HockeyPeriodOvertimeDurationSevenMin       = "7_min"
	HockeyPeriodOvertimeDurationTenMin         = "10_min"
	HockeyPeriodOvertimeDurationTwentyMin      = "20_min"

	HockeyPeriodDurationUnknown   = "unknown"
	HockeyPeriodDurationThreeMin  = "3_min"
	HockeyPeriodDurationFourMin   = "4_min"
	HockeyPeriodDurationSevenMin  = "7_min"
	HockeyPeriodDurationTenMin    = "10_min"
	HockeyPeriodDurationTwentyMin = "20_min"
)

func (s HockeyStatistic) Typ() string {
	return s.Type
}

type HockeyStatistic struct {
	Type                   string         `json:"type"`
	OvertimePeriodDuration string         `json:"overtime_period_duration"`
	PeriodDuration         string         `json:"period_duration"`
	Periods                []HockeyPeriod `json:"periods"`
}

type HockeyPeriodType = string

type HockeyPeriod struct {
	Number       int                 `json:"number"`
	Type         HockeyPeriodType    `json:"type"`
	Ended        bool                `json:"ended"`
	Timer        Timer               `json:"timer"`
	Points       []HockeyPoint       `json:"points"`
	BulletThrows []HockeyBulletThrow `json:"bullet_throws"`
}

type HockeyPoint struct {
	Number   int    `json:"number"`
	Team     string `json:"team"`
	Canceled bool   `json:"canceled"`
}

type HockeyBulletThrow struct {
	Number   int  `json:"number"`
	Team     Team `json:"team"`
	Canceled bool `json:"canceled"`
	Active   bool `json:"active"`
	Scored   bool `json:"scored"`
}
