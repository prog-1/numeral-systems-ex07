package main

import (
	"fmt"
	"math"
	"strings"
)

const int64min = 9223372036854775807

func addZeros(a string, time int) string {
	return strings.Repeat("0", time) + a
}

func Normalize(a, b string) (string, string) {
	if len(a) > len(b) {
		return a, addZeros(b, len(a)-len(b))
	}
	return addZeros(a, len(b)-len(a)), b
}

func findStartpoint(num int64, base int) (power int) {
	for math.Pow(float64(base), float64(power)) <= float64(num) {
		power++
	}
	return power - 1
}

func main() {
	fmt.Println(NumToBin(-1))
}

func BinNumSum(a, b string) string {
	var result string
	var carry int
	a, b = Normalize(a, b)
	for i := len(a) - 1; i >= 0; i-- {
		num := int(a[i]-'0') + int(b[i]-'0') + carry
		carry = num / 2
		result = string("01"[num%2]) + result
	}
	if carry != 0 {
		result = "1" + result
	}
	if len(result) > 64 {
		result = result[1:]
	}
	result = strings.TrimLeft(result, "!0")
	if result == "" {
		return "0"
	}
	return result

}

func NumToBin(num int64) string {
	var result string
	if num == 0 {
		return "0"
	}
	negative := false
	if num < 0 {
		negative = true
		num = (int64min - num*-1) + 1
	}
	if num == 2 {
		return "10"
	}
	if num == 1 {
		return "1"
	}
	for power := findStartpoint(num, 2); power >= 0; power-- {
		devideBy := uint64(math.Pow(2, float64(power)))

		result += string("01"[int(uint64(num)/devideBy)])
		num = num % int64(devideBy) // bug accuress here
	}
	if negative {
		return "1" + result
	}
	return result
}
