package statistics

type SoccerPeriodType string

const (
	SoccerPeriodUnknown     SoccerPeriodType = "unknown"
	SoccerPeriodFirstHalf   SoccerPeriodType = "1st_half"
	SoccerPeriodSecondHalf  SoccerPeriodType = "2nd_half"
	SoccerPeriodFirstExtra  SoccerPeriodType = "1st_extra"
	SoccerPeriodSecondExtra SoccerPeriodType = "2nd_extra"
	SoccerPeriodPenalties   SoccerPeriodType = "penalties"
	SoccerPeriodLast        SoccerPeriodType = "last"
)

func (s SoccerStatistic) GetType() Type {
	return s.Type
}

type SoccerStatistic struct {
	Type    Type                    `json:"type"`
	Periods map[string]SoccerPeriod `json:"periods"`
}

type SoccerPlayer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SoccerPenalty struct {
	Number    int           `json:"number"`
	MatchTime int           `json:"match_time"`
	Status    PenaltyStatus `json:"status"`
	Player    SoccerPlayer  `json:"player"`
}

type SoccerGoal struct {
	Number   int  `json:"number"`
	Canceled bool `json:"canceled"`
}

type SoccerTeams struct {
	Home SoccerTeam `json:"home"`
	Away SoccerTeam `json:"away"`
}

type SoccerTeam struct {
	Goals       []SoccerGoal    `json:"goals"`
	Penalties   []SoccerPenalty `json:"penalties"`
	Corners     int             `json:"corners"`
	YellowCards int             `json:"yellow_cards"`
	RedCards    int             `json:"red_cards"`
}

type SoccerPeriod struct {
	Name      string           `json:"name"`
	Type      SoccerPeriodType `json:"type"`
	IsEnded   bool             `json:"is_ended"`
	Teams     SoccerTeams      `json:"teams"`
	VARActive bool             `json:"video_assistant_referee_active"`
	Timer     Timer            `json:"timer"`
}
