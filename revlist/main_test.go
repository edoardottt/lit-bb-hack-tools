package main

import (
	"fmt"
	"testing"
)

//testEqStringSlice (helper)
func testEqStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestReverseTableDriven(t *testing.T) {
	var tests = []struct {
		a    []string
		want []string
	}{
		{[]string{}, []string{}},
		{[]string{"", ""}, []string{"", ""}},
		{[]string{"a", "a", ""}, []string{"", "a", "a"}},
		{[]string{"a", "a", "a", "b", "b"}, []string{"b", "b", "a", "a", "a"}},
		{[]string{"a", "c"}, []string{"c", "a"}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"", "a", "c", "b", "f", "e", "c", "d", "a", "e", "f"},
			[]string{"f", "e", "a", "d", "c", "e", "f", "b", "c", "a", ""}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%s", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := Reverse(tt.a)
			if !testEqStringSlice(ans, tt.want) {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
