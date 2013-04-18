// Digit canceling fractions
//
// The fraction 49/98 is a curious fraction, as an inexperienced
// mathematician in attempting to simplify it may incorrectly
// believe that 49/98 = 4/8, which is correct, is obtained by
// cancelling the 9s.
//
// We shall consider fractions like, 30/50 = 3/5,
// to be trivial examples. 
//
// There are exactly four non-trivial examples of this type of fraction,
// less than one in value, and containing two digits in
// the numerator and denominator.
//
// If the product of these four fractions is given in its lowest
// common terms, find the value of the denominator.
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Fraction struct {
	Num int
	Den int
}

func GCD(m, n int) int {
	if m < n {
		m, n = n, m
	}
	if n == 0 {
		return m
	}
	return GCD(n, m%n)
}

func Power(x, n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret = ret * x
	}
	return ret
}

func (f *Fraction) EliminateCommonDigit() (*Fraction, error) {
	num := strconv.Itoa(f.Num)
	den := strconv.Itoa(f.Den)
	var newNum string
CANCELING:
	for _, n := range num {
		for j, d := range den {
			if n == d {
				den = den[:j] + den[j+1:]
				continue CANCELING
			}
		}
		newNum += string(n)
	}

	numN, err := strconv.Atoi(newNum)
	if err != nil {
		return nil, err
	}
	denN, err := strconv.Atoi(den)
	if err != nil {
		return nil, err
	}
	return &Fraction{numN, denN}, nil
}

func (f *Fraction) Reduction() *Fraction {
	gcd := GCD(f.Num, f.Den)
	if gcd == 0 {
		fmt.Println(f)
		return &Fraction{1, 1}
	}
	return &Fraction{f.Num / gcd, f.Den / gcd}
}

func (f *Fraction) IsCurious() bool {
	r := f.Reduction()
	e, err := f.EliminateCommonDigit()
	if err != nil {
		return false
	}
	er := e.Reduction()
	if reflect.DeepEqual(f, e) {
		return false
	}
	return reflect.DeepEqual(r, er)
}

func NonTrivialFractions(digits int, c chan<- Fraction) {
	start := Power(10, digits-1)
	end := Power(10, digits)
	for d := start; d < end; d++ {
		for n := start; n < d; n++ {
			if n%10 != 0 || d%10 != 0 {
				c <- Fraction{n, d}
			}
		}
	}
	close(c)
}

func Solver() {
	fractions := make(chan Fraction, 1000)
	go NonTrivialFractions(2, fractions)

	r := Fraction{1, 1}
	for f := range fractions {
		if f.IsCurious() == true {
			fmt.Println(f)
			r = Fraction{r.Num * f.Num, r.Den * f.Den}
		}
	}
	fmt.Println(r.Reduction().Den)
}

func main() {
	Solver()
}
