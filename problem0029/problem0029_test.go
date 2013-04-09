package main

import (
	"reflect"
	"testing"
)

func TestPrimesBelow(t *testing.T) {
	want := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	out := PrimesBelow(30)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestStringfy(t *testing.T) {
	a := []int{2, 3, 4, 6}
	b := []int{2, 3, 4, 3}
	want := []string{
		"2^2",
		"3^3",
		"2^8",
		"2^3*3^3",
	}
	primes := PrimesBelow(10)
	for i, w := range want {
		out := Stringfy(primes, a[i], b[i])
		if w != out {
			t.Errorf("%v^%v -> want:%v, out=%v", a[i], b[i], w, out)
		}
	}
}
