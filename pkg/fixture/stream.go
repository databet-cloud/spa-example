package fixture

import (
	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

type Stream struct {
	ID        string    `json:"id"`
	Locale    string    `json:"locale"`
	URL       string    `json:"url"`
	Platforms Platforms `json:"platforms"`
	Priority  int       `json:"priority"`
}

func (s Stream) WithPatch(tree patch.Tree) Stream {
	if v, ok := patch.GetFromTree[string](tree, "id"); ok {
		s.ID = v
	}

	if v, ok := patch.GetFromTree[string](tree, "locale"); ok {
		s.Locale = v
	}

	if v, ok := patch.GetFromTree[string](tree, "url"); ok {
		s.URL = v
	}

	if subTree := tree.SubTree("platforms"); !subTree.Empty() {
		s.Platforms = patch.PatchMap(s.Platforms, subTree)
	}

	if v, ok := patch.GetFromTree[int](tree, "priority"); ok {
		s.Priority = v
	}

	return s
}

type Streams map[string]Stream
