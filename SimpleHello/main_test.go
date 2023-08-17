package main

import "testing"

func Test_add(t *testing.T) {
	d1 := 10
	d2 := 5

	r := add(d1, d2)

	if r != d1+d2 {
		t.Errorf("Expected outcome is %d, expected is %d", d1+d2, r)
	}
}
