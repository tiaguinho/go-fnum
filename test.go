package fnum

import "testing"

//testList
type testList []struct {
	exp, got string
}

//Validate
func (tl testList) validate(t *testing.T) {
	for _, test := range tl {
		if test.got != test.exp {
			t.Error("Expected: ", test.exp, " - Got: ", test.got)
		}
	}
}
