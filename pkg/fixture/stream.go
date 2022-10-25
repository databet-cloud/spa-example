//go:generate go run github.com/mailru/easyjson/easyjson stream.go
package fixture

import (
	"encoding/json"
	"fmt"
	"strings"
)

//easyjson:json
type Stream struct {
	ID        string    `json:"id"`
	Locale    string    `json:"locale"`
	URL       string    `json:"url"`
	Platforms Platforms `json:"platforms"`
	Priority  int       `json:"priority"`
}

func (s *Stream) ApplyPatch(path string, value json.RawMessage) error {
	var (
		unmarshaller     any
		key, rest, found = strings.Cut(path, "/")
		partialPatch     = found
	)

	switch key {
	case "id":
		unmarshaller = &s.ID
	case "locale":
		unmarshaller = &s.Locale
	case "url":
		unmarshaller = &s.URL
	case "priority":
		unmarshaller = &s.Priority
	case "platforms":
		if partialPatch {
			if s.Platforms == nil {
				return fmt.Errorf("patch nil platfroms")
			}

			return s.Platforms.ApplyPatch(rest, value)
		}

		unmarshaller = &s.Platforms
	default:
		return nil
	}

	err := json.Unmarshal(value, unmarshaller)
	if err != nil {
		return fmt.Errorf("unmarshal value: %w, path: %q", err, path)
	}

	return nil
}

type Streams map[string]Stream

func (s Streams) ApplyPatch(path string, value json.RawMessage) error {
	key, rest, found := strings.Cut(path, "/")
	stream, ok := s[key]

	if !found {
		err := json.Unmarshal(value, &stream)
		if err != nil {
			return fmt.Errorf("unmarshal stream: %w, path: %q", err, path)
		}

		s[key] = stream
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent stream: %q", path)
	}

	err := stream.ApplyPatch(rest, value)
	if err != nil {
		return fmt.Errorf("patch stream: %w, path: %q", err, rest)
	}

	s[key] = stream
	return nil
}
