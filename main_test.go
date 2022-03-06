package main

import (
	"testing"
)

func TestNumToBin(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    int64
		want string
	}{
		{"From README", 0, "0"},
		{"From README", 1, "1"},
		{"From README", 8, "1000"},
		{"From README", -1, "1111111111111111111111111111111111111111111111111111111111111111"},
		{"From README", -1, "1111111111111111111111111111111111111111111111111111111111111111"},
		{"From README", 1000_0000_0000_0000, "11100011010111111010100100110001101000000000000000"},
		{"From README", -1000_0000_0000_0001, "1111111111111100011100101000000101011011001110010111111111111111"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := NumToBin(tc.a); got != tc.want {
				t.Errorf("NumToBin(%v) = %v, want = %v", tc.a, got, tc.want)
			}
		})
	}
}

func TestBinNumSum(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"From README", "1", "11", "100"},
		{"From README", "11001", "1111111111111111111111111111111111111111111111111111111111110010", "1011"},
		{"From README", "1111111111111111111111111111111111111111111111111111111111110110", "111", "1111111111111111111111111111111111111111111111111111111111111101"},
		{"From README", "11", "1111111111111111111111111111111111111111111111111111111111111101", "0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := BinNumSum(tc.a, tc.b); got != tc.want {
				t.Errorf("BinToSum(%v) = %v, %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
