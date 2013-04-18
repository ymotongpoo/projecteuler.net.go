// Digit factorials
//
// 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
//
// Find the sum of all numbers which are equal to the sum of
// the factorial of their digits.
//
// Note: as 1! = 1 and 2! = 2 are not sums they are not included.
package main

import (
	"fmt"
)

func Factorial(n int) uint64 {
	accu := uint64(1)
	for i := 2; i <= n; i++ {
		accu *= uint64(i)
	}
	return accu
}

func SumOfFactorial(ns []int) uint64 {
	sum := uint64(0)
	for _, n := range ns {
		sum += Factorial(n)
	}
	return sum
}

// NumToDigits returns digits in n in reverse order.
func NumToDigits(n uint64) []int {
	digits := []int{}
	for n > uint64(10) {
		digits = append(digits, int(n%uint64(10)))
		n /= uint64(10)
	}
	digits = append(digits, int(n))
	return digits
}

func Power(x, n int) uint64 {
	pow := uint64(1)
	for i := 0; i < n; i++ {
		pow *= uint64(x)
	}
	return pow
}

func Solver() {
	total := uint64(0)
	fact9 := Factorial(9)
	for i := uint64(3); i < fact9*10; i++ {
		if SumOfFactorial(NumToDigits(i)) == i {
			total += i
		}
	}
	fmt.Println(total)
}

func main() {
	Solver()
}
