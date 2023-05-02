package statistics

const (
	TennisCommonSetStatus   SetStatus = "common"
	TennisTieBreakSetStatus SetStatus = "tie_break"
)

const (
	TennisServeTypeFirst  ServeAttemptType = 1
	TennisServeTypeSecond ServeAttemptType = 2
)

const (
	Point0     GamePoint = "0"
	Point15    GamePoint = "15"
	Point30    GamePoint = "30"
	Point40    GamePoint = "40"
	PointAbove GamePoint = "A"
)

func (s TennisStatistic) GetType() Type {
	return s.Type
}

type TennisStatistic struct {
	Type Type        `json:"type"`
	Sets []TennisSet `json:"sets"`
}

type SetStatus = string

type ServeAttemptType = int

type GamePoint = string

type TennisSet struct {
	Number   int       `json:"number"`
	Winner   Team      `json:"winner"`
	State    SetStatus `json:"state"`
	Games    []Game    `json:"games"`
	TieBreak TieBreak  `json:"tie_break"`
}

type Game struct {
	Number int    `json:"number"`
	Winner Team   `json:"winner"`
	Serve  Serve  `json:"serve"`
	Points Points `json:"points"`
}

type TieBreak struct {
	Number int            `json:"number"`
	Winner Team           `json:"winner"`
	Serve  Serve          `json:"serve"`
	Points TieBreakPoints `json:"points"`
}

type Serve struct {
	ServeAttempt ServeAttemptType `json:"serve_attempt"`
	Server       Team             `json:"server"`
}

type Points struct {
	Home GamePoint `json:"home"`
	Away GamePoint `json:"away"`
}

type TieBreakPoints struct {
	Home int `json:"home"`
	Away int `json:"away"`
}
