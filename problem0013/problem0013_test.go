package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestReadAllNumbers(t *testing.T) {
	in := []byte(`
12345
67890
12345
`)
	want := []*big.Int{
		big.NewInt(12345),
		big.NewInt(67890),
		big.NewInt(12345),
	}
	out, err := BytesToNumbers(in)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
