package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadCipher(t *testing.T) {
	in := []byte("11,22,33")
	want := []int{11, 22, 33}
	b := bytes.NewReader(in)
	out, err := readCipher(b)
	if err != nil {
		t.Errorf("%v", err)
	}
	if reflect.DeepEqual(want, out) {
		t.Errorf("want: %v out: %v", want, out)
	}
}

func TestDecrypt(t *testing.T) {
	in := []byte("ABCABC")
	key := []byte("*+,")
	want := []byte("kkkkkk")
	out, err := decrypt(key, in)
	if err != nil {
		t.Errorf("%v", err)
	}
	if reflect.DeepEqual(want, out) {
		t.Errorf("want: %v, out: %v", want, out)
	}
}
