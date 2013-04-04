// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
package main

import (
	"fmt"
)

func Generator(c chan<- uint64) {
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

func SumOfPrimeBelow(n uint64) uint64 {
	in := make(chan uint64)
	go Generator(in)
	sum := uint64(0)
	for {
		prime := <-in
		if prime > n {
			break
		}
		sum += prime
		out := make(chan uint64)
		go Filter(in, out, prime)
		in = out
	}
	return sum
}

func main() {
	fmt.Println(SumOfPrimeBelow(2000000))
}
