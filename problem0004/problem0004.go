// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

const Limit = 1000

func IsPalindrome(n int64) bool {
	s := strconv.FormatInt(n, 10)

	var is_palindrome func(s string) bool
	is_palindrome = func(s string) bool {
		if utf8.RuneCountInString(s) <= 1 {
			return true
		}
		first, sizeOfFirst := utf8.DecodeRuneInString(s)
		last, sizeOfLast := utf8.DecodeLastRuneInString(s)
		if first != last {
			return false
		}
		return is_palindrome(s[sizeOfFirst : len(s)-sizeOfLast])
	}

	return is_palindrome(s)
}

func main() {
	max := 0
	for i := 0; i < Limit; i++ {
		for j := 0; j < Limit; j++ {
			n := i * j
			if IsPalindrome(int64(n)) && max < n {
				max = n
			}
		}
	}
	fmt.Println(max)
}
