package mts

import (
	apierror "github.com/databet-cloud/databet-go-sdk/pkg/error"
	"github.com/databet-cloud/databet-go-sdk/pkg/restriction"
)

type errorResponse struct {
	Error *apierror.Error `json:"error,omitempty"`
}

type placeBetResponse struct {
	Bet          *Bet                      `json:"bet,omitempty"`
	Error        *apierror.Error           `json:"error,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
}

type PlaceBetResponse struct {
	Bet          *Bet                      `json:"bet,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
}

type calculateCashOutResponse struct {
	Amount       *CashOutAmount            `json:"amount,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
	Error        *apierror.Error           `json:"error,omitempty"`
}

type CalculateCashOutResponse struct {
	Amount       *CashOutAmount            `json:"amount,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
}

type placeCashOutOrderResponse struct {
	Bet          *Bet                      `json:"bet,omitempty"`
	CashOutOrder *CashOutOrder             `json:"cash_out_order,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
	Error        *apierror.Error           `json:"error,omitempty"`
}

type PlaceCashOutOrderResponse struct {
	Bet          *Bet                      `json:"bet,omitempty"`
	CashOutOrder *CashOutOrder             `json:"cash_out_order,omitempty"`
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
}

type getRestrictionsResponse struct {
	Restrictions []restriction.Restriction `json:"restrictions,omitempty"`
	Error        *apierror.Error           `json:"error,omitempty"`
}

type getMaxBetResponse struct {
	MaxBet string          `json:"max_bet"`
	Error  *apierror.Error `json:"error,omitempty"`
}
