package main

import (
	"testing"
)

func TestSumOfDiag(t *testing.T) {
	in := []int{2, 5}
	want := []uint64{uint64(0), uint64(101)}
	for i, n := range in {
		out := SumOfDiag(n)
		if want[i] != out {
			t.Errorf("%v -> want:%v, out=%v", n, want[i], out)
		}
	}
}
