// If the numbers 1 to 5 are written out in words:
// one, two, three, four, five, then there are
// 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
// If all the numbers from 1 to 1000 (one thousand) inclusive
// were written out in words, how many letters would be used?
//
// NOTE: Do not count spaces or hyphens.
// For example, 342 (three hundred and forty-two) contains 23 letters
// and 115 (one hundred and fifteen) contains 20 letters.
// The use of "and" when writing out numbers is in compliance with British usage.
package main

import (
	"fmt"
)

type Pair struct {
	Number int
	Length int
}

var SingleNumbers = map[int]int{
	1: 3, // one
	2: 3, // two
	3: 5, // three
	4: 4, // four
	5: 4, // five
	6: 3, // six
	7: 5, // seven
	8: 5, // eight
	9: 4, // nine
}

var DoubleNumbers = []Pair{
	Pair{90, 6}, // ninety
	Pair{80, 6}, // eighty
	Pair{70, 7}, // seventy
	Pair{60, 5}, // sixty
	Pair{50, 5}, // fifty
	Pair{40, 5}, // forty
	Pair{30, 6}, // thirty
	Pair{20, 6}, // twenty
	Pair{19, 8}, // nineteen
	Pair{18, 8}, // eighteen
	Pair{17, 9}, // seventeen
	Pair{16, 7}, // sixteen
	Pair{15, 7}, // fifteen
	Pair{14, 8}, // fourteen
	Pair{13, 8}, // thirteen
	Pair{12, 6}, // twelve
	Pair{11, 6}, // eleven
	Pair{10, 3}, // ten
}

func LengthInWords(n int) int {
	length := 0
	if t := n / 1000; t != 0 {
		length += SingleNumbers[t] + 8 // "{t} thousand"
		n = n % 1000
	}
	if h := n / 100; h != 0 {
		length += SingleNumbers[h] + 7 // "{h} hundred"
		n = n % 100
		if n != 0 {
			length += 3 // "and"
		}
	}
	for _, p := range DoubleNumbers {
		if d := n / p.Number; d != 0 {
			length += p.Length
			n = n % p.Number
			break
		}
	}
	if v, ok := SingleNumbers[n]; ok {
		length += v
	}
	return length
}

func Solver() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += LengthInWords(i)
	}
	fmt.Println(sum)
	fmt.Println(LengthInWords(1000))
}

func main() {
	Solver()
}
