package fixture

import "github.com/databet-cloud/databet-go-sdk/pkg/patch"

type Venue struct {
	ID string `json:"id"`
}

func (v Venue) WithPatch(tree patch.Tree) Venue {
	id, ok := patch.GetFromTree[string](tree, "id")
	if ok {
		v.ID = id
	}

	return v
}
