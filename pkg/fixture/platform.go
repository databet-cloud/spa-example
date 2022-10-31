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

func (p *Platform) UnmarshalSimdJSON(obj *simdjson.Object) error {
	iter := new(simdjson.Iter)

	for {
		name, elementType, err := obj.NextElement(iter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		err = p.unmarshalFieldSimdJSON(name, iter)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Platform) ApplyPatchSimdJSON(key string, iter *simdjson.Iter) error {
	return p.unmarshalFieldSimdJSON(key, iter)
}

func (p *Platform) unmarshalFieldSimdJSON(key string, iter *simdjson.Iter) error {
	var err error

	switch key {
	case "type":
		p.Type, err = simdutil.UnsafeStrFromIter(iter)
	case "allowed_countries":
		array, err := iter.Array(nil)
		if err != nil {
			return err
		}

		strArray, err := array.AsString()
		if err != nil {
			return err
		}

		p.AllowedCountries = strArray

	case "enabled":
		p.Enabled, err = iter.Bool()
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
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

func (p Platforms) UnmarshalSimdJSON(obj *simdjson.Object) error {
	tmpIter := new(simdjson.Iter)
	oddObj := new(simdjson.Object)
	tmpPlatform := new(Platform)

	for {
		name, elementType, err := obj.NextElement(tmpIter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		oddObj, err = tmpIter.Object(oddObj)
		if err != nil {
			return fmt.Errorf("create %q object: %w", name, err)
		}

		err = tmpPlatform.UnmarshalSimdJSON(oddObj)
		if err != nil {
			return fmt.Errorf("unmarshal %q odd: %w", name, err)
		}

		p[name] = *tmpPlatform
	}

	return nil
}

func (p Platforms) FromIter(iter *simdjson.Iter) error {
	obj, err := iter.Object(nil)
	if err != nil {
		return err
	}

	return p.UnmarshalSimdJSON(obj)
}

func (s Platforms) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	platform, ok := s[key]

	if !partialPatch {
		obj, err := iter.Object(nil)
		if err != nil {
			return err
		}

		err = platform.UnmarshalSimdJSON(obj)
		if err != nil {
			return fmt.Errorf("platform %q unmarshal simdjson: %w", key, err)
		}

		s[key] = platform
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent platform: %q", key)
	}

	err := platform.ApplyPatchSimdJSON(rest, iter)
	if err != nil {
		return fmt.Errorf("apply platform patch: %w", err)
	}

	s[key] = platform
	return nil
}
