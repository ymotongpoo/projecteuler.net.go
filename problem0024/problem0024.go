// A permutation is an ordered arrangement of objects.
// For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
// If all of the permutations are listed numerically or alphabetically,
// we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:
//
// 012   021   102   120   201   210
//
// What is the millionth lexicographic permutation of
// the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
package main

import (
	"fmt"
	"sort"
)

const Target = uint64(1000000)

var Digits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func Factorial(n int) uint64 {
	mul := uint64(1)
	if n > 0 {
		for i := 1; i <= n; i++ {
			mul *= uint64(i)
		}
	}
	return mul
}

// NumCandidates returns number of candidates if n numbers out of m elements
// are fixed for permutation.
func NumCandidates(n, m int) uint64 {
	return Factorial(m - n)
}

// NthLexicographicPerm returns nth lexicographic permutation from nums.
func NthLexicographicPerm(n uint64, nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)
	baseCands := Factorial(len(nums) - 1)

	index := 0
	for n > baseCands {
		n -= baseCands
		index++
	}

	result := make([]int, len(nums))
	result[0] = nums[index]
	// Create index'th element extracted slice from nums.
	newNums := make([]int, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		switch {
		case i < index:
			newNums[i] = nums[i]
		case i > index:
			newNums[i-1] = nums[i]
		}
	}

	subs := NthLexicographicPerm(n, newNums)
	for i := 0; i < len(subs); i++ {
		result[i+1] = subs[i]
	}
	return result

}

func Solver() {
	fmt.Println(NthLexicographicPerm(Target, Digits))
}

func main() {
	Solver()
}
