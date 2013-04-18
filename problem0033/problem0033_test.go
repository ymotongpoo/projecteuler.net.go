package main

import (
	"reflect"
	"testing"
)

func TestGCD(t *testing.T) {
	n := []int{10, 24, 98}
	m := []int{12, 32, 49}
	want := []int{2, 8, 49}
	for i := 0; i < len(n); i++ {
		out := GCD(n[i], m[i])
		if want[i] != out {
			t.Errorf("n:%v, m:%v, want:%v, out=%v",
				n[i], m[i], want[i], out)
		}
	}
}

func TestPower(t *testing.T) {
	base := []int{10, 10, 3}
	power := []int{0, 2, 3}
	want := []int{1, 100, 27}
	for i := 0; i < len(base); i++ {
		out := Power(base[i], power[i])
		if want[i] != out {
			t.Errorf("base:%v, power:%v, want:%v, out=%v",
				base[i], power[i], want[i], out)
		}
	}
}

func TestReduction(t *testing.T) {
	in := []*Fraction{
		&Fraction{32, 40},
		&Fraction{49, 98},
		&Fraction{88, 99},
	}
	want := []*Fraction{
		&Fraction{4, 5},
		&Fraction{1, 2},
		&Fraction{8, 9},
	}
	for i, f := range in {
		out := f.Reduction()
		if !reflect.DeepEqual(want[i], out) {
			t.Errorf("%v -> want:%v, out=%v", f, want[i], out)
		}
	}
}

func TestIsCurious(t *testing.T) {
	in := []*Fraction{
		&Fraction{49, 98},
		&Fraction{19, 98},
		&Fraction{21, 31},
		&Fraction{88, 99},
	}
	want := []bool{true, false, false, false}
	for i, f := range in {
		out := f.IsCurious()
		if want[i] != out {
			t.Errorf("%v -> want:%v, out=%v", f, want[i], out)
		}
	}
}

func TestEliminateCommonDigit(t *testing.T) {
	fractions := []*Fraction{
		&Fraction{49, 98},
		&Fraction{13, 23},
		&Fraction{88, 99},
	}
	want := []*Fraction{
		&Fraction{4, 8},
		&Fraction{1, 2},
		&Fraction{88, 99},
	}
	for i, f := range fractions {
		out, err := f.EliminateCommonDigit()
		if err != nil {
			t.Errorf("%v", err)
		}
		if !reflect.DeepEqual(want[i], out) {
			t.Errorf("%v -> want:%v, out:=%v",
				f, want[i], out)
		}
	}
}
