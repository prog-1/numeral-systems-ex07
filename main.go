package main

import (
	"fmt"
	"math"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NumToBin(num int64) (convert string) {
	if num < 0 {
		//will be done soon
	}
	if num == 0 {
		return "0"
	}
	var z int64
	var number float64
	for number <= float64(num) {
		z++
		number = math.Pow(2, float64(z))
	}
	number /= float64(2)
	for z > 0 {
		z--
		convert += string(base36[int(num/int64(number))])
		num %= int64(number)
		number /= 2
	}
	return convert
}

func BinNumSum(a, b string) string {
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
	re := string(res)
	if len(res) > 64 {
		re = re[len(re)-64:]
		if len(re) == 64 {
			re = strings.TrimLeft(re, "0")
		}
	}
	if len(re) == 0 {
		re = "0"
	}
	return re
}

func main() {
	for _, t := range []struct {
		a, b string
	}{
		{"1", "11"},
		{"11001", "1111111111111111111111111111111111111111111111111111111111110010"},
		{"1111111111111111111111111111111111111111111111111111111111110110", "111"},
		{"11", "1111111111111111111111111111111111111111111111111111111111111101"},
	} {
		fmt.Printf("%v + %v = %v\n", t.a, t.b, BinNumSum(t.a, t.b))
	}
	var a, b string
	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)
	fmt.Printf("%v * %v = %v\n", a, b, BinNumSum(a, b))
	var z int64
	fmt.Scan(&z)
	fmt.Printf("%v = %v", z, NumToBin(z))
}
