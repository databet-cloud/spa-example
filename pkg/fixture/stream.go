package fixture

import (
	"fmt"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

type Stream struct {
	ID        string    `json:"id"`
	Locale    string    `json:"locale"`
	URL       string    `json:"url"`
	Platforms Platforms `json:"platforms"`
	Priority  int       `json:"priority"`
}

type StreamPatch struct {
	ID       *string `mapstructure:"id"`
	Locale   *string `mapstructure:"locale"`
	URL      *string `mapstructure:"url"`
	Priority *int    `mapstructure:"priority"`
}

func (s Stream) WithPatch(tree patch.Tree) (Stream, error) {
	var streamPatch StreamPatch

	err := tree.UnmarshalPatch(&streamPatch)
	if err != nil {
		return Stream{}, fmt.Errorf("unmarshal stream patch: %w", err)
	}

	s.applyStreamPatch(&streamPatch)

	if subTree := tree.SubTree("platforms"); !subTree.Empty() {
		s.Platforms, err = patch.MapPatchable(s.Platforms, subTree)
		if err != nil {
			return Stream{}, fmt.Errorf("patch platforms: %w", err)
		}
	}

	return s, nil
}

func (s *Stream) ApplyPatch(tree patch.Tree) error {
	var streamPatch StreamPatch

	err := tree.UnmarshalPatch(&streamPatch)
	if err != nil {
		return fmt.Errorf("unmarshal stream patch: %w", err)
	}

	s.applyStreamPatch(&streamPatch)

	if subTree := tree.SubTree("platforms"); !subTree.Empty() {
		s.Platforms, err = patch.MapPatchable(s.Platforms, subTree)
		if err != nil {
			return fmt.Errorf("patch platforms: %w", err)
		}
	}

	return nil
}

func (s *Stream) applyStreamPatch(patch *StreamPatch) {
	if patch.ID != nil {
		s.ID = *patch.ID
	}

	if patch.Locale != nil {
		s.Locale = *patch.Locale
	}

	if patch.URL != nil {
		s.URL = *patch.URL
	}

	if patch.Priority != nil {
		s.Priority = *patch.Priority
	}
}

type Streams map[string]Stream
