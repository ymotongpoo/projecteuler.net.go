package main

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	want := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	out := make([]int, len(want))
	go func() {
		for i := 0; i < len(want); i++ {
			out[i] = <-c
		}
		quit <- 0
	}()
	Fibonacci(c, quit)
	if !reflect.DeepEqual(want, out) {
		t.Errorf("want: %v, out=%v", want, out)
	}
}
