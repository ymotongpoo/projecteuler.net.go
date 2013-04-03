// 2520 is the smallest number that can be divided by each of the numbers
// from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible
// by all of the numbers from 1 to 20?
package main

import (
	"fmt"
)

const Limit = 20

func Generate(c chan<- int) {
	for i := 2; ; i++ {
		c <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func Primes(max int) []int {
	primes := []int{}
	in := make(chan int)
	go Generate(in)
	for i := 0; ; i++ {
		prime := <-in
		if prime > max {
			break
		}
		primes = append(primes, prime)
		out := make(chan int)
		go Filter(in, out, prime)
		in = out
	}
	return primes
}

func Pow(x, n int) int {
	ret := 1
	for n > 0 {
		ret *= x
		n--
	}
	return ret
}

func main() {
	primes := Primes(Limit)
	factors := make(map[int]int)
	for i := 2; i < Limit+1; i++ {
		number := i
		for _, p := range primes {
			count := 0
			if p > number {
				break
			}
			for number%p == 0 {
				number /= p
				count++
				if count > factors[p] {
					factors[p] = count
				}
			}
		}
	}

	answer := 1
	for k, v := range factors {
		answer *= Pow(k, v)
	}

	fmt.Println(answer)
}
