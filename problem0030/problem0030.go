// Digit fifth powers
//
// Surprisingly there are only three numbers that can be written
// as the sum of fourth powers of their digits:
//
//   1634 = 1^4 + 6^4 + 3^4 + 4^4
//   8208 = 8^4 + 2^4 + 0^4 + 8^4
//   9474 = 9^4 + 4^4 + 7^4 + 4^4
//   As 1 = 1^4 is not a sum it is not included.
//
// The sum of these numbers is 1634 + 8208 + 9474 = 19316.
//
// Find the sum of all the numbers that can be written
// as the sum of fifth powers of their digits.
package main

import (
	"fmt"
	"strconv"
)

func FifthPower(n int) int {
	return n * n * n * n * n
}

func NumToDigits(n int) []int {
	intStr := strconv.Itoa(n)
	ret := make([]int, len(intStr))
	for i, c := range intStr {
		value, err := strconv.Atoi(string(c))
		if err != nil {
			return nil
		}
		ret[i] = value
	}
	return ret
}

func Solver() {
	total := 0
	for i := 2; i < 500000; i++ {
		digits := NumToDigits(i)
		sum := 0
		for _, d := range digits {
			sum += FifthPower(d)
		}
		if sum == i {
			total += sum
		}
	}
	fmt.Println(total)
}

func main() {
	Solver()
}
