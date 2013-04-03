package main

import (
	"testing"
)

func TestIsPythagorean(t *testing.T) {
	if !IsPythagorean(3, 4, 5) {
		t.Errorf("wrong calculation.")
	}
}
