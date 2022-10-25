//go:generate go run github.com/mailru/easyjson/easyjson venue.go
package fixture

import (
	"encoding/json"
)

//easyjson:json
type Venue struct {
	ID string `json:"id"`
}

func (v *Venue) ApplyPatch(path string, value json.RawMessage) error {
	if path == "id" {
		return json.Unmarshal(value, &v.ID)
	}

	return nil
}
