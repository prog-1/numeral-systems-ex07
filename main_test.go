package main

import (
	"strings"
	"testing"
)

var testBinNumSum = BinNumSum
var testNumToBin = NumToBin

func TestNumToBin(t *testing.T) {
	for _, tc := range []struct {
		num  int64
		want string
	}{
		{0, "0"},
		{1, "1"},
		{3, "11"},
		{4, "100"},
		{8, "1000"},
		{-1, "1111111111111111111111111111111111111111111111111111111111111111"},
		{-1, "1111111111111111111111111111111111111111111111111111111111111111"},
		{1000_0000_0000_0000, "11100011010111111010100100110001101000000000000000"},
		{-1000_0000_0000_0001, "1111111111111100011100101000000101011011001110010111111111111111"},
	} {
		if got := testNumToBin(tc.num); got != tc.want {
			t.Errorf("NumToBin(%v) = %v, want = %v", tc.num, got, tc.want)
		}
	}
}

func TestBinNumSum(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want string
	}{
		{"1", "11", "100"},
		{"11", "10", "101"},
		{"11000", "1011111", "1110111"},
		{"11001" /* 25 */, "1111111111111111111111111111111111111111111111111111111111110010" /* -14 */, "1011" /* 11 */},
		{"1111111111111111111111111111111111111111111111111111111111110110", "111" /* 7 */, "1111111111111111111111111111111111111111111111111111111111111101" /* -3 */},
		{"11", "1111111111111111111111111111111111111111111111111111111111111101", "0" /* 0 */},
	} {
		if got := testBinNumSum(tc.a, tc.b); got != tc.want {
			t.Errorf("BinNumSum(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
		}
	}
}

var (
	benchA = strings.Repeat("1", 10000)
	benchB = strings.Repeat("9", 1)
)

func BenchmarkBinNumSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = testBinNumSum(benchA, benchB)
	}
}
