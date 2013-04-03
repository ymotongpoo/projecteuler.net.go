package main

import (
	"testing"
)

func TestNthPrime(t *testing.T) {
	want := uint64(13)
	out := NthPrime(5)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
