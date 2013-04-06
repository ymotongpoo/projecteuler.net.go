package main

import (
	"testing"
)

func TestWordScore(t *testing.T) {
	want := 53
	out := WordScore("COLIN")
	if want != out {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
