package fixture

import "github.com/databet-cloud/databet-go-sdk/pkg/patch"

type Tournament struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	MasterID    string `json:"master_id"`
	CountryCode string `json:"country_code"`
}

func (t Tournament) WithPatch(patchTree patch.Tree) Tournament {
	if v, ok := patch.GetFromTree[string](patchTree, "id"); ok {
		t.ID = v
	}

	if v, ok := patch.GetFromTree[string](patchTree, "name"); ok {
		t.Name = v
	}

	if v, ok := patch.GetFromTree[string](patchTree, "master_id"); ok {
		t.MasterID = v
	}

	if v, ok := patch.GetFromTree[string](patchTree, "country_code"); ok {
		t.CountryCode = v
	}

	return t
}
