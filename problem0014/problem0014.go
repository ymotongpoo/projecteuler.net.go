// The following iterative sequence is defined for the set of positive integers:
// is sequence (starting at 13 and finishing at 1) contains 10 terms.
//
//  n -> n/2 (n is even)
//  n -> 3n + 1 (n is odd)
//
// Using the rule above and starting with 13, we generate the following sequence:
//
//  13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
//
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
// Although it has not been proved yet (Collatz Problem),
// it is thought that all starting numbers finish at 1.
//
// Which starting number, under one million, produces the longest chain?
//
// NOTE: Once the chain starts the terms are allowed to go above one million.
package main

import (
	"fmt"
)

const Limit = 1000000

func CollatzSeq(n uint64) uint64 {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func CountSequence(n uint64) int {
	for i := 1; ; i++ {
		ret := CollatzSeq(n)
		if ret == 1 {
			return i + 1
		}
		n = ret
	}
	return 0
}

func Solver(limit int) {
	max := 0
	longest := 1
	for i := 1; i < limit; i++ {
		length := CountSequence(uint64(i))
		if length > max {
			max = length
			longest = i
		}
	}
	fmt.Println(longest)
}

func main() {
	Solver(Limit)
}
