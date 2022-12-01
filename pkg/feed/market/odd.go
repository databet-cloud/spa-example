package market

const (
	OddStatusNotResulted OddStatus = iota
	OddStatusWin
	OddStatusLoss
	OddStatusHalfWin
	OddStatusHalfLoss
	OddStatusRefunded
	OddStatusCancelled
)

type OddStatus int

// nolint:gochecknoglobals // package dictionary of odd statuses
var oddStatuses = []OddStatus{
	OddStatusNotResulted,
	OddStatusWin,
	OddStatusLoss,
	OddStatusHalfWin,
	OddStatusHalfLoss,
	OddStatusRefunded,
	OddStatusCancelled,
}

func (os OddStatus) IsValid() bool {
	for _, status := range oddStatuses {
		if os == status {
			return true
		}
	}

	return false
}

func (os OddStatus) IsResulted() bool {
	return os != OddStatusNotResulted
}

// Odds is a map with odd id in its key and odd as a value
type Odds map[string]Odd

func (c Odds) Equals(other Odds) bool {
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

// Clone odds and return it's deep copy
func (c Odds) Clone() Odds {
	newC := make(Odds, len(c))
	for id, odd := range c {
		newC[id] = odd.Clone()
	}

	return newC
}

type Odd struct {
	// ID of the current odd
	ID string `json:"id"`
	// Template is a name with specifiers, that could be replaced to format market's name
	Template string `json:"template"`
	// IsActive indicates whether this odd is active and you could accept bets on it or not
	IsActive bool `json:"is_active"`
	// Status indicates odd status
	Status OddStatus `json:"status"`
	// Value is a float coefficient in string representation
	Value string `json:"value"`
	// Reason indicates why the current status is set to the current odd
	StatusReason string `json:"status_reason"`
}

func (o Odd) Equals(other Odd) bool {
	return o.ID == other.ID &&
		o.Template == other.Template &&
		o.IsActive == other.IsActive &&
		o.Status == other.Status &&
		o.Value == other.Value &&
		o.StatusReason == other.StatusReason
}

// Clone odd and return it's deep copy
func (o Odd) Clone() Odd {
	return Odd{
		ID:           o.ID,
		Template:     o.Template,
		IsActive:     o.IsActive,
		Status:       o.Status,
		Value:        o.Value,
		StatusReason: o.StatusReason,
	}
}
