package main

import (
	"reflect"
	"testing"
)

func TestSeqToSlice(t *testing.T) {
	in := `
12345
67890
`
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	out, err := SeqToSlice(in)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
