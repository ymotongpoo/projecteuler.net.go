package main

import (
	"reflect"
	"testing"
)

const in = `
12 34 56
78 90 01
02 03 04
`

func TestGridToDoubleSlice(t *testing.T) {
	want := [][]int{
		[]int{12, 34, 56},
		[]int{78, 90, 1},
		[]int{2, 3, 4},
	}
	out := GridToDoubleSlice(in)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestProdHorizontal(t *testing.T) {
	want := 7020
	grid := GridToDoubleSlice(in)
	out := ProdHorizontal(grid, 2)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestProdVertical(t *testing.T) {
	want := 3060
	grid := GridToDoubleSlice(in)
	out := ProdVertical(grid, 2)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestProdLeftCross(t *testing.T) {
	want := 5040
	grid := GridToDoubleSlice(in)
	out := ProdRightCross(grid, 2)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
