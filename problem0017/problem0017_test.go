package main

import (
	"testing"
)

func TestLengthInWords(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 342, 115}
	want := []int{3, 3, 5, 4, 4, 23, 20}
	for i, n := range input {
		out := LengthInWords(n)
		if want[i] != out {
			t.Errorf("%v: want:%v, out=%v", n, want[i], out)
		}
	}
}
