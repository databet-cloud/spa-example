package statistics

func (s ETennisStatistic) GetType() Type {
	return s.Type
}

type ETennisStatistic struct {
	Type Type         `json:"type"`
	Sets []ETennisSet `json:"sets"`
}

type ETennisSet struct {
	Number   int             `json:"number"`
	Winner   Team            `json:"winner"`
	State    SetStatus       `json:"state"`
	Games    []ETennisGame   `json:"games"`
	TieBreak ETennisTieBreak `json:"tie_break"`
}

type ETennisGame struct {
	Number int           `json:"number"`
	Winner Team          `json:"winner"`
	Serve  ETennisServe  `json:"serve"`
	Points ETennisPoints `json:"points"`
}

type ETennisTieBreak struct {
	Number int                   `json:"number"`
	Winner Team                  `json:"winner"`
	Serve  ETennisServe          `json:"serve"`
	Points ETennisTieBreakPoints `json:"points"`
}

type ETennisServe struct {
	ServeAttempt ServeAttemptType `json:"serve_attempt"`
	Server       Team             `json:"server"`
}

type ETennisPoints struct {
	Home GamePoint `json:"home"`
	Away GamePoint `json:"away"`
}

type ETennisTieBreakPoints struct {
	Home int `json:"home"`
	Away int `json:"away"`
}
