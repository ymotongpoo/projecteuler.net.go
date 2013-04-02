package main

import (
	"testing"
)

func TestFindLargestFactor(t *testing.T) {
	in := int64(13195)
	out := FindLargestFactor(in)
	want := int64(29)
	if out != want {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
