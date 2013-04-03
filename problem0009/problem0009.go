// A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,
//
// a^2 + b^2 = c^2
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
package main

import (
	"fmt"
)

const Limit = 1000

func IsPythagorean(x, y, z int) bool {
	if x*x+y*y == z*z {
		return true
	}
	return false
}

func main() {
	ave := Limit / 3 * 2
SOLVE:
	for x := 2; x < ave+1; x++ {
		for a := 1; a < x/2+1; a++ {
			b := x - a
			c := 1000 - x
			if IsPythagorean(a, b, c) {
				fmt.Println(a * b * c)
				break SOLVE
			}
		}
	}
}
