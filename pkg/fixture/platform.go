package fixture

import (
	"fmt"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

type Platform struct {
	Type             string   `json:"type"`
	AllowedCountries []string `json:"allowed_countries"`
	Enabled          bool     `json:"enabled"`
}

type PlatformPatch struct {
	Type             *string  `mapstructure:"type"`
	AllowedCountries []string `mapstructure:"allowed_countries"`
	Enabled          *bool    `mapstructure:"enabled"`
}

func (p Platform) WithPatch(tree patch.Tree) (Platform, error) {
	var platformPatch PlatformPatch

	err := tree.UnmarshalPatch(&platformPatch)
	if err != nil {
		return Platform{}, fmt.Errorf("decode platform patch: %w", err)
	}

	p.applyPlatformPatch(&platformPatch)

	return p, nil
}
func (p *Platform) ApplyPatch(tree patch.Tree) error {
	var platformPatch PlatformPatch

	err := tree.UnmarshalPatch(&platformPatch)
	if err != nil {
		return fmt.Errorf("decode platform patch: %w", err)
	}

	p.applyPlatformPatch(&platformPatch)

	return nil
}

func (p *Platform) applyPlatformPatch(patch *PlatformPatch) {
	if patch.Type != nil {
		p.Type = *patch.Type
	}

	if patch.AllowedCountries != nil {
		p.AllowedCountries = patch.AllowedCountries
	}

	if patch.Enabled != nil {
		p.Enabled = *patch.Enabled
	}
}

type Platforms map[string]Platform

func (p Platforms) Clone() Platforms {
	result := make(Platforms, len(p))
	for k, v := range p {
		result[k] = v
	}

	return result
}
