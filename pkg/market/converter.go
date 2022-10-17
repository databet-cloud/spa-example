package market

type StatusMarketConverter struct{}

func NewMarketStatusConverter() *StatusMarketConverter {
	return &StatusMarketConverter{}
}

func (smc *StatusMarketConverter) ConvertMarket(m Market) StatusMarket {
	statusOdds := make(StatusOddCollection, len(m.Odds))
	for id, odd := range m.Odds {
		statusOdds[id] = StatusOdd{
			ID:           odd.ID,
			Status:       odd.Status,
			StatusReason: odd.StatusReason,
		}
	}

	return StatusMarket{
		ID:     m.ID,
		TypeID: m.TypeID,
		Status: m.Status,
		Odds:   statusOdds,
	}
}

func (smc *StatusMarketConverter) ConvertCollection(c Collection) StatusMarketCollection {
	statusCollection := make(StatusMarketCollection, len(c))
	for id, m := range c {
		statusCollection[id] = smc.ConvertMarket(m)
	}

	return statusCollection
}
