package statistics

type ESoccerPeriodType string

const (
	ESoccerPeriodUnknown     ESoccerPeriodType = "unknown"
	ESoccerPeriodFirstHalf   ESoccerPeriodType = "1st_half"
	ESoccerPeriodSecondHalf  ESoccerPeriodType = "2nd_half"
	ESoccerPeriodFirstExtra  ESoccerPeriodType = "1st_extra"
	ESoccerPeriodSecondExtra ESoccerPeriodType = "2nd_extra"
	ESoccerPeriodPenalties   ESoccerPeriodType = "penalties"
)

func (s ESoccerStatistic) GetType() Type {
	return s.Type
}

type ESoccerStatistic struct {
	Type    Type            `json:"type"`
	Periods []ESoccerPeriod `json:"periods"`
}

type ESoccerGoal struct {
	Number   int  `json:"number"`
	Canceled bool `json:"canceled"`
	Team     Team `json:"team"`
}

type ESoccerCard struct {
	Number   int      `json:"number"`
	Type     CardType `json:"type"`
	Canceled bool     `json:"canceled"`
	Team     Team     `json:"team"`
}

type ESoccerPenalty struct {
	Number    int           `json:"number"`
	MatchTime int           `json:"match_time"`
	Status    PenaltyStatus `json:"status"`
	Team      Team          `json:"team"`
}

type ESoccerPeriod struct {
	Name      string            `json:"name"`
	Type      ESoccerPeriodType `json:"type"`
	IsEnded   bool              `json:"is_ended"`
	Timer     Timer             `json:"timer"`
	Cards     []ESoccerCard     `json:"cards"`
	Goals     []ESoccerGoal     `json:"goals"`
	Penalties []ESoccerPenalty  `json:"penalties"`
}
