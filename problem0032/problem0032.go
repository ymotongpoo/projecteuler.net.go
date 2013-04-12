// Pandigital products
//
// We shall say that an n-digit number is pandigital
// if it makes use of all the digits 1 to nexactly once;
// for example, the 5-digit number, 15234, is 1 through 5 pandigital.
//
// The product 7254 is unusual, as the identity, 
// 39 x 186 = 7254, containing multiplicand, multiplier,
// and product is 1 through 9 pandigital.
//
// Find the sum of all products whose multiplicand/multiplier/product
// identity can be written as a 1 through 9 pandigital.
//
// HINT: Some products can be obtained in more than one way
// so be sure to only include it once in your sum.
package main

import (
	"fmt"
)

var Digits = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Permutuation returns all candidats choose choose out of from.
func Permutuation(from []int, choose int) [][]int {
	if choose > len(from) {
		return Permutuation(from, len(from))
	}
	if choose == 1 {
		candidates := make([][]int, len(from))
		for i, n := range from {
			candidates[i] = []int{n}
		}
		return candidates
	}
	candidates := [][]int{}
	for i, n := range from {
		left := make([]int, 0, len(from)-1)
		left = append(left, from[:i]...)
		left = append(left, from[i+1:]...)
		subCandidates := Permutuation(left, choose-1)
		for _, s := range subCandidates {
			out := make([]int, 0, choose)
			out = append(out, n)
			out = append(out, s...)
			candidates = append(candidates, out)
		}
	}
	return candidates
}

// Solver works as follows:
//  1. List up all permutuation candidates from 1..9
//  2. Chop out multiplicand, multiplier and product candidate
//     from a permutuation of 1.
//  3. Check if the product is already listed up.
func Solver() {
	sum := 0

	candidates := Permutuation(Digits, len(Digits))
	products := map[int]struct{}{}

	// Case 1) 3 digits x 2 digits = 4 digits
	for _, c := range candidates {
		multiplicand := c[0]*100 + c[1]*10 + c[2]
		multiplier := c[3]*10 + c[4]
		prod := c[5]*1000 + c[6]*100 + c[7]*10 + c[8]
		if prod == multiplicand*multiplier {
			if _, ok := products[prod]; !ok {
				products[prod] = struct{}{}
				sum += prod
			}
		}
	}

	// Case 2) 1 digits x 4 digits = 4 digits
	for _, c := range candidates {
		multiplicand := c[0]
		multiplier := c[1]*1000 + c[2]*100 + c[3]*10 + c[4]
		prod := c[5]*1000 + c[6]*100 + c[7]*10 + c[8]
		if prod == multiplicand*multiplier {
			if _, ok := products[prod]; !ok {
				products[prod] = struct{}{}
			}
			sum += prod
		}
	}

	fmt.Println(sum)
}

func main() {
	Solver()
}
