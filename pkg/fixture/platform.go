//go:generate go run github.com/mailru/easyjson/easyjson platform.go
package fixture

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
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

func (p *Platform) UnmarshalSimdJSON(obj *simdjson.Object, reuseIter *simdjson.Iter) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
	}

	for {
		name, elementType, err := obj.NextElementBytes(reuseIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		switch string(name) {
		case "type":
			p.Type, err = simdutil.UnsafeStrFromIter(reuseIter)
		case "allowed_countries":
			array, err := reuseIter.Array(nil)
			if err != nil {
				return err
			}

			strArray, err := array.AsString()
			if err != nil {
				return err
			}

			p.AllowedCountries = strArray

		case "enabled":
			p.Enabled, err = reuseIter.Bool()
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", string(name), err)
		}
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

func (p Platforms) UnmarshalSimdJSON(obj *simdjson.Object, reuseIter *simdjson.Iter) error {
	if reuseIter == nil {
		reuseIter = new(simdjson.Iter)
	}

	platformObj := new(simdjson.Object)
	tmpPlatform := new(Platform)

	for {
		name, elementType, err := obj.NextElement(reuseIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		platformObj, err = reuseIter.Object(platformObj)
		if err != nil {
			return fmt.Errorf("create %q object: %w", name, err)
		}

		err = tmpPlatform.UnmarshalSimdJSON(platformObj, reuseIter)
		if err != nil {
			return fmt.Errorf("unmarshal %q odd: %w", name, err)
		}

		p[name] = *tmpPlatform
	}

	return nil
}
