package fixture

import (
	"github.com/databet-cloud/databet-go-sdk/internal/generic"
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

type Platform struct {
	Type             string   `json:"type"`
	AllowedCountries []string `json:"allowed_countries"`
	Enabled          bool     `json:"enabled"`
}

func (p Platform) WithPatch(tree patch.Tree) Platform {
	if v, ok := patch.GetFromTree[string](tree, "type"); ok {
		p.Type = v
	}

	if tree.Has("allowed_countries") {
		allowedCountries, ok := patch.GetFromTree[[]string](tree, "allowed_countries")
		if !ok {
			if v, ok := patch.GetFromTree[[]interface{}](tree, "allowed_countries"); ok {
				allowedCountries = generic.CastSlice[string](v)
			}
		}

		p.AllowedCountries = allowedCountries
	}

	if v, ok := patch.GetFromTree[bool](tree, "enabled"); ok {
		p.Enabled = v
	}

	return p
}

type Platforms map[string]Platform

func (p Platforms) Clone() Platforms {
	result := make(Platforms, len(p))
	for k, v := range p {
		result[k] = v
	}

	return result
}
