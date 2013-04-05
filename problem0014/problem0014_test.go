package main

import (
	"testing"
)

func TestCollatzSeq(t *testing.T) {
	in := []uint64{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	for i := 0; i < len(in)-2; i++ {
		if in[i+1] != CollatzSeq(in[i]) {
			t.Errorf("want:%v, out=%v", in[i+1], CollatzSeq(in[i]))
		}
	}
}

func TestCountSequence(t *testing.T) {
	in := uint64(13)
	want := 10
	out := CountSequence(in)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
