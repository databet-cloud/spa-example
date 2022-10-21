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

	if tournamentPatch.ID != nil {
		t.ID = *tournamentPatch.ID
	}

	if tournamentPatch.Name != nil {
		t.Name = *tournamentPatch.Name
	}

	if tournamentPatch.MasterID != nil {
		t.MasterID = *tournamentPatch.MasterID
	}

	if tournamentPatch.CountryCode != nil {
		t.CountryCode = *tournamentPatch.CountryCode
	}

	return t, nil
}
