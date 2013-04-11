package main

import (
	"testing"
)

func TestNumPatterns(t *testing.T) {
	coins := []int{5, 2, 1}
	pense := 5
	want := 1 + (1 + 1 + 1)
	out := NumPatterns(pense, coins)
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
