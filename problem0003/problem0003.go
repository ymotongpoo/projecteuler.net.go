// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?
//
// Original code is in http://golang.org/doc/play/sieve.go.
package main

import (
	"fmt"
)

const Target = int64(600851475143)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int64) {
	for i := int64(2); ; i++ {
		ch <- i
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int64, out chan<- int64, prime int64) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func FindLargestFactor(target int64) int64 {
	ch := make(chan int64) // Create a new channel.
	go Generate(ch)        // Launch Generate goroutine.
	var result int64
	for {
		prime := <-ch
		if prime > target {
			break
		}

		if target%prime == 0 {
			result = prime
			for target%prime == 0 {
				target /= prime
			}
		}
		ch1 := make(chan int64)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return result
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	fmt.Println(FindLargestFactor(Target))
}
