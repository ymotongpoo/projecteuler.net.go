package main

import (
	"reflect"
	"testing"
)

func TestPrimes(t *testing.T) {
	in := 30
	want := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	out := Primes(in)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestPow(t *testing.T) {
	want := 8
	out := Pow(2, 3)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
