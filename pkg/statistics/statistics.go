package statistics

import (
	"encoding/json"

	"github.com/bytedance/sonic"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

type Statistics []Statistic

type Statistic interface {
	Typ() string
}

// easyjson:json
type typedStatistic struct {
	Type string `json:"type"`
}

func (s *Statistics) UnmarshalJSON(data []byte) error {
	var sItems []json.RawMessage

	if err := json.Unmarshal(data, &sItems); err != nil {
		return err
	}

	result := make(Statistics, 0, len(sItems))

	var typed typedStatistic

	for _, sItem := range sItems {
		if err := json.Unmarshal(sItem, &typed); err != nil {
			return err
		}

		var statistic Statistic

		switch typed.Type {
		case SportCSGO:
			statistic = new(CSGOStatistic)
		case SportDota2:
			statistic = new(Dota2Statistic)
		case SportLOL:
			statistic = new(LOLStatistic)
		case SportSoccer:
			statistic = new(SoccerStatistic)
		case SportBasketball:
			statistic = new(BasketballStatistic)
		case SportHockey:
			statistic = new(HockeyStatistic)
		case SportEHockey:
			statistic = new(EHockeyStatistic)
		case SportEBasketball:
			statistic = new(EBasketballStatistic)
		case SportESoccer:
			statistic = new(ESoccerStatistic)
		case SportTennis:
			statistic = new(TennisStatistic)
		case SportETennis:
			statistic = new(ETennisStatistic)
		case SportTableTennis:
			statistic = new(TableTennisStatistic)
		case SportVolleyball:
			statistic = new(VolleyballStatistic)
		case SportBeachVolleyball:
			statistic = new(BeachVolleyballStatistic)
		case SportEVolleyball:
			statistic = new(EVolleyballStatistic)
		case SportAmericanFootball:
			statistic = new(AmericanFootballStatistic)
		case SportIndoorSoccer:
			statistic = new(IndoorSoccerStatistic)
		case SportEUFC:
			statistic = new(EUFCStatistic)
		default:
			// append raw json for unknown sport
			result = append(result, RawStatistic{typ: typed.Type, RawMessage: sItem})
			continue
		}

		if err := sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(sItem), statistic); err != nil {
			return err
		}

		result = append(result, statistic)

	}

	*s = result

	return nil
}
