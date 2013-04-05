// The sequence of triangle numbers is generated by adding the natural numbers.
// So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28.
// The first ten terms would be:
//
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
//
// Let us list the factors of the first seven triangle numbers:
//
// 1: 1
// 3: 1,3
// 6: 1,2,3,6
// 10: 1,2,5,10
// 15: 1,3,5,15
// 21: 1,3,7,21
// 28: 1,2,4,7,14,28
// We can see that 28 is the first triangle number to have over five divisors.
//
// What is the value of the first triangle number to have over five hundred divisors?
package main

import (
	"fmt"
)

// NthTriangleNumber returns the value of nth triangle number.
func NthTriangleNumber(n int) int {
	if n%2 == 0 {
		return (n + 1) * n / 2
	}
	return (n*n + n + 1) / 2
}

func Generator(c chan int) {
	for i := 2; ; i++ {
		c <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		n := <-in
		if n%prime != 0 {
			out <- n
		}
	}
}

// Primes returns a slice of primes below n.
func Primes(n int) []int {
	in := make(chan int)
	go Generator(in)
	primes := []int{}
	for {
		prime := <-in
		primes = append(primes, prime)
		if len(primes) == n {
			break
		}
		out := make(chan int)
		go Filter(in, out, prime)
		in = out
	}
	return primes
}

// Factors returns the number of n's factors.
func NumFactors(n int, primes []int) int {
	result := 1
	for _, p := range primes {
		for i := 0; ; i++ {
			if n%p != 0 {
				result *= i + 1
				break
			}
			n /= p
		}
	}
	return result
}

func main() {
	primes := Primes(500)
	for i := 1; ; i++ {
		t := NthTriangleNumber(i)
		if NumFactors(t, primes) > 500 {
			fmt.Println(t)
			return
		}
	}
}