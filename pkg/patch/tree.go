package patch

import "strings"

type Tree struct {
	patch          Patch
	level          string
	levelDelimiter string
}

func NewTree(patch Patch, levelDelimiter string) Tree {
	return Tree{
		patch:          patch,
		level:          "",
		levelDelimiter: levelDelimiter,
	}
}

func (t Tree) SubTree(level string) Tree {
	if subPatch, ok := t.Get(level).(Patch); ok {
		return t.newSubTree(level, subPatch)
	}

	subPatch := make(Patch)
	levelPrefix := level + t.levelDelimiter

	for field, value := range t.patch {
		if strings.HasPrefix(field, levelPrefix) {
			subPatch[strings.TrimPrefix(field, levelPrefix)] = value
		}
	}

	return t.newSubTree(level, subPatch)
}

func (t Tree) SubTrees() map[string]Tree {
	subTrees := make(map[string]Tree)

	for field := range t.patch {
		level := field

		if strings.Contains(field, t.levelDelimiter) {
			level = strings.SplitN(field, t.levelDelimiter, 2)[0]
		}

		subTrees[level] = t.SubTree(level)
	}

	return subTrees
}

func (t Tree) newSubTree(level string, patch Patch) Tree {
	t.patch = patch
	t.level = t.level + t.levelDelimiter + level

	return t
}

func (t Tree) Has(field string) bool {
	_, ok := t.patch[field]
	return ok
}

func (t Tree) Get(field string) any {
	return t.patch[field]
}

func (t Tree) LastLevel() string {
	values := strings.Split(t.level, t.levelDelimiter)
	if len(values) == 0 {
		return ""
	}

	return values[len(values)-1]
}

func (t Tree) Patch() Patch {
	return t.patch
}

func (t Tree) Empty() bool {
	return len(t.patch) == 0
}

func GetFromTree[T any](patchTree Tree, field string) (value T, exists bool) {
	if !patchTree.Has(field) {
		return value, false
	}

	v, ok := patchTree.Get(field).(T)
	return v, ok
}
