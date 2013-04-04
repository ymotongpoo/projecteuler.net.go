package main

import (
	"testing"
)

func TestSumOfPrimeBelow(t *testing.T) {
	want := uint64(17)
	out := SumOfPrimeBelow(10)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
