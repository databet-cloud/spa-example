package patch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/pkg/patch"
)

func TestTreeSubTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    patch.Patch
		level    string
		expected patch.Patch
	}{
		{
			name: "full object subtree",
			input: patch.Patch{
				"field1": 1,
				"field2": 2,
				"field3": 3,
				"field4": map[string]any{
					"subfield1": 1,
					"subfield2": 2,
					"subfield3": 3,
				},
			},
			level: "field4",
			expected: patch.Patch{
				"subfield1": 1,
				"subfield2": 2,
				"subfield3": 3,
			},
		},
		{
			name: "subfields with delimiters",
			input: patch.Patch{
				"field1":           1,
				"field2":           2,
				"field3":           3,
				"field4/subfield1": 1,
				"field4/subfield2": 2,
				"field4/subfield3": patch.Patch{
					"subsubfield": 123,
				},
			},
			level: "field4",
			expected: patch.Patch{
				"subfield1": 1,
				"subfield2": 2,
				"subfield3": patch.Patch{
					"subsubfield": 123,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tree := patch.NewMapTree(tc.input, "/")
			subTree := tree.SubTree(tc.level)
			assert.Equal(t, tc.expected, subTree.Patch())
		})
	}
}
