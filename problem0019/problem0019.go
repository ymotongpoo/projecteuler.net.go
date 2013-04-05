// You are given the following information, but you may prefer to do some research for yourself.
//
// * 1 Jan 1900 was a Monday.
// * Thirty days has September,
//   April, June and November.
//   All the rest have thirty-one,
//   Saving February alone,
//   Which has twenty-eight, rain or shine.
//   And on leap years, twenty-nine.
// * A leap year occurs on any year evenly divisible by 4,
//   but not on a century unless it is divisible by 400.
//
// How many Sundays fell on the first of the month during the twentieth century
// (1 Jan 1901 to 31 Dec 2000)?
package main

import (
	"fmt"
)

var DaysInMonth = []int{
	31, // January
	28, // February
	31, // March
	30, // April
	31, // May
	30, // June
	31, // July
	31, // August
	30, // September
	31, // October
	30, // November
	31, // December
}

type Date struct {
	Year  int
	Month int
	Day   int
}

func isLeap(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

func (d Date) DaysFromStart() int {
	if d.Year-1900 < 0 {
		return 0
	}
	days := 0
	for year := 1900; year < d.Year; year++ {
		if isLeap(year) {
			days += 366
		} else {
			days += 365
		}
	}
	if d.Month > 1 {
		for month := 1; month < d.Month; month++ {
			days += DaysInMonth[month-1]
		}
	}
	days += d.Day - 1
	if isLeap(d.Year) && d.Month > 2 {
		days++
	}
	return days
}

func (d Date) DayOfDate() int {
	return d.DaysFromStart() % 7
}

func Solver() {
	count := 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			d := Date{y, m, 1}
			if d.DayOfDate() == 6 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	Solver()
}
