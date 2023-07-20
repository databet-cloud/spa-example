package statistics

type CSGOSide string

const (
	CSGOSideT  CSGOSide = "t"
	CSGOSideCT CSGOSide = "ct"
)

const (
	RoundTime = 60 + 55
)

type CSGOMapName string

const (
	CSGOMapNameDeInferno     CSGOMapName = "de_inferno"
	CSGOMapNameDeMirage      CSGOMapName = "de_mirage"
	CSGOMapNameDeVertigo     CSGOMapName = "de_vertigo"
	CSGOMapNameDeOverpass    CSGOMapName = "de_overpass"
	CSGOMapNameDeNuke        CSGOMapName = "de_nuke"
	CSGOMapNameDeDust2       CSGOMapName = "de_dust2"
	CSGOMapNameDeTrain       CSGOMapName = "de_train"
	CSGOMapNameDeCache       CSGOMapName = "de_cache"
	CSGOMapNameDeAncient     CSGOMapName = "de_ancient"
	CSGOMapNameDeCobblestone CSGOMapName = "de_cobblestone"
	CSGOMapNameDeAnubis      CSGOMapName = "de_anubis"
)

func (s CSGOStatistic) GetType() Type {
	return s.Type
}

type CSGOStatistic struct {
	Type             Type                  `json:"type"`
	Maps             []CSGOMap             `json:"maps"`
	PlayersStatistic []CSGOPlayerStatistic `json:"players"`
}

type CSGOMap struct {
	Number int          `json:"number"`
	Name   CSGOMapName  `json:"name"`
	Rounds []CSGORound  `json:"rounds"`
	Score  CSGOMapScore `json:"score"`
	Winner Team         `json:"winner"`
}

type CSGOMapScore struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type CSGORound struct {
	Number       int      `json:"number"`
	Timer        Timer    `json:"timer"`
	HomeTeamSide CSGOSide `json:"home_team_side"`
	AwayTeamSide CSGOSide `json:"away_team_side"`
	BombPlanted  bool     `json:"bomb_planted"`
	BombTime     int      `json:"bomb_time"`
	GameState    string   `json:"game_state"`
}

type CSGOPlayerStatistic struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}
