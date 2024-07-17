package alib

import "testing"

func TestAverage(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	v := Average(s)
	if v != 3 {
		t.Errorf("%v != %v", v, 3)
	}
}
