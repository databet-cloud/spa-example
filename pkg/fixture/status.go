package fixture

type Status int

const (
	StatusNotStarted = iota
	StatusLive
	StatusSuspended
	StatusEnded
	StatusClosed
	StatusCancelled
	StatusAbandoned
	StatusDelayed
	StatusUnknown
)
