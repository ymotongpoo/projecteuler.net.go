package main

import (
	"reflect"
	"testing"
)

func TestReadAllNumbers(t *testing.T) {
	in := []byte(`
12345
67890
12345
`)
	want := []int64{int64(12345), int64(67890), int64(12345)}
	out, err := BytesToNumbers(in)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want:%v, out=%v", want, out)
	}
}
