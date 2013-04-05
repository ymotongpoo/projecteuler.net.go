// Starting in the top left corner of a 2x2 grid,
// and only being able to move to the right and down,
// there are exactly 6 routes to the bottom right corner.
//
// How many such routes are there through a 20x20 grid?
package main

import (
	"fmt"
	"math/big"
)

const EdgeLength = 20

func CountRoutes(n int) *big.Int {
	nums := big.NewInt(int64(1))
	dens := big.NewInt(int64(1))
	for i := 1; i <= n; i++ {
		nums = nums.Mul(big.NewInt(int64(n+i)), nums)
		dens = dens.Mul(big.NewInt(int64(i)), dens)
	}
	ret := new(big.Int)
	return ret.Div(nums, dens)
}

func Solver(n int) {
	fmt.Println(CountRoutes(n))
}

func main() {
	Solver(EdgeLength)
}
