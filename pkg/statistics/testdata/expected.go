package testdata

import "github.com/databet-cloud/databet-go-sdk/pkg/statistics"

var ExpectedStatisticsResponse = statistics.Statistics{
	&statistics.CSGOStatistic{
		Type: "esports_counter_strike",
		Maps: []statistics.CSGOMap{
			{
				Number: 1,
				Name:   "unknown",
				Rounds: []statistics.CSGORound{
					{
						Number: 1,
						Timer: statistics.Timer{
							IsActive:  false,
							StartedAt: 0,
							EndedAt:   0,
							TimeDelta: 25000,
						},
						HomeTeamSide: "t",
						AwayTeamSide: "ct",
						BombPlanted:  false,
						BombTime:     0,
						GameState:    "after_end_time",
					},
				},
				Score: statistics.CSGOMapScore{
					Home: 16,
					Away: 5,
				},
				Winner: statistics.Home,
			},
		},
	},
}
