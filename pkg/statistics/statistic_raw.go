package statistics

import "encoding/json"

type RawStatistic struct {
	typ string
	json.RawMessage
}

func (r RawStatistic) Typ() string {
	return r.typ
}
