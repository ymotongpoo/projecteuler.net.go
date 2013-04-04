package main

import (
	"reflect"
	"testing"
)

func TestNthTriangleNumber(t *testing.T) {
	want := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	out := make([]int, 10)
	for i := 0; i < 10; i++ {
		out[i] = NthTriangleNumber(i + 1)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestPrimes(t *testing.T) {
	want := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	out := Primes(10)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestNumFactors(t *testing.T) {
	want := map[int]int{
		1:  1,
		3:  2,
		6:  4,
		10: 4,
		15: 4,
		21: 4,
		28: 6}
	primes := Primes(10)
	for w, v := range want {
		out := NumFactors(w, primes)
		if v != out {
			t.Errorf("want:%v, out=%v", w, out)
		}
	}
}
