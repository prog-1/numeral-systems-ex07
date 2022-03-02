package main

import (
	"strings"
	"testing"
)

var testBinNumSum = BinNumSum

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

func BenchmarkLongMul(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = testBinNumSum(benchA, benchB)
	}
}
