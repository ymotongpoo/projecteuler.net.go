package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	in := []int64{4, 10, 11, 100, 111, 121, 999, 1000, 1001, 1111, 1221}
	want := []bool{true, false, true, false, true, true, true, false, true, true, true}
	out := make([]bool, len(in))
	for i, n := range in {
		out[i] = IsPalindrome(n)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
