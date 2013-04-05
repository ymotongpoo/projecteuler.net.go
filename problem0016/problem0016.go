// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//
// What is the sum of the digits of the number 2^1000?
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// Pow return nth power of x.
func Pow(x *big.Int, n int) *big.Int {
	ret := big.NewInt(int64(1))
	for i := 0; i < n; i++ {
		ret = ret.Mul(ret, x)
	}
	return ret
}

func SumOfDigits(n *big.Int) (int, error) {
	digits := n.String()
	sum := 0
	for _, s := range digits {
		n, err := strconv.ParseInt(string(s), 10, 0)
		if err != nil {
			return 0, err
		}
		sum += int(n)
	}
	return sum, nil
}

func Solver() {
	n, err := SumOfDigits(Pow(big.NewInt(2), 1000))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func main() {
	Solver()
}
