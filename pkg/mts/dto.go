package mts

type AcceptStrategyID int8

const (
	AcceptStrategyDeviationAllowed AcceptStrategyID = 1 << iota
	AcceptStrategyGreaterAllowed
	AcceptStrategyAlwaysAllowed
)

type RestrictionsSelection struct {
	SportEventID string  `json:"sport_event_id"`
	MarketID     string  `json:"market_id"`
	OddID        string  `json:"odd_id"`
	Value        Decimal `json:"value"`
}
