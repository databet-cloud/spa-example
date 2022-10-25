//go:generate go run github.com/mailru/easyjson/easyjson score.go
package fixture

import (
	"encoding/json"
	"fmt"
	"strings"
)

//easyjson:json
type Score struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Points string `json:"points"`
	Number int    `json:"number"`
}

func (s *Score) ApplyPatch(key string, value json.RawMessage) error {
	var unmarshaller any

	switch key {
	case "id":
		unmarshaller = &s.ID
	case "type":
		unmarshaller = &s.Type
	case "points":
		unmarshaller = &s.Points
	case "number":
		unmarshaller = &s.Number
	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("unmarshal score value: %w, field: %q", err, key)
	}

	return nil
}

//easyjson:json
type Scores map[string]Score

func (s Scores) Clone() Scores {
	result := make(Scores, len(s))
	for k, v := range s {
		result[k] = v
	}

	return result
}

func (s Scores) ApplyPatch(path string, value json.RawMessage) error {
	key, rest, found := strings.Cut(path, "/")
	score, ok := s[key]

	if !found {
		err := json.Unmarshal(value, &score)
		if err != nil {
			return fmt.Errorf("unmarshal competitor: %w, key: %q", err, key)
		}

		s[key] = score
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent score: %q", key)
	}

	err := score.ApplyPatch(rest, value)
	if err != nil {
		return fmt.Errorf("patch competitor: %w, path: %q", err, rest)
	}

	s[key] = score
	return nil
}
