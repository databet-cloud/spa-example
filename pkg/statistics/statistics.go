package statistics

import (
	"encoding/json"

	"github.com/bytedance/sonic"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

type Statistics []Statistic

type Statistic interface {
	GetType() Type
}

// easyjson:json
type typedStatistic struct {
	Type Type `json:"type"`
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
		case TypeCSGO:
			statistic = new(CSGOStatistic)
		case TypeDota2:
			statistic = new(Dota2Statistic)
		case TypeLOL:
			statistic = new(LOLStatistic)
		case TypeSoccer:
			statistic = new(SoccerStatistic)
		case TypeBasketball:
			statistic = new(BasketballStatistic)
		case TypeHockey:
			statistic = new(HockeyStatistic)
		case TypeEHockey:
			statistic = new(EHockeyStatistic)
		case TypeEBasketball:
			statistic = new(EBasketballStatistic)
		case TypeESoccer:
			statistic = new(ESoccerStatistic)
		case TypeTennis:
			statistic = new(TennisStatistic)
		case TypeETennis:
			statistic = new(ETennisStatistic)
		case TypeTableTennis:
			statistic = new(TableTennisStatistic)
		case TypeVolleyball:
			statistic = new(VolleyballStatistic)
		case TypeBeachVolleyball:
			statistic = new(BeachVolleyballStatistic)
		case TypeEVolleyball:
			statistic = new(EVolleyballStatistic)
		case TypeAmericanFootball:
			statistic = new(AmericanFootballStatistic)
		case TypeIndoorSoccer:
			statistic = new(IndoorSoccerStatistic)
		case TypeEUFC:
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
