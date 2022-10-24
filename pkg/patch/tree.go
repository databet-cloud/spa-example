package patch

import (
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Tree interface {
	SubTree(level string) Tree
	SubTrees() map[string]Tree
	Has(field string) bool
	Get(field string) any
	LastLevel() string
	Patch() Patch
	Empty() bool
	UnmarshalPatch(destination any) error
}

type MapTree struct {
	patch          Patch
	level          string
	levelDelimiter string
}

func NewMapTree(patch Patch, levelDelimiter string) *MapTree {
	return &MapTree{
		patch:          patch,
		level:          "",
		levelDelimiter: levelDelimiter,
	}
}

func (t *MapTree) SubTree(level string) Tree {
	if subPatch := castToPatch(t.Get(level)); subPatch != nil {
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

func castToPatch(value any) Patch {
	switch v := value.(type) {
	case Patch:
		return v
	case map[string]any:
		return v
	default:
		return nil
	}
}

func (t *MapTree) SubTrees() map[string]Tree {
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

func (t *MapTree) newSubTree(level string, patch Patch) *MapTree {
	return &MapTree{
		patch:          patch,
		level:          t.level + t.levelDelimiter + level,
		levelDelimiter: t.levelDelimiter,
	}
}

func (t *MapTree) Has(field string) bool {
	_, ok := t.patch[field]
	return ok
}

func (t *MapTree) Get(field string) any {
	return t.patch[field]
}

func (t *MapTree) LastLevel() string {
	values := strings.Split(t.level, t.levelDelimiter)
	if len(values) == 0 {
		return ""
	}

	return values[len(values)-1]
}

func (t *MapTree) Patch() Patch {
	return t.patch
}

func (t *MapTree) UnmarshalPatch(destination any) error {
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.OrComposeDecodeHookFunc(
				mapstructure.StringToTimeHookFunc(time.RFC3339),
				mapstructure.StringToTimeHookFunc(time.RFC3339Nano),
			),
		),
		Result: destination,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(t.patch)
}

func (t *MapTree) Empty() bool {
	return len(t.patch) == 0
}

func GetFromTree[T any](patchTree Tree, field string) (value T, exists bool) {
	if !patchTree.Has(field) {
		return value, false
	}

	v, ok := patchTree.Get(field).(T)
	return v, ok
}
