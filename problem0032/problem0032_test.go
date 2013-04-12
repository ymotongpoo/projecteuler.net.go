package main

import (
	"reflect"
	"testing"
)

func TestPermutuation(t *testing.T) {
	in := []int{1, 2, 3, 4}
	want := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 4},
		[]int{1, 3, 2},
		[]int{1, 3, 4},
		[]int{1, 4, 2},
		[]int{1, 4, 3},
		[]int{2, 1, 3},
		[]int{2, 1, 4},
		[]int{2, 3, 1},
		[]int{2, 3, 4},
		[]int{2, 4, 1},
		[]int{2, 4, 3},
		[]int{3, 1, 2},
		[]int{3, 1, 4},
		[]int{3, 2, 1},
		[]int{3, 2, 4},
		[]int{3, 4, 1},
		[]int{3, 4, 2},
		[]int{4, 1, 2},
		[]int{4, 1, 3},
		[]int{4, 2, 1},
		[]int{4, 2, 3},
		[]int{4, 3, 1},
		[]int{4, 3, 2},
	}
	out := Permutuation(in, 3)
	if len(out) == 0 {
		t.Errorf("no candidates")
	}
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("recovered: %v", r)
		}
	}()
	for _, o := range out {
		found := false
		for _, w := range want {
			if reflect.DeepEqual(o, w) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("not found: candidate:%v, want:%v", o, want)
		}
	}
}
