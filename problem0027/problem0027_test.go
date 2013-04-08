package main

import (
	"reflect"
	"testing"
)

func TestPrimes(t *testing.T) {
	want := map[uint64]struct{}{
		uint64(2):  struct{}{},
		uint64(3):  struct{}{},
		uint64(5):  struct{}{},
		uint64(7):  struct{}{},
		uint64(11): struct{}{},
		uint64(13): struct{}{},
		uint64(17): struct{}{},
		uint64(19): struct{}{},
		uint64(23): struct{}{},
		uint64(29): struct{}{},
	}
	out := Primes(10)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%V, out=%v", want, out)
	}
}

func TestCountQuadraticPrimes(t *testing.T) {
	primes := Primes(500)
	a, b := 1, 41
	want := 40
	out := CountQuadraticPrimes(a, b, primes)
	if want != out {
		t.Errorf("want:%V, out=%v", want, out)
	}

	a, b = -79, 1601
	want = 80
	out = CountQuadraticPrimes(a, b, primes)
	if want != out {
		t.Errorf("want:%V, out=%v", want, out)
	}
}
