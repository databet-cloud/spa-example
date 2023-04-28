package statistics

const (
	ESoccerCardTypeRed       = "red"
	ESoccerCardTypeYellow    = "yellow"
	ESoccerCardTypeYellowRed = "yellow_red"
)

const (
	ESoccerPeriodUnknown     = "unknown"
	ESoccerPeriodFirstHalf   = "1st_half"
	ESoccerPeriodSecondHalf  = "2nd_half"
	ESoccerPeriodFirstExtra  = "1st_extra"
	ESoccerPeriodSecondExtra = "2nd_extra"
	ESoccerPeriodPenalties   = "penalties"
)

func (s ESoccerStatistic) Typ() string {
	return s.Type
}

type ESoccerStatistic struct {
	Type    string          `json:"type"`
	Periods []ESoccerPeriod `json:"periods"`
}

type ESoccerGoal struct {
	Number   int  `json:"number"`
	Canceled bool `json:"canceled"`
	Team     Team `json:"team"`
}

type ESoccerCard struct {
	Number   int    `json:"number"`
	Type     string `json:"type"`
	Canceled bool   `json:"canceled"`
	Team     Team   `json:"team"`
}

type ESoccerPenalty struct {
	Number    int           `json:"number"`
	MatchTime int           `json:"match_time"`
	Status    PenaltyStatus `json:"status"`
	Team      Team          `json:"team"`
}

type ESoccerPeriod struct {
	Name      string           `json:"name"`
	IsEnded   bool             `json:"is_ended"`
	Timer     Timer            `json:"timer"`
	Cards     []ESoccerCard    `json:"cards"`
	Goals     []ESoccerGoal    `json:"goals"`
	Penalties []ESoccerPenalty `json:"penalties"`
}
