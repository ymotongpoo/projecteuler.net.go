package main

import (
	"math/big"
	"testing"
)

func TestNumDigits(t *testing.T) {
	in := big.NewInt(144)
	want := 3
	out := NumDigits(in)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestSolver(t *testing.T) {
	want := 12
	out := Solver(3)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
