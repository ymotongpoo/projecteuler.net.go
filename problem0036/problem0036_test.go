package main

import (
	"testing"
)

func TestIsPalindromic(t *testing.T) {
	in10 := []uint64{
		uint64(11),
		uint64(12),
		uint64(131),
		uint64(132),
		uint64(1551),
		uint64(1552),
		uint64(1000001),
		uint64(1000002),
	}
	want10 := []bool{true, false, true, false, true, false, true, false}
	for i, n := range in10 {
		out := IsPalindromic(n, 10)
		if out != want10[i] {
			t.Errorf("%v -> want:%v, out=%v", n, want10[i], out)
		}
	}

	in2 := []uint64{
		uint64(2),
		uint64(3),
		uint64(6),
		uint64(7),
		uint64(8),
		uint64(9),
	}
	want2 := []bool{false, true, false, true, false, true}
	for i, n := range in2 {
		out := IsPalindromic(n, 2)
		if out != want2[i] {
			t.Errorf("%v -> want:%v, out=%v", n, want2[i], out)
		}
	}
}
