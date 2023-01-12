package mts

import (
	"context"
)

type Client interface {
	// PlaceBet notifies that you've placed the bet on your side
	PlaceBet(ctx context.Context, req *PlaceBetRequest) (*PlaceBetResponse, error)

	// DeclineBet notifies that you've declined the bet that was previously accepted by the trading bet accounting system
	DeclineBet(ctx context.Context, req *DeclineBetRequest) error

	// CalculateCashOut calculates cash-out refund amounts for certain bets
	CalculateCashOut(ctx context.Context, req *CalculateCashOutRequest) (*CalculateCashOutResponse, error)

	// PlaceCashOutOrder notifies aboud cash-out on the bet
	PlaceCashOutOrder(ctx context.Context, req *PlaceCashOutOrderRequest) (*PlaceCashOutOrderResponse, error)

	// CancelCashOutOrder notifies that you've declined the cash-out that was previously accepted by the trading bet accounting system
	CancelCashOutOrder(ctx context.Context, req *CancelCashOutOrderRequest) error

	// GetRestrictions retrieves restrictions that could be violated when you tro to place a bet with the same parameters
	GetRestrictions(ctx context.Context, req *GetRestrictionsRequest) ([]Restriction, error)
}
