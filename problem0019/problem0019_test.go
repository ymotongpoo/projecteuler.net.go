package main

import (
	"testing"
)

func Test_isLeap(t *testing.T) {
	in := []int{1900, 1901, 1904, 2000}
	want := []bool{false, false, true, true}
	for i, y := range in {
		out := isLeap(y)
		if want[i] != out {
			t.Errorf("%v: want:%v, out=%v", y, want[i], out)
		}
	}
}

func TestDaysFromStart(t *testing.T) {
	in := []Date{
		Date{1900, 1, 1},
		Date{1901, 1, 1},
		Date{1904, 1, 1},
		Date{1904, 2, 1},
		Date{1904, 3, 31},
	}
	want := []int{0, 365, 365 * 4, 365*4 + 31, 365*4 + 31 + 29 + 30}
	for i, d := range in {
		out := d.DaysFromStart()
		if want[i] != out {
			t.Errorf("%v: want:%v, out=%v", d, want[i], out)
		}
	}
}
