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

func TestSquare(t *testing.T) {
	want := []int{1, 4, 9, 16, 25}
	out := make([]int, len(want))
	for i := range want {
		out[i] = square(i + 1)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want: %v, out: %v", want, out)
	}
}

func TestPentagonal(t *testing.T) {
	want := []int{1, 5, 12, 22, 35}
	out := make([]int, len(want))
	for i := range want {
		out[i] = pentagonal(i + 1)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want: %v, out: %v", want, out)
	}
}
