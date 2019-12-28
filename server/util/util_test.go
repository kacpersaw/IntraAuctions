package util

import "testing"

func TestStringInSlice(t *testing.T) {
	type Args struct {
		Str string
		Slice []string
	}
	tests := []struct {
		Name string
		Args Args
		Want bool
	}{
		{
			Name: "success",
			Args: Args{
				Str: "b",
				Slice: []string{"a", "b", "c"},
			},
			Want: true,
		},
		{
			Name: "failure",
			Args: Args{
				Str: "d",
				Slice: []string{"a", "b", "c"},
			},
			Want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := StringInSlice(tt.Args.Str, tt.Args.Slice)
			if got != tt.Want {
				t.Errorf("StringInSlice() got = %v, want  = %v", got, tt.Want)
			}
		})
	}
}
