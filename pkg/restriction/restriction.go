package restriction

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	MaxBet                Type = "max_bet"
	MaxOddBet             Type = "max_odd_bet"
	BetType               Type = "bet_type"
	BetStatus             Type = "bet_status"
	BetInterval           Type = "bet_interval"
	BetDelay              Type = "bet_delay"
	BetSelectionExistence Type = "bet_selection_existence"
	SelectionValue        Type = "selection_value"
	SelectionValueChange  Type = "selection_value_change"
	SportEventType        Type = "sport_event_type"
	SportEventStatus      Type = "sport_event_status"
	SportEventExistence   Type = "sport_event_existence"
	SportEventBetStop     Type = "sport_event_bet_stop"
	MarketStatus          Type = "market_status"
	MarketExistence       Type = "market_existence"
	MarketDefective       Type = "market_defective"
	OddStatus             Type = "odd_status"
	OddExistence          Type = "odd_existence"
	CashOutBetType        Type = "cash_out_bet_type"
	BetCashOutSelection   Type = "bet_cash_out_selections_mismatch"
	CashOutUnavailable    Type = "cash_out_unavailable"
	CashOutAmountLimit    Type = "cash_out_amount_limit"
	CashOutRefundAmount   Type = "cash_out_refund_amount"
	CashOutOrderStatus    Type = "cash_out_order_status"
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
	Type    Type                   `json:"type"`
	Context map[string]interface{} `json:"context"`
}

func NewRestriction(restrictionType Type, context map[string]interface{}) Restriction {
	return Restriction{Type: restrictionType, Context: context}
}
