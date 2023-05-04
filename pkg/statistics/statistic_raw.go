package statistics

import "encoding/json"

type RawStatistic struct {
	typ Type
	json.RawMessage
}

func (r RawStatistic) GetType() Type {
	return r.typ
}
