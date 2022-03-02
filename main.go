package main

import "fmt"

func decimalToBinary(number int64) (s string) {
	fmt.Println(number)
	if number == 0 {
		return "0"
	}
	for number != 0 {
		s = fmt.Sprint(number%2, s)
		number /= 2
	}
	return s
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
func NumToBin(num int64) (s string) {
	switch {
	case num >= 0:
		s = string(decimalToBinary(num))
	default:
		var st string
		st = LongAdd("1111111111111111111111111111111111111111111111111111111111111111", decimalToBinary(num*(-1)))
		// if num != -1 {
		// 	st = LongAdd(st, "1")
		// }
		fmt.Println(st)
		if len(st) > 64 {
			for i := 64; i > 0; i-- {
				if st[i] == '0' {
					s = "1" + s
				} else {
					s = "0" + s
				}
			}
		}

	}
	return s
}

func main() {
	fmt.Println(NumToBin(-512))
}
