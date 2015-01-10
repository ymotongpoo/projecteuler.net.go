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

func TestMinmax(t *testing.T) {
	in := []byte("ABCDEabcde")
	wantMin := int('A')
	wantMax := int('e')
	ints := make([]int, len(in))
	for i, c := range in {
		ints[i] = int(c)
	}
	outMin, outMax := minmax(ints)
	if wantMin != outMin {
		t.Errorf("want: %v, out: %v", wantMin, outMin)
	}
	if wantMax != outMax {
		t.Errorf("want: %v, out: %v", wantMax, outMax)
	}
}
