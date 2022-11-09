package mts

import (
	"time"

	"github.com/databet-cloud/databet-go-sdk/pkg/restriction"
)

type CashOutSelection struct {
	SportEventID string  `json:"sport_event_id"`
	MarketID     string  `json:"market_id"`
	OddID        string  `json:"odd_id"`
	Value        float64 `json:"value,string"`
}

type CashOutOrderSelection struct {
	OddID        string  `json:"odd_id"`
	MarketID     string  `json:"market_id"`
	SportEventID string  `json:"sport_event_id"`
	Value        float64 `json:"value,string"`
}

type CashOutOrderID string

func (o CashOutOrderID) String() string {
	return string(o)
}

type CashOutOrder struct {
	ID           CashOutOrderID          `json:"id"`
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
	RefundAmount    float64        `json:"refund_amount,string"`
	MinAmount       float64        `json:"min_amount,string"`
	MinRefundAmount float64        `json:"min_refund_amount,string"`
	MaxAmount       float64        `json:"max_amount,string"`
	MaxRefundAmount float64        `json:"max_refund_amount,string"`
	Ranges          []CashOutRange `json:"ranges"`
}

type CashOutRange struct {
	FromAmount  float64 `json:"from_amount,string"`
	ToAmount    float64 `json:"to_amount,string"`
	RefundRatio float64 `json:"refund_ratio,string"`
}
