package util

import "testing"

type IncludesTestCases struct {
	name string
	s    []uint
	v    uint
	want bool
}

func Test_Includes(t *testing.T) {
	includesTestCases := []IncludesTestCases{
		{name: "element_found", s: []uint{1, 2, 3, 4}, v: 2, want: true},
		{name: "element_not_found", s: []uint{1, 2, 3, 4}, v: 5, want: false},
		{name: "empty_slice", s: []uint{}, v: 2, want: false},
	}

	for _, tc := range includesTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Includes(tc.s, tc.v)

			if got != tc.want {
				t.Errorf("Test case %s: expected %v, got %v", tc.name, tc.want, got)
			}
		})
	}
}
