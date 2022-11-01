package mts

import (
	"context"

	"github.com/databet-cloud/databet-go-sdk/pkg/restriction"
)

type Client interface {
	PlaceBet(ctx context.Context, req *PlaceBetRequest) (*Bet, []restriction.Restriction, error)
	DeclineBet(ctx context.Context, req *DeclineBetRequest) error

	CalculateCashOut(ctx context.Context, req *CalculateCashOutRequest) (*CashOutAmount, []restriction.Restriction, error)
	PlaceCashOutOrder(ctx context.Context, req *PlaceCashOutOrderRequest) (*Bet, *CashOutOrder, []restriction.Restriction, error)
	CancelCashOutOrder(ctx context.Context, req *CancelCashOutOrderRequest) error

	GetRestrictions(ctx context.Context, req *GetRestrictionsRequest) ([]restriction.Restriction, error)
	GetMaxBet(ctx context.Context, req *GetMaxBetRequest) (string, error)
}
