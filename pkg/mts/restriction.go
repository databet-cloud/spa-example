package mts

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	RestrictionMaxBet                          Type = "max_bet"
	RestrictionMaxOddBet                       Type = "max_odd_bet"
	RestrictionBetType                         Type = "bet_type"
	RestrictionBetStatus                       Type = "bet_status"
	RestrictionBetInterval                     Type = "bet_interval"
	RestrictionBetDelay                        Type = "bet_delay"
	RestrictionBetSelectionExistence           Type = "bet_selection_existence"
	RestrictionSelectionValue                  Type = "selection_value"
	RestrictionRestrictionSelectionValueChange Type = "selection_value_change"
	RestrictionRestrictionSportEventType       Type = "sport_event_type"
	RestrictionRestrictionSportEventStatus     Type = "sport_event_status"
	RestrictionRestrictionSportEventExistence  Type = "sport_event_existence"
	RestrictionRestrictionSportEventBetStop    Type = "sport_event_bet_stop"
	RestrictionMarketStatus                    Type = "market_status"
	RestrictionMarketExistence                 Type = "market_existence"
	RestrictionMarketDefective                 Type = "market_defective"
	RestrictionOddStatus                       Type = "odd_status"
	RestrictionOddExistence                    Type = "odd_existence"
	RestrictionCashOutBetType                  Type = "cash_out_bet_type"
	RestrictionBetCashOutSelection             Type = "bet_cash_out_selections_mismatch"
	RestrictionCashOutUnavailable              Type = "cash_out_unavailable"
	RestrictionCashOutAmountLimit              Type = "cash_out_amount_limit"
	RestrictionCashOutRefundAmount             Type = "cash_out_refund_amount"
	RestrictionCashOutOrderStatus              Type = "cash_out_order_status"
)

const (
	CtxKeyCashOutRefundAmount           = "cash_out_refund_amount"
	CtxKeyCashOutCalculatedRefundAmount = "cash_out_calculated_refund_amount"
	CtxKeyCashOutAmount                 = "cash_out_amount"
	CtxKeyCashOutMinAmount              = "cash_out_min_amount"
	CtxKeyCashOutMaxAmount              = "cash_out_max_amount"
	CtxKeyMaxBet                        = "max_bet"
	CtxKeyMaxOddBet                     = "max_odd_bet"
	CtxKeyReason                        = "reason"
	CtxKeySportEventID                  = "sport_event_id"
	CtxKeyMarketID                      = "market_id"
	CtxKeyOddID                         = "odd_id"
	CtxKeyBetType                       = "bet_type"
	CtxKeyBetStatus                     = "bet_status"
	CtxKeyValue                         = "value"
	CtxKeyMarge                         = "marge"
	CtxKeyType                          = "type"
	CtxKeyDelay                         = "delay"
	CtxKeyTimeToWait                    = "time_to_wait"
	CtxKeyStatus                        = "status"
	CtxKeyIsActive                      = "is_active"
	CtxKeyBetSelections                 = "bet_selections"
	CtxKeyCashOutSelections             = "cash_out_selections"
	CtxKeyCashOutBetType                = "cash_out_bet_type"
	CtxKeyCashOutOrderID                = "cash_out_order_id"
	CtxKeyCashOutOrderStatus            = "cash_out_order_status"
)

type Restriction struct {
	Type    Type           `json:"type"`
	Context map[string]any `json:"context"`
}

func NewRestriction(restrictionType Type, context map[string]any) Restriction {
	return Restriction{Type: restrictionType, Context: context}
}
