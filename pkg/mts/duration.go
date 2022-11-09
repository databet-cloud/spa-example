package mts

import "time"

type DurationMS int64

func (ms DurationMS) ToTimeDuration() time.Duration {
	return time.Duration(ms) * time.Millisecond
}
