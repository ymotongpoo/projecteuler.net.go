package main

import (
	"testing"
)

func primesBelow(n uint64) (primes []uint64) {
	in := make(chan uint64, Buffer)
	go Generator(in)

	for {
		prime := <-in
		if prime > n {
			break
		}
		primes = append(primes, prime)
		out := make(chan uint64, Buffer)
		go Filter(in, out, prime)
		in = out
	}
	return primes
}

func TestIsTruncatablePrime(t *testing.T) {
	in := []uint64{uint64(29), uint64(3797)}
	want := []bool{false, true}
	primes := primesBelow(uint64(4000))
	for i, n := range in {
		out := IsTruncatablePrime(n, primes)
		if want[i] != out {
			t.Errorf("%v -> want:%v, out=%v", n, want[i], out)
		}
	}
}
