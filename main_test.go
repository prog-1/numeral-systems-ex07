package main

import "testing"

func TestNumToBin(t *testing.T) {
	for _, tc := range []struct {
		num    int64
		binnum string
	}{

		{0, "0"},
		{1, "1"},
		{8, "1000"},
		{-1, "1111111111111111111111111111111111111111111111111111111111111111"},
		{-1, "1111111111111111111111111111111111111111111111111111111111111111"},
		{-8, "1111111111111111111111111111111111111111111111111111111111111000"},
		{1000_0000_0000_0000, "11100011010111111010100100110001101000000000000000"},
		{-1000_0000_0000_0001, "1111111111111100011100101000000101011011001110010111111111111111"},
	} {
		if got := NumToBin(tc.num); got != tc.binnum {
			t.Errorf("Error: got = %v, want = %v", got, tc.binnum)
		}
	}
}
func TestBinNumSum(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		want string
	}{
		{"1", "11", "100"},
		{"11001" /* 25 */, "1111111111111111111111111111111111111111111111111111111111110010" /* -14 */, "1011" /* 11 */},
		{"1111111111111111111111111111111111111111111111111111111111110110", "111" /* 7 */, "1111111111111111111111111111111111111111111111111111111111111101" /* -3 */},
		{"11" /* 3 */, "1111111111111111111111111111111111111111111111111111111111111101" /* -3 */, "0" /* 0 */},
	} {
		t.Run("", func(t *testing.T) {
			if got := BinNumSum(tc.a, tc.b); got != tc.want {
				t.Errorf("Error: got = %v, want %v", got, tc.want)
			}
		})
	}
}