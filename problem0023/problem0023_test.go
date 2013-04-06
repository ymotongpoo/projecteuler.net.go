package main

import (
	"testing"
)

func TestIsAbundant(t *testing.T) {
	in := []int64{12, 28}
	want := []bool{true, false}
	for i, n := range in {
		out := IsAbundant(n)
		if want[i] != out {
			t.Errorf("%v -> want:%v, out=%v", n, want[i], out)
		}
	}
}
