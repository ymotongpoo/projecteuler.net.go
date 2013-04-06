package main

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	in := []int{3, 5}
	want := []uint64{uint64(6), uint64(120)}
	for i, n := range in {
		out := Factorial(n)
		if want[i] != out {
			t.Errorf("%v -> want:%v, out=%v", n, want[i], out)
		}
	}
}

func TestNthLexicographicPerm(t *testing.T) {
	in := []int{0, 1, 2}
	want := []int{2, 0, 1}
	out := NthLexicographicPerm(uint64(5), in)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
