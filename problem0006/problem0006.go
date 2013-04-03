// The sum of the squares of the first ten natural numbers is,
//
// 1^2 + 2^2 + ... + 10^2 = 385
//
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)^2 = 552 = 3025
//
// Hence the difference between the sum of the squares of
// the first ten natural numbers and the square of the sum is 3025  385 = 2640.
//
// Find the difference between the sum of the squares of
// the first one hundred natural numbers and the square of the sum.
package main

import (
	"fmt"
)

func SumOfSquare(max int64) int64 {
	sum := int64(0)
	for i := int64(0); i < max; i++ {
		sum += (i + int64(1)) * (i + int64(1))
	}
	return sum
}

func SquareOfSum(max int64) int64 {
	sum := int64(0)
	for i := int64(0); i < max; i++ {
		sum += i + int64(1)
	}
	return sum * sum
}

func main() {
	limit := int64(100)
	fmt.Println(SquareOfSum(limit) - SumOfSquare(limit))
}
