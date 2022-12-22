package mts

import (
	"context"
)

type Client interface {
	PlaceBet(ctx context.Context, req *PlaceBetRequest) (*PlaceBetResponse, error)
	DeclineBet(ctx context.Context, req *DeclineBetRequest) error

	CalculateCashOut(ctx context.Context, req *CalculateCashOutRequest) (*CalculateCashOutResponse, error)
	PlaceCashOutOrder(ctx context.Context, req *PlaceCashOutOrderRequest) (*PlaceCashOutOrderResponse, error)
	CancelCashOutOrder(ctx context.Context, req *CancelCashOutOrderRequest) error

	GetRestrictions(ctx context.Context, req *GetRestrictionsRequest) ([]Restriction, error)
}
