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

	s.applyScorePatch(&scorePatch)

	return s, nil
}

func (s *Score) ApplyPatch(tree patch.Tree) error {
	var scorePatch ScorePatch

	err := tree.UnmarshalPatch(&scorePatch)
	if err != nil {
		return fmt.Errorf("unmarshal score patch: %w", err)
	}

	s.applyScorePatch(&scorePatch)

	return nil
}

func (s *Score) applyScorePatch(patch *ScorePatch) {
	if patch.ID != nil {
		s.ID = *patch.ID
	}

	if patch.Type != nil {
		s.Type = *patch.Type
	}

	if patch.Points != nil {
		s.Points = *patch.Points
	}

	if patch.Number != nil {
		s.Number = *patch.Number
	}
}

type Scores map[string]Score

func (s Scores) Clone() Scores {
	result := make(Scores, len(s))
	for k, v := range s {
		result[k] = v
	}

	return result
}
