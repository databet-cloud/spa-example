package mts

import (
	"time"
)

type CashOutOrderSelection struct {
	OddID        string  `json:"odd_id"`
	MarketID     string  `json:"market_id"`
	SportEventID string  `json:"sport_event_id"`
	Value        Decimal `json:"value"`
}

type CashOutOrderID string

func (o CashOutOrderID) String() string {
	return string(o)
}

type CashOutOrder struct {
	// ID of the cash-out sent by the PlaceCashOutOrder request
	ID CashOutOrderID `json:"id"`
	// BetID sent by the PlaceBet request
	BetID string `json:"bet_id"`
	// Amount is the part of the bet stake that was cashed out
	Amount MultiMoney `json:"amount"`
	// RefundAmount of money that should be credited to the player`s balance
	RefundAmount MultiMoney         `json:"refund_amount"`
	Status       CashOutOrderStatus `json:"status"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	// Selections state at the moment of cash-out - must contain all selections in the bet,
	// but have an actual value(coefficients)
	Selections []CashOutOrderSelection `json:"selections"`
	Context    *CashOutContext         `json:"context"`
}

type CashOutOrderStatus struct {
	Code   CashOutOrderStatusCode `json:"code"`
	Reason string                 `json:"reason"`
}

type CashOutOrderStatusCode string

const (
	CashOutOrderStatusPending  CashOutOrderStatusCode = "pending"
	CashOutOrderStatusAccepted CashOutOrderStatusCode = "accepted"
	CashOutOrderStatusDeclined CashOutOrderStatusCode = "declined"
)

type CashOutContext struct {
	Restrictions []Restriction `json:"restrictions"`
}

type CalculatedCashOut struct {
	Amount       CashOutAmount `json:"amount"`
	Restrictions []Restriction `json:"restrictions"`
}

type CashOutAmount struct {
	RefundAmount    Decimal        `json:"refund_amount"`
	MinAmount       Decimal        `json:"min_amount"`
	MinRefundAmount Decimal        `json:"min_refund_amount"`
	MaxAmount       Decimal        `json:"max_amount"`
	MaxRefundAmount Decimal        `json:"max_refund_amount"`
	Ranges          []CashOutRange `json:"ranges"`
}

type CashOutRange struct {
	FromAmount  Decimal `json:"from_amount"`
	ToAmount    Decimal `json:"to_amount"`
	RefundRatio Decimal `json:"refund_ratio"`
}
