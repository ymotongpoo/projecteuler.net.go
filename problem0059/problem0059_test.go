package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadCipher(t *testing.T) {
	in := []byte("11,22,33")
	b := bytes.NewReader(in)
	out, err := readCipher(b)
	if err != nil {
		t.Errorf("%v", err)
	}
	if reflect.DeepEqual(in, out) {
		t.Errorf("in:%v out:%v", in, out)
	}
}
