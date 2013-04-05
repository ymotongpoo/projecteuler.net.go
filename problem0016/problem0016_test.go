package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestPow(t *testing.T) {
	want := big.NewInt(32768)
	out := Pow(big.NewInt(2), 15)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestSumOfDigits(t *testing.T) {
	want := 26
	out, err := SumOfDigits(Pow(big.NewInt(2), 15))
	if err != nil {
		t.Errorf("%v", err)
	}
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
