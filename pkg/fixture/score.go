package fixture

import (
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

type Score struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Points string `json:"points"`
	Number int    `json:"number"`
}

func (s Score) WithPatch(patchTree patch.Tree) Score {
	if v, ok := patch.GetFromTree[string](patchTree, "id"); ok {
		s.ID = v
	}

	if v, ok := patch.GetFromTree[string](patchTree, "type"); ok {
		s.Type = v
	}

	if v, ok := patch.GetFromTree[string](patchTree, "points"); ok {
		s.Points = v
	}

	if v, ok := patch.GetFromTree[int](patchTree, "number"); ok {
		s.Number = v
	} else if v, ok := patch.GetFromTree[float64](patchTree, "number"); ok {
		s.Number = int(v)
	}

	return s
}

type Scores map[string]Score

func (s Scores) Clone() Scores {
	result := make(Scores, len(s))
	for k, v := range s {
		result[k] = v
	}

	return result
}
