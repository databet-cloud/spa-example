package statistics

type PenaltyStatus = string

const (
	UNKNOWN  PenaltyStatus = "unknown"
	AWARDED  PenaltyStatus = "awarded"
	MISSED   PenaltyStatus = "missed"
	ENDED    PenaltyStatus = "ended"
	CANCELED PenaltyStatus = "canceled"
)
