package main

import (
	"reflect"
	"testing"
)

func testPolygonal(t *testing.T, want []int, f func(int) int) {
	out := make([]int, len(want))
	for i := range want {
		out[i] = f(i + 1)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want: %v, out: %v", want, out)
	}
}

func TestTriangle(t *testing.T) {
	want := []int{1, 3, 6, 10, 15}
	testPolygonal(t, want, triangle)
}

func TestSquare(t *testing.T) {
	want := []int{1, 4, 9, 16, 25}
	testPolygonal(t, want, square)
}

func TestPentagonal(t *testing.T) {
	want := []int{1, 5, 12, 22, 35}
	testPolygonal(t, want, pentagonal)
}

func TestHexagonal(t *testing.T) {
	want := []int{1, 6, 15, 28, 45}
	testPolygonal(t, want, hexagonal)
}

func TestHeptagonal(t *testing.T) {
	want := []int{1, 7, 18, 34, 55}
	testPolygonal(t, want, heptagonal)
}

func TestOctagonal(t *testing.T) {
	want := []int{1, 8, 21, 40, 65}
	testPolygonal(t, want, octagonal)
}
