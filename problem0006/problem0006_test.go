package main

import (
	"testing"
)

func TestSumOfSquare(t *testing.T) {
	want := int64(385)
	out := SumOfSquare(10)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestSquareOfSum(t *testing.T) {
	want := int64(3025)
	out := SquareOfSum(10)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
