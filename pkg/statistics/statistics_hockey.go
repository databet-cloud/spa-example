package statistics

type HockeyPeriodType string

const (
	HockeyPeriodRegular  HockeyPeriodType = "regular"
	HockeyPeriodOvertime HockeyPeriodType = "overtime"
	HockeyPeriodBullets  HockeyPeriodType = "bullets"
)

type HockeyPeriodOvertimeDuration string

const (
	HockeyPeriodOvertimeDurationUnknown        HockeyPeriodOvertimeDuration = "unknown"
	HockeyPeriodOvertimeDurationNotAllowed     HockeyPeriodOvertimeDuration = "not_allowed"
	HockeyPeriodOvertimeDurationUntilFirstGoal HockeyPeriodOvertimeDuration = "until_first_goal"
	HockeyPeriodOvertimeDurationFiveMin        HockeyPeriodOvertimeDuration = "5_min"
	HockeyPeriodOvertimeDurationSevenMin       HockeyPeriodOvertimeDuration = "7_min"
	HockeyPeriodOvertimeDurationTenMin         HockeyPeriodOvertimeDuration = "10_min"
	HockeyPeriodOvertimeDurationTwentyMin      HockeyPeriodOvertimeDuration = "20_min"
)

type HockeyPeriodDuration string

const (
	HockeyPeriodDurationUnknown   = "unknown"
	HockeyPeriodDurationThreeMin  = "3_min"
	HockeyPeriodDurationFourMin   = "4_min"
	HockeyPeriodDurationSevenMin  = "7_min"
	HockeyPeriodDurationTenMin    = "10_min"
	HockeyPeriodDurationTwentyMin = "20_min"
)

func (s HockeyStatistic) GetType() Type {
	return s.Type
}

type HockeyStatistic struct {
	Type                   Type                         `json:"type"`
	OvertimePeriodDuration HockeyPeriodOvertimeDuration `json:"overtime_period_duration"`
	PeriodDuration         HockeyPeriodDuration         `json:"period_duration"`
	Periods                []HockeyPeriod               `json:"periods"`
}

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
