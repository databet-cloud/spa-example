//go:generate go run github.com/mailru/easyjson/easyjson platform.go
package fixture

import (
	"encoding/json"
	"fmt"
	"strings"
)

//easyjson:json
type Platform struct {
	Type             string   `json:"type"`
	AllowedCountries []string `json:"allowed_countries"`
	Enabled          bool     `json:"enabled"`
}

func (p *Platform) ApplyPatch(path string, value json.RawMessage) error {
	var unmarshaller any

	switch path {
	case "type":
		unmarshaller = &p.Type
	case "allowed_countries":
		unmarshaller = &p.AllowedCountries
	case "enabled":
		unmarshaller = &p.Enabled
	default:
		return nil
	}

	err := json.Unmarshal(value, &unmarshaller)
	if err != nil {
		return fmt.Errorf("unmarshal field: %w, field: %q", err, path)
	}

	return nil
}

//easyjson:json
type Platforms map[string]Platform

func (p Platforms) Clone() Platforms {
	result := make(Platforms, len(p))
	for k, v := range p {
		result[k] = v
	}

	return result
}

func (p Platforms) ApplyPatch(path string, value json.RawMessage) error {
	key, rest, found := strings.Cut(path, "/")
	platform, ok := p[key]

	if !found {
		err := json.Unmarshal(value, &platform)
		if err != nil {
			return fmt.Errorf("unmarshal platform: %w, path: %q", err, path)
		}

		p[key] = platform
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent platform: %q", path)
	}

	err := platform.ApplyPatch(rest, value)
	if err != nil {
		return fmt.Errorf("patch platform: %w, path: %q", err, rest)
	}

	p[key] = platform
	return nil
}
