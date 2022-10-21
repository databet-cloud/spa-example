package fixture

import (
	"fmt"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

type Score struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Points string `json:"points"`
	Number int    `json:"number"`
}

type ScorePatch struct {
	ID     *string `mapstructure:"id"`
	Type   *string `mapstructure:"type"`
	Points *string `mapstructure:"points"`
	Number *int    `mapstructure:"number"`
}

func (s Score) WithPatch(tree patch.Tree) (Score, error) {
	var scorePatch ScorePatch

	err := tree.UnmarshalPatch(&scorePatch)
	if err != nil {
		return Score{}, fmt.Errorf("unmarshal score patch: %w", err)
	}

	if scorePatch.ID != nil {
		s.ID = *scorePatch.ID
	}

	if scorePatch.Type != nil {
		s.Type = *scorePatch.Type
	}

	if scorePatch.Points != nil {
		s.Points = *scorePatch.Points
	}

	if scorePatch.Number != nil {
		s.Number = *scorePatch.Number
	}

	return s, nil
}

type Scores map[string]Score

func (s Scores) Clone() Scores {
	result := make(Scores, len(s))
	for k, v := range s {
		result[k] = v
	}

	return result
}
