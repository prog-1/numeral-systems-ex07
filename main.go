package main

import (
	"fmt"
	"math"
	"strings"
)

func NumToBin(num int64) (res string) {
	const dictionary = "0123456789abcdefghijklmnopqrstuvwxyz"
	if num == 0 {
		return "0"
	}
	if num > 0 {
		var i int64
		var number float64
		for number <= float64(num) {
			i++
			number = math.Pow(2, float64(i))
		}
		number /= float64(2)
		for i > 0 {
			i--
			res += string(dictionary[int(num/int64(number))])
			num %= int64(number)
			number /= 2
		}
		return res
	}
	if num < 0 {

	}
	return res
}
func addZeros(a string, i int) string {
	for ; i > 0; i-- {
		a = "0" + a
	}
	return a
}
func LongAdd(a, b string) (res string) {
	if len(a) > len(b) {
		a, b = b, a
	}
	var carry byte
	a = addZeros(a, len(b)-len(a))
	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + carry
		carry = num / 10
		res = string(num%10+'0') + res
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}

func BinNumSum(a, b string) string {
	res := LongAdd(a, b)
	if len(res) > 64 {
		res = res[1:]
	}
	res = strings.TrimLeft(res, "0")
	if len(res) == 0 {
		return "0"
	}
	return res
}

func main() {
	fmt.Println()
	var choice int
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Println("Enter number:")
		var num int64
		fmt.Scan(&num)
		fmt.Println("Number", num, "in binary is", NumToBin(num))
	}
	if choice == 2 {
		fmt.Println("Enter two integers:")
		var num1, num2 string
		fmt.Scan(&num1, &num2)
		fmt.Println("Sum of two integers is", BinNumSum(num1, num2))
	}
}
