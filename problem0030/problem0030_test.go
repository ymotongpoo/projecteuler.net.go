package main

import (
	"reflect"
	"testing"
)

func TestFifthPower(t *testing.T) {
	in := []int{1, 2, 3}
	want := []int{1, 32, 243}
	for i, n := range in {
		out := FifthPower(n)
		if want[i] != out {
			t.Errorf("%v -> want:%v, out=%v", n, want[i], out)
		}
	}
}

func TestNumToDigits(t *testing.T) {
	in := []int{1634, 8208, 9474}
	want := [][]int{
		[]int{1, 6, 3, 4},
		[]int{8, 2, 0, 8},
		[]int{9, 4, 7, 4},
	}
	for i, n := range in {
		out := NumToDigits(n)
		if !reflect.DeepEqual(want[i], out) {
			t.Errorf("%v -> want:%v, out=%v", n, want[i], out)
		}
	}
}
