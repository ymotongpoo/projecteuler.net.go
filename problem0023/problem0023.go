// A perfect number is a number for which the sum of its proper divisors
// is exactly equal to the number. For example, the sum of the proper divisors
// of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
//
// A number n is called deficient if the sum of its proper divisors is
// less than n and it is called abundant if this sum exceeds n.
//
// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16,
// the smallest number that can be written as the sum of two abundant numbers is 24.
// By mathematical analysis, it can be shown that all integers greater than 28123
// can be written as the sum of two abundant numbers.
// However, this upper limit cannot be reduced any further by analysis even though
// it is known that the greatest number that cannot be expressed as the sum of
// two abundant numbers is less than this limit.
//
// Find the sum of all the positive integers which cannot be written as the sum of
// two abundant numbers.
package main

import (
	"fmt"
)

const Limit = int64(28123)

func IsAbundant(n int64) bool {
	total := int64(0)
	for i := int64(1); i < n/int64(2)+int64(1); i++ {
		if n%i == 0 {
			total += i
		}
	}
	if int64(total) > n {
		return true
	}
	return false
}

func AbundantNumbersBelow(n int64) []int64 {
	nums := []int64{}
	for i := int64(1); i < n; i++ {
		if IsAbundant(i) {
			nums = append(nums, i)
		}
	}
	return nums
}

func HasAbundantPair(nums []int64, n int64) bool {
	for i, m := range nums {
		if m > n {
			break
		}
		left := n - m
		for _, k := range nums[i:] {
			if k == left {
				return true
			}
		}
	}
	return false
}

func Solver() {
	nums := AbundantNumbersBelow(Limit)
	total := int64(0)
	for i := int64(1); i < Limit; i++ {
		if !HasAbundantPair(nums, i) {
			total += i
		}
	}
	fmt.Println(total)
}

func main() {
	Solver()
}
