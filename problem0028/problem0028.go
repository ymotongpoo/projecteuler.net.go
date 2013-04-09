// Number spiral diagonals
//
// Starting with the number 1 and moving to the right in a clockwise direction
// a 5 by 5 spiral is formed as follows:
//
// [21] 22 23 24 [25]
// 20  [7] 8 [9] 10
// 19  6  [1]  2 11
// 18  [5] 4 [3] 12
// [17] 16 15 14 [13]
//
// It can be verified that the sum of the numbers on the diagonals is 101.
//
// What is the sum of the numbers on the diagonals in a
// 1001 by 1001 spiral formed in the same way?
package main

import (
	"fmt"
)

// SumOfDiag returns the sum of all diagonal number of n x n square.
func SumOfDiag(n int) uint64 {
	// Edge length must be odd number.
	if n%2 == 0 {
		return 0
	}
	level := (n + 1) / 2

	sum := uint64(0)
	for i := 0; i < level; i++ {
		if i == 0 {
			sum++
		} else {
			// Right top number is square of each level.
			rt := uint64(2*i+1) * uint64(2*i+1)
			sum += uint64(4)*rt - uint64(12*i)
		}
	}
	return sum
}

func Solver() {
	fmt.Println(SumOfDiag(1001))
}

func main() {
	Solver()
}
