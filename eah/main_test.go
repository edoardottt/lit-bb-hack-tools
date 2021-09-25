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
		if !sliceContains(a, b[i]) {
			return false
		}
	}
	return true
}

//sliceContains (helper)
func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestRemoveDuplicateValuesTableDriven(t *testing.T) {
	var tests = []struct {
		a    []string
		want []string
	}{
		{[]string{}, []string{}},
		{[]string{"", ""}, []string{""}},
		{[]string{"a", "a", ""}, []string{"a", ""}},
		{[]string{"a", "a", "a", "b", "b"}, []string{"a", "b"}},
		{[]string{"a", "c"}, []string{"a", "c"}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"", "a", "c", "c", "c", "b", "f", "b", "f", "e", "c", "d", "a", "e", "f"},
			[]string{"", "a", "b", "c", "d", "e", "f"}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%s", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := RemoveDuplicateValues(tt.a)
			if !testEqStringSlice(ans, tt.want) {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
