package main

import (
	"testing"
)

func TestRecurringCycleLength(t *testing.T) {
	in := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{0, 1, 0, 0, 1, 6, 0, 1, 0}
	for i, n := range in {
		out := RecurringCycleLength(n)
		if want[i] != out {
			t.Errorf("%v -> want:%v, out=%v", n, want[i], out)
		}
	}
}
