// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?
package main

import (
	"fmt"
)

func Generate(ch chan<- uint64) {
	for i := uint64(2); ; i++ {
		ch <- i
	}
}

func Filter(in <-chan uint64, out chan<- uint64, prime uint64) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func NthPrime(n int) uint64 {
	in := make(chan uint64)
	go Generate(in)
	for i := 0; ; i++ {
		prime := <-in
		if i == n {
			return prime
		}
		out := make(chan uint64)
		go Filter(in, out, prime)
		in = out
	}
	return uint64(0)
}

func main() {
	fmt.Println(NthPrime(10000))
}
