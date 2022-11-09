package mts

import (
	"github.com/databet-cloud/databet-go-sdk/pkg/restriction"
)

type errorResponse struct {
	Error *apiError `json:"error,omitempty"`
}

type placeBetResponse struct {
	Bet          *Bet                      `json:"bet,omitempty"`
	Error        *apiError                 `json:"error,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
}

type PlaceBetResponse struct {
	Bet          *Bet                      `json:"bet,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
}

type calculateCashOutResponse struct {
	Amount       *CashOutAmount            `json:"amount,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
	Error        *apiError                 `json:"error,omitempty"`
}

type CalculateCashOutResponse struct {
	Amount       *CashOutAmount            `json:"amount,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
}

type placeCashOutOrderResponse struct {
	Bet          *Bet                      `json:"bet,omitempty"`
	CashOutOrder *CashOutOrder             `json:"cash_out_order,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
	Error        *apiError                 `json:"error,omitempty"`
}

type PlaceCashOutOrderResponse struct {
	Bet          *Bet                      `json:"bet,omitempty"`
	CashOutOrder *CashOutOrder             `json:"cash_out_order,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
}

type getRestrictionsResponse struct {
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
	Error        *apiError                 `json:"error,omitempty"`
}

type getMaxBetResponse struct {
	MaxBet float64   `json:"max_bet,string"`
	Error  *apiError `json:"error,omitempty"`
}
