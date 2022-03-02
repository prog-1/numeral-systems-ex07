package main

import (
	"fmt"
	"math"
	"strings"
)

const int64min = 9223372036854775807

func findStartpoint(num int64, base int) (power int) {
	for math.Pow(float64(base), float64(power)) <= float64(num) {
		power++
	}
	return power - 1
}

func NumToBin(num int64) string {
	if num == 0 {
		return "0"
	}
	negative := false
	if num < 0 {
		negative = true
		num = (int64min - num*-1) + 1
	}
	var result string
	for power := findStartpoint(num, 2); power >= 0; power-- {
		devideBy := uint64(math.Pow(2, float64(power)))
		result += string("01"[int(uint64(num)/devideBy)])
		num %= int64(devideBy)
	}
	if negative {
		return "1" + result
	}
	return result
}

func LongAdd(a, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	var carry byte
	res := make([]byte, len(b)+1)
	for i := range a {
		sum := carry + a[len(a)-1-i] - '0' + b[len(b)-1-i] - '0'
		res[i], carry = sum%2+'0', sum/2
	}
	for i := len(a); i < len(b); i++ {
		sum := carry + b[len(b)-1-i] - '0'
		res[i], carry = sum%2+'0', sum/2
	}
	if carry > 0 {
		res[len(res)-1] = carry + '0'
	} else {
		res = res[:len(res)-1]
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func BinNumSum(a, b string) string {
	result := LongAdd(a, b)
	if len(result) > 64 {
		result = result[1:]
	}
	result = strings.TrimLeft(result, "0")
	if len(result) == 0 {
		return "0"
	}
	return result
}

func main() {
	var a, b int64
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Println(NumToBin(a))
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	fmt.Println("The addition of binary numbers: ", BinNumSum(NumToBin(a), NumToBin(b)))
}
