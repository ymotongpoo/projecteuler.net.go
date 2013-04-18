package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	in := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := []uint64{
		uint64(1),
		uint64(1),
		uint64(2),
		uint64(6),
		uint64(24),
		uint64(120),
		uint64(720),
		uint64(5040),
		uint64(40320),
		uint64(362880),
	}
	for i, n := range in {
		out := Factorial(n)
		if want[i] != out {
			t.Errorf("want:%v, out=%v", want[i], out)
		}
	}
}

func TestSumOfFactorial(t *testing.T) {
	in := []int{1, 4, 5}
	want := uint64(145)
	out := SumOfFactorial(in)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func NumToDigits(t *testing.T) {
	in := []uint64{
		uint64(1),
		uint64(20),
		uint64(145),
	}
	want := [][]int{
		[]int{1},
		[]int{0, 2},
		[]int{5, 4, 1},
	}
	for i, n := range in {
		out := NumToDigits(n)
		if want[i] != out {
			t.Errorf("want:%v, out=%v", want[i], out)
		}
	}
}
