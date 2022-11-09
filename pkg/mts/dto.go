package mts

type AcceptStrategyID int8

const (
	AcceptStrategyDeviationAllowed AcceptStrategyID = 1 << iota
	AcceptStrategyGreaterAllowed
	AcceptStrategyAlwaysAllowed
)

type MaxBetSelection struct {
	SportEventID string  `json:"sport_event_id"`
	MarketID     string  `json:"market_id"`
	OddID        string  `json:"odd_id"`
	Value        float64 `json:"value,string"`
}

type RestrictionsSelection struct {
	SportEventID string  `json:"sport_event_id"`
	MarketID     string  `json:"market_id"`
	OddID        string  `json:"odd_id"`
	Value        float64 `json:"value,string"`
}
