package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

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
	fmt.Println("spam")
}
