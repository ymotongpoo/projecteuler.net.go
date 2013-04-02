package main

import (
	"testing"
)

func TestSumMultiplesBelow(t *testing.T) {
	want := 23
	out := SumMultiplesBelow(10)
	if want != out {
		t.Errorf("want: %v, out=%v\n", want, out)
	}
}
