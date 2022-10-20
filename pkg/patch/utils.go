package patch

import (
	"golang.org/x/exp/maps"
)

type Patchable[T any] interface {
	WithPatch(tree Tree) T
}

func MapPatchable[P Patchable[P]](m map[string]P, tree Tree) map[string]P {
	res := maps.Clone(m)
	if res == nil {
		res = make(map[string]P, len(tree.Patch()))
	}

	for id, subTree := range tree.SubTrees() {
		res[id] = res[id].WithPatch(subTree)
	}

	return res
}

func PatchMap[V any](m map[string]V, tree Tree) map[string]V {
	res := maps.Clone(m)

	if res == nil {
		res = map[string]V{}
	}

	for k, v := range tree.Patch() {
		value, ok := v.(V)
		if ok {
			res[k] = value
		}
	}

	return res
}
