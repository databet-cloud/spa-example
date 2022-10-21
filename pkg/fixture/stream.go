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

	if streamPatch.ID != nil {
		s.ID = *streamPatch.ID
	}

	if streamPatch.Locale != nil {
		s.Locale = *streamPatch.Locale
	}

	if streamPatch.URL != nil {
		s.URL = *streamPatch.URL
	}

	if streamPatch.Priority != nil {
		s.Priority = *streamPatch.Priority
	}

	if subTree := tree.SubTree("platforms"); !subTree.Empty() {
		s.Platforms, err = patch.MapPatchable(s.Platforms, subTree)
		if err != nil {
			return Stream{}, fmt.Errorf("patch platforms: %w", err)
		}
	}

	return s, nil
}

type Streams map[string]Stream
