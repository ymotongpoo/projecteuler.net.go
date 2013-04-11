// In England the currency is made up of pound, and pence,
// and there are eight coins in general circulation:
//
//   1p, 2p, 5p, 10p, 20p, 50p, 1pound (100p) and 2 pound (200p)
//
// It is possible to make 2 pound in the following way:
//
//   1x1pound+1x50p+2x20p+1x5p+1x2p+3x1p
//
// How many different ways can 2pound be made using
// any number of choice?
package main

import (
	"fmt"
)

var Coins = []int{200, 100, 50, 20, 10, 5, 2, 1}

func NumPatterns(pense int, coins []int) (pattens int) {
	if pense == 0 || len(coins) == 1 {
		return 1
	}
	coin := coins[0]
	patterns := 0
	for i := 0; pense >= coin*i; i++ {
		patterns += NumPatterns(pense-coin*i, coins[1:])
	}
	return patterns
}

func Solver() {
	fmt.Println(NumPatterns(200, Coins))
}

func main() {
	Solver()
}
