package statistics

type PenaltyStatus string

const (
	PenaltyStatusUnknown       PenaltyStatus = "unknown"
	PenaltyStatusAwarded       PenaltyStatus = "awarded"
	PenaltyStatusMissed        PenaltyStatus = "missed"
	PenaltyStatusEnded         PenaltyStatus = "ended"
	PenaltyStatusEndedCanceled PenaltyStatus = "canceled"
)
