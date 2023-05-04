package statistics

type Timer struct {
	IsActive  bool  `json:"is_active"`
	StartedAt int64 `json:"started_at"`
	EndedAt   int64 `json:"ended_at"`
	TimeDelta int64 `json:"time_delta"`
}
