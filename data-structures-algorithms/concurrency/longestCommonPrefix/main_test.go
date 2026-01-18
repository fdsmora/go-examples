package longestcommonprefix

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			"same length",
			[]string{"flower", "flowir", "flowes"},
			"flow",
		},
		{
			"variable length",
			[]string{"flower", "flow", "flxwesass"},
			"fl",
		},
		{
			"no common prefix, variable length",
			[]string{"asba", "ewe", "c"},
			"",
		},
		{
			"one word",
			[]string{"asba"},
			"asba",
		},
		{
			"many words",
			[]string{"asba", "asbafwef", "asbasdf", "asbaqweqwe", "asba3ijfwaiejf;as", "asbaasgasdgasdgiiasdgasg"},
			"asba",
		},
		{
			"many words",
			[]string{"xasba", "asbafwef", "asbasdf", "asbaqweqwe", "asba3ijfwaiejf;as", "asbaasgasdgasdgiiasdgasg"},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.input)
			if tt.want != got {
				t.Errorf("Got: '%s', want: '%s'", got, tt.want)
			}
		})
	}
}
