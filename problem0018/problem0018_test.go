package main

import (
	"reflect"
	"testing"
)

const in = `
3
7 4
2 4 6
8 5 9 3
`

func TestTriangleToData(t *testing.T) {
	want := [][]int{
		[]int{3},
		[]int{7, 4},
		[]int{2, 4, 6},
		[]int{8, 5, 9, 3},
	}
	out, err := TriangleToData(in)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestPreviousRow(t *testing.T) {
	want := [][]int{
		[]int{3},
		[]int{10, 7},
		[]int{12, 14, 13},
		[]int{20, 19, 23, 16},
	}
	out, _ := TriangleToData(in)
	for i := 0; i < len(data); i++ {
		LoadPreviousRow(out, i)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}

func TestMaximumTotal(t *testing.T) {
	want := 23
	data, _ := TriangleToData(in)
	out := FindMaximumTotal(data)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
