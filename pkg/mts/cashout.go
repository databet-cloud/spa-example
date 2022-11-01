package mts

import (
	"time"

	"github.com/databet-cloud/databet-go-sdk/pkg/restriction"
)

type CashOutMoney struct {
	Value        string `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type CashOutSelection struct {
	SportEventID string `json:"sport_event_id"`
	MarketID     string `json:"market_id"`
	OddID        string `json:"odd_id"`
	Value        string `json:"value"`
}

type CashOutOrderSelection struct {
	OddID        string `json:"odd_id"`
	MarketID     string `json:"market_id"`
	SportEventID string `json:"sport_event_id"`
	Value        string `json:"value"`
}

type CashOutOrder struct {
	ID           string                  `json:"id"`
	BetID        string                  `json:"bet_id"`
	Amount       MultiMoney              `json:"amount"`
	RefundAmount MultiMoney              `json:"refund_amount"`
	Status       CashOutOrderStatus      `json:"status"`
	CreatedAt    time.Time               `json:"created_at"`
	UpdatedAt    time.Time               `json:"updated_at"`
	Selections   []CashOutOrderSelection `json:"selections"`
	Context      *CashOutContext         `json:"context"`
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
	Restrictions []restriction.Restriction `json:"restrictions"`
}

type CalculatedCashOut struct {
	Amount       CashOutAmount             `json:"amount"`
	Restrictions []restriction.Restriction `json:"restrictions"`
}

type CashOutAmount struct {
	RefundAmount    string         `json:"refund_amount"`
	MinAmount       string         `json:"min_amount"`
	MinRefundAmount string         `json:"min_refund_amount"`
	MaxAmount       string         `json:"max_amount"`
	MaxRefundAmount string         `json:"max_refund_amount"`
	Ranges          []CashOutRange `json:"ranges"`
}

type CashOutRange struct {
	FromAmount  string `json:"from_amount"`
	ToAmount    string `json:"to_amount"`
	RefundRatio string `json:"refund_ratio"`
}
