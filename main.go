package main

import (
	"fmt"
	"math"
	"strings"
)

const int64min = 9223372036854775807

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
	if num == 1 {
		return "1"
	}
	if num == 2 {
		return "10"
	}
	for power := findStartpoint(num, 2); power >= 0; power-- {
		devideBy := uint64(math.Pow(2, float64(power)))

		result = result + string("01"[int(uint64(num)/devideBy)])
		num = num % int64(devideBy)
	}
	if negative {
		return "1" + result
	}
	return result
}

func findStartpoint(num int64, base int) (power int) {
	for math.Pow(float64(base), float64(power)) <= float64(num) {
		power++
	}
	return power - 1
}

func main() {
	var num int64
	fmt.Scan(&num)
	fmt.Println(NumToBin(num))
}

func BinNumSum(a, b string) string {
	var result string
	var carry int

	if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	} else {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}

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
	result = strings.TrimLeft(result, "0")
	if result == "" {
		return "0"
	}
	return result

}
