package main

import (
	"reflect"
	"testing"
)

func TestTriangle(t *testing.T) {
	want := []int{1, 3, 6, 10, 15}
	out := make([]int, len(want))
	for i := range want {
		out[i] = triangle(i + 1)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want: %v, out: %v", want, out)
	}
}
