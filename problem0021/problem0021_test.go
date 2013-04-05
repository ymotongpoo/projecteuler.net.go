package main

import (
	"testing"
)

func TestSumOfDivisors(t *testing.T) {
	in := []int{220, 284}
	want := []int{284, 220}
	for i, n := range in {
		out := SumOfDivisors(n)
		if want[i] != out {
			t.Errorf("%v: want:%v, out=%v", n, want[i], out)
		}
	}
}
