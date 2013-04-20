// Double-base palindromes
//
// The decimal number, 585 = 1001001001_2 (binary),
// is palindromic in both bases.
//
// Find the sum of all numbers, less than one million,
// which are palindromic in base 10 and base 2.
//
// (Please note that the palindromic number, in either base,
//  may not include leading zeros.)
package main

import (
	"fmt"
	"strconv"
)

const Limit = 1000000

func IsPalindromic(n uint64, base int) bool {
	str := strconv.FormatUint(n, base)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func Generator(ch chan uint64, limit uint64) {
	for i := uint64(0); i < limit; i++ {
		ch <- i
	}
	close(ch)
}

func Filter(in <-chan uint64, out chan<- uint64) {
	for n := range in {
		if IsPalindromic(n, 10) && IsPalindromic(n, 2) {
			out <- n
		}
	}
	close(out)
}

func Solver() {
	in := make(chan uint64, 10000)
	out := make(chan uint64, 10000)
	go Generator(in, Limit)
	go Filter(in, out)
	sum := uint64(0)
	for n := range out {
		sum += n
	}
	fmt.Println(sum)
}

func main() {
	Solver()
}
