// A unit fraction contains 1 in the numerator.
// The decimal representation of the unit fractions with denominators 2 to 10 are given:
//
// 1/2	= 	0.5
// 1/3	= 	0.(3)
// 1/4	= 	0.25
// 1/5	= 	0.2
// 1/6	= 	0.1(6)
// 1/7	= 	0.(142857)
// 1/8	= 	0.125
// 1/9	= 	0.(1)
// 1/10	= 	0.1
//
// Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle.
// It can be seen that 1/7 has a 6-digit recurring cycle.
//
// Find the value of d < 1000 for which 1/d contains the longest recurring cycle
// in its decimal fraction part.
package main

import (
	"fmt"
)

const Limit = 1000

// RecurringCycleLength returns the length or recurring cycle length
// in 1/n unit fraction.
func RecurringCycleLength(n int) int {
	length := 0
	buf := map[int]struct{}{}
	mod := 1
	for mod != 0 {
		base := mod * 10
		mod = base % n
		if mod == 0 {
			length = 0
			break
		}
		_, ok := buf[mod]
		if ok {
			break
		} else {
			buf[mod] = struct{}{}
		}
		length++
	}
	return length
}

func Solver(n int) {
	max := 0
	max_n := 0
	for i := 2; i < n; i++ {
		length := RecurringCycleLength(i)
		if length > max {
			max = length
			max_n = i
		}
	}
	fmt.Println(max_n)
}

func main() {
	Solver(Limit)
}
