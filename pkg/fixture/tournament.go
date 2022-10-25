package fixture

import (
	"encoding/json"
	"fmt"
)

type Tournament struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	MasterID    string `json:"master_id"`
	CountryCode string `json:"country_code"`
}

func (t *Tournament) ApplyPatch(key string, value json.RawMessage) error {
	var unmarshaller any

	switch key {
	case "id":
		unmarshaller = &t.ID
	case "name":
		unmarshaller = &t.Name
	case "master_id":
		unmarshaller = &t.MasterID
	case "country_code":
		unmarshaller = &t.CountryCode
	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}
