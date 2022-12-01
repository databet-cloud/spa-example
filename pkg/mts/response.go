package mts

type errorResponse struct {
	Error *apiError `json:"error,omitempty"`
}

type placeBetResponse struct {
	Bet          *Bet          `json:"bet,omitempty"`
	Error        *apiError     `json:"error,omitempty"`
	Restrictions []Restriction `json:"restrictions,omitempty"`
}

type PlaceBetResponse struct {
	// Bet that was placed in case of successful response
	Bet *Bet `json:"bet,omitempty"`
	// Restrictions is the list of restrictions that aren't met and the bet is declined
	Restrictions []Restriction `json:"restrictions,omitempty"`
}

type calculateCashOutResponse struct {
	Amount       *CashOutAmount `json:"amount,omitempty"`
	Restrictions []Restriction  `json:"restrictions,omitempty"`
	Error        *apiError      `json:"error,omitempty"`
}

type CalculateCashOutResponse struct {
	Amount       *CashOutAmount `json:"amount,omitempty"`
	Restrictions []Restriction  `json:"restrictions,omitempty"`
}

type placeCashOutOrderResponse struct {
	Bet          *Bet          `json:"bet,omitempty"`
	CashOutOrder *CashOutOrder `json:"cash_out_order,omitempty"`
	Restrictions []Restriction `json:"restrictions,omitempty"`
	Error        *apiError     `json:"error,omitempty"`
}

type PlaceCashOutOrderResponse struct {
	Bet          *Bet          `json:"bet,omitempty"`
	CashOutOrder *CashOutOrder `json:"cash_out_order,omitempty"`
	Restrictions []Restriction `json:"restrictions,omitempty"`
}

type getRestrictionsResponse struct {
	Restrictions []Restriction `json:"restrictions,omitempty"`
	Error        *apiError     `json:"error,omitempty"`
}
