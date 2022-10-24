package fixture

import (
	"fmt"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

type Tournament struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	MasterID    string `json:"master_id"`
	CountryCode string `json:"country_code"`
}

type TournamentPatch struct {
	ID          *string `mapstructure:"id"`
	Name        *string `mapstructure:"name"`
	MasterID    *string `mapstructure:"master_id"`
	CountryCode *string `mapstructure:"country_code"`
}

func (t Tournament) WithPatch(tree patch.Tree) (Tournament, error) {
	var tournamentPatch TournamentPatch

	err := tree.UnmarshalPatch(&tournamentPatch)
	if err != nil {
		return Tournament{}, fmt.Errorf("unmarshal tournament patch: %w", err)
	}

	t.applyTournamentPatch(&tournamentPatch)

	return t, nil
}

func (t *Tournament) ApplyPatch(tree patch.Tree) error {
	var tournamentPatch TournamentPatch

	err := tree.UnmarshalPatch(&tournamentPatch)
	if err != nil {
		return fmt.Errorf("unmarshal tournament patch: %w", err)
	}

	t.applyTournamentPatch(&tournamentPatch)

	return nil
}

func (t *Tournament) applyTournamentPatch(patch *TournamentPatch) {
	if patch.ID != nil {
		t.ID = *patch.ID
	}

	if patch.Name != nil {
		t.Name = *patch.Name
	}

	if patch.MasterID != nil {
		t.MasterID = *patch.MasterID
	}

	if patch.CountryCode != nil {
		t.CountryCode = *patch.CountryCode
	}

}
