// The Fibonacci sequence is defined by the recurrence relation:
//
// Fn = Fn1 + Fn2, where F1 = 1 and F2 = 1.
// Hence the first 12 terms will be:
//
// F1 = 1
// F2 = 1
// F3 = 2
// F4 = 3
// F5 = 5
// F6 = 8
// F7 = 13
// F8 = 21
// F9 = 34
// F10 = 55
// F11 = 89
// F12 = 144
//
// The 12th term, F12, is the first term to contain three digits.
//
// What is the first term in the Fibonacci sequence to contain 1000 digits?
package main

import (
	"fmt"
	"math/big"
)

const Limit = 1000

func NumDigits(n *big.Int) int {
	count := 0
	ten := big.NewInt(10)
	for n.Cmp(big.NewInt(0)) == 1 {
		n = n.Div(n, ten)
		count++
	}
	return count
}

func Fibonacci(c chan *big.Int, quit chan int) {
	x := big.NewInt(1)
	y := big.NewInt(1)
	for {
		select {
		case c <- x:
			z := new(big.Int)
			tmp := x
			x = y
			y = z.Add(tmp, y)
		case <-quit:
			return
		}
	}
}

func Solver(n int) int {
	c := make(chan *big.Int)
	quit := make(chan int)

	var value *big.Int
	term := 1
	go func() {
		for {
			value = <-c
			fmt.Println("") // TODO(ymotongpoo): これがないと正しい回答にならない
			if NumDigits(value) < n {
				term++
			} else {
				quit <- 0
				break
			}
		}
	}()
	Fibonacci(c, quit)
	return term
}

func main() {
	fmt.Println(Solver(Limit))
}
