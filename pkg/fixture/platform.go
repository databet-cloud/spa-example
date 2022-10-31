package fixture

import (
	"fmt"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

type Platform struct {
	Type             string   `json:"type"`
	AllowedCountries []string `json:"allowed_countries"`
	Enabled          bool     `json:"enabled"`
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

type Platforms map[string]Platform

func (p Platforms) Clone() Platforms {
	result := make(Platforms, len(p))
	for k, v := range p {
		result[k] = v
	}

	return result
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
