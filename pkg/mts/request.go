package mts

import "time"

type PlaceBetRequest struct {
	BetID             string                     `json:"bet_id"`
	BetType           PlaceBetRequestBetType     `json:"type"`
	Stake             PlaceBetRequestStake       `json:"stake"`
	Selections        []PlaceBetRequestSelection `json:"selections"`
	OddAcceptStrategy AcceptStrategyID           `json:"odd_accept_strategy"`
	PlayerID          string                     `json:"player_id"`
	ClientIP          string                     `json:"client_ip"`
	// CountryCode ISO 3166-1 alpha-2
	CountryCode string    `json:"country_code"`
	CreatedAt   time.Time `json:"created_at"`
}

type PlaceBetRequestBetType struct {
	Code int   `json:"code"`
	Size []int `json:"size"`
}

type PlaceBetRequestStake struct {
	Value float64 `json:"value"`
	// CurrencyCode ISO 4217 (alfa-3)
	CurrencyCode string `json:"currency_code"`
}

type PlaceBetRequestSelection struct {
	SportEventID string  `json:"sport_event_id"`
	MarketID     string  `json:"market_id"`
	OddID        string  `json:"odd_id"`
	Value        float64 `json:"value"`
}

type DeclineBetRequest struct {
	BetID  string `json:"bet_id"`
	Reason string `json:"reason"`
}

type CalculateCashOutRequest struct {
	BetID      string             `json:"-"`
	Amount     Money              `json:"amount"`
	Selections []CashOutSelection `json:"selections"`
}

type CancelCashOutOrderRequest struct {
	BetID          string          `json:"foreign_id"`
	CashOutOrderID string          `json:"cash_out_order_id"`
	Context        *CashOutContext `json:"context"`
}

type AcceptCashOutOrderRequest struct {
	Context *CashOutContext `json:"context"`
}

type DeclineCashOutOrderRequest struct {
	Context *CashOutContext `json:"context"`
}

type PlaceCashOutOrderRequest struct {
	ID           string             `json:"id"`
	BetID        string             `json:"-"`
	Amount       Money              `json:"amount"`
	RefundAmount Money              `json:"refund_amount"`
	CreatedAt    string             `json:"created_at"`
	Selections   []CashOutSelection `json:"selections"`
}

type GetRestrictionsRequest struct {
	PlayerID   string
	BetType    int
	Selections []RestrictionsSelection
	// SystemSizes: every element in array defines odds grouping. For Singles it's always [1],
	// for express - [len(BetPlaceRequest.Selections)], for system - user defined size.
	SystemSizes []int
	// CurrencyCode ISO 4217 (alfa-3)
	CurrencyCode      string
	OddAcceptStrategy AcceptStrategyID
}
