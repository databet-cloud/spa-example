package statistics

const (
	SoccerPeriodUnknown     = "unknown"
	SoccerPeriodFirstHalf   = "1st_half"
	SoccerPeriodSecondHalf  = "2nd_half"
	SoccerPeriodFirstExtra  = "1st_extra"
	SoccerPeriodSecondExtra = "2nd_extra"
	SoccerPeriodPenalties   = "penalties"
	SoccerPeriodLast        = "last"
)

func (s SoccerStatistic) Typ() string {
	return s.Type
}

type SoccerStatistic struct {
	Type    string                  `json:"type"`
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
	Name      string      `json:"name"`
	IsEnded   bool        `json:"is_ended"`
	Teams     SoccerTeams `json:"teams"`
	VARActive bool        `json:"video_assistant_referee_active"`
	Timer     Timer       `json:"timer"`
}
