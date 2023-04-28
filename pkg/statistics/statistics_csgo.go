package statistics

const (
	CSGOSideT  = "t"
	CSGOSideCT = "ct"

	HomeWinner = "home"
	AwayWinner = "away"

	RoundTime = 60 + 55
)

const (
	CSGOMapNameDeInferno     = "de_inferno"
	CSGOMapNameDeMirage      = "de_mirage"
	CSGOMapNameDeVertigo     = "de_vertigo"
	CSGOMapNameDeOverpass    = "de_overpass"
	CSGOMapNameDeNuke        = "de_nuke"
	CSGOMapNameDeDust2       = "de_dust2"
	CSGOMapNameDeTrain       = "de_train"
	CSGOMapNameDeCache       = "de_cache"
	CSGOMapNameDeAncient     = "de_ancient"
	CSGOMapNameDeCobblestone = "de_cobblestone"
	CSGOMapNameDeAnubis      = "de_anubis"
)

func (s CSGOStatistic) Typ() string {
	return s.Type
}

type CSGOStatistic struct {
	Type string    `json:"type"`
	Maps []CSGOMap `json:"maps"`
}

type CSGOMap struct {
	Number int          `json:"number"`
	Name   string       `json:"name"`
	Rounds []CSGORound  `json:"rounds"`
	Score  CSGOMapScore `json:"score"`
	Winner Team         `json:"winner"`
}

type CSGOMapScore struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type CSGORound struct {
	Number       int    `json:"number"`
	Timer        Timer  `json:"timer"`
	HomeTeamSide string `json:"home_team_side"`
	AwayTeamSide string `json:"away_team_side"`
	BombPlanted  bool   `json:"bomb_planted"`
	BombTime     int    `json:"bomb_time"`
	GameState    string `json:"game_state"`
}
