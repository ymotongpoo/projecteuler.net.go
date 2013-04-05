// n! means n  (n  1)  ...  3  2  1
//
// For example, 10! = 10  9  ...  3  2  1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
//
// Find the sum of the digits in the number 100!
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func Fraction(n int) *big.Int {
	ret := big.NewInt(int64(1))
	for i := 1; i <= n; i++ {
		ret = ret.Mul(ret, big.NewInt(int64(i)))
	}
	return ret
}

func SumOfDigits(n *big.Int) (int, error) {
	sum := 0
	for _, s := range n.String() {
		i, err := strconv.Atoi(string(s))
		if err != nil {
			return 0, err
		}
		sum += i
	}
	return sum, nil
}

func Solver() {
	answer, err := SumOfDigits(Fraction(100))
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}

func main() {
	Solver()
}
