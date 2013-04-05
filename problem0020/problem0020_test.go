package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFraction(t *testing.T) {
	want := big.NewInt(int64(3628800))
	out := Fraction(10)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestSumOfDigits(t *testing.T) {
	want := 27
	out, err := SumOfDigits(Fraction(10))
	if err != nil {
		t.Errorf("%v", err)
	}
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
