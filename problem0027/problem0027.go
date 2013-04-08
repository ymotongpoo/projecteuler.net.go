// Quadratic primes
//
// Euler published the remarkable quadratic formula:
//
//   n^2 + n + 41
//
// It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39.
// However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41,
// and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.
//
// Using computers, the incredible formula n^2 - 79n + 1601 was discovered,
// which produces 80 primes for the consecutive values n = 0 to 79.
// The product of the coefficients, -79 and 1601, is -126479.
//
// Considering quadratics of the form:
//
//   n^2 + an + b, where |a| < 1000 and |b| < 1000
//
//   where |n| is the modulus/absolute value of n
//   e.g. |11| = 11 and |4| = 4
//
// Find the product of the coefficients, a and b, for the quadratic expression
// that produces the maximum number of primes for consecutive values of n,
// starting with n = 0.
package main

import (
	"fmt"
)

func Generator(c chan uint64) {
	for i := uint64(2); ; i++ {
		c <- i
	}
}

func Filter(in <-chan uint64, out chan<- uint64, prime uint64) {
	for {
		n := <-in
		if n%prime != 0 {
			out <- n
		}
	}
}

func Primes(n int) map[uint64]struct{} {
	primes := make(map[uint64]struct{})
	index := 0
	in := make(chan uint64)
	go Generator(in)
	for {
		prime := <-in
		out := make(chan uint64)
		go Filter(in, out, prime)
		if index == n {
			break
		}
		primes[prime] = struct{}{}
		index++
		in = out
	}
	return primes
}

func CountQuadraticPrimes(a, b int, primes map[uint64]struct{}) int {
	for i := 0; ; i++ {
		value := i*i + a*i + b
		if value < 1 {
			return i
		}
		if _, ok := primes[uint64(value)]; !ok {
			return i
		}
	}
	return 0
}

func Solver() {
	var max_seq int
	var prod int
	primes := Primes(500)
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			seq := CountQuadraticPrimes(a, b, primes)
			if seq > max_seq {
				max_seq = seq
				prod = a * b
			}
		}
	}
	fmt.Println(prod)
}

func main() {
	Solver()
}
