// Truncatable primes
//
// The number 3797 has an interesting property.
// Being prime itself, it is possible to continuously remove digits
// from left to right, and remain prime at each stage:
// 3797, 797, 97, and 7.
// Similarly we can work from right to left: 3797, 379, 37, and 3.
//
// Find the sum of the only eleven primes that are both
// truncatable from left to right and right to left.
//
// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
package main

import (
	"fmt"
)

const Buffer = 10000

func Generator(ch chan uint64) {
	for i := uint64(2); ; i++ {
		ch <- i
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

func IsTruncatablePrime(n uint64, primes []uint64) bool {
	l2r, r2l := n, n

	// Left to right
	for l2r > uint64(0) {
		found := false
		for _, p := range primes {
			if p == l2r {
				found = true
				break
			}
			if l2r%p == 0 {
				return false
			}
		}
		if !found {
			return false
		}
		l2r /= uint64(10)
	}
	// Right to left
	base := uint64(10)
	for n%base < n {
		found := false
		r2l = n % base
		for _, p := range primes {
			if p == r2l {
				found = true
				break
			}
			if r2l%p == 0 {
				return false
			}
		}
		if !found {
			return false
		}
		base *= uint64(10)
	}
	return true
}

func Solver() {
	in := make(chan uint64, Buffer)
	go Generator(in)

	count := 11
	sum := uint64(0)
	primes := []uint64{}
	for count > 0 {
		prime := <-in
		// Except 2, 3, 5, and 7
		if prime < uint64(10) {
			continue
		}
		primes = append(primes, prime)
		if IsTruncatablePrime(prime, primes) {
			sum += prime
			count--
		}
		out := make(chan uint64, Buffer)
		go Filter(in, out, prime)
		in = out
	}
	fmt.Println(sum)
}

func main() {
	Solver()
}
