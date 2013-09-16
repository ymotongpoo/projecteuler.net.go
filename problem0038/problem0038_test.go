package main

import (
	"testing"
)

func TestIsConcatPandigital(t *testing.T) {
	in := []int64{int64(192384576), int64(918273645)}
	want := []bool{true, true}

	for i, num := range in {
		out := IsConcatPandigital(num)
		if want[i] != out {
			t.Error("want: %v, out: %v", want[i], out)
		}
	}
}
