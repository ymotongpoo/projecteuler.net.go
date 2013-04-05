package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCountRoutes(t *testing.T) {
	want := big.NewInt(6)
	out := CountRoutes(2)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
