package market

type StatusMarket struct {
	ID     string              `json:"id"`
	TypeID int                 `json:"type_id"`
	Status Status              `json:"status"`
	Odds   StatusOddCollection `json:"odds"`
}

func (sm StatusMarket) Equals(other StatusMarket) bool {
	return sm.ID == other.ID &&
		sm.TypeID == other.TypeID &&
		sm.Status == other.Status &&
		sm.Odds.Equals(other.Odds)
}

type StatusMarketCollection map[string]StatusMarket

func (c StatusMarketCollection) Equals(other StatusMarketCollection) bool {
	if len(c) != len(other) {
		return false
	}

	for id, m := range c {
		otherM, ok := other[id]
		if !ok {
			return false
		}

		if !m.Equals(otherM) {
			return false
		}
	}

	return true
}

type StatusOdd struct {
	ID           string    `json:"id"`
	Status       OddStatus `json:"status"`
	StatusReason string    `json:"status_reason"`
}

func (so StatusOdd) Equals(other StatusOdd) bool {
	return so.ID == other.ID &&
		so.Status == other.Status &&
		so.StatusReason == other.StatusReason
}

type StatusOddCollection map[string]StatusOdd

func (c StatusOddCollection) Equals(other StatusOddCollection) bool {
	if len(c) != len(other) {
		return false
	}

	for id, odd := range c {
		otherOdd, ok := other[id]
		if !ok {
			return false
		}

		if !odd.Equals(otherOdd) {
			return false
		}
	}

	return true
}
