package main

import (
	"fmt"
	"strings"
)

// func delzeros(res string) string {
// 	for len(res) > 64 {
// 		res = strings.TrimLeft(string(res), "10")
// 		if len(res) == 64 {
// 			res = strings.TrimPrefix(string(res), "0")
// 		}
// 	}
// 	return res
// }

// func delzeros(res string) string {
// 	if len(res) > 64 {
// 		res = strings.TrimLeft(res, "10")
// 		if len(res) == 64 {
// 			res = strings.TrimPrefix(res, "0")
// 		}
// 	}
// 	return res
// }

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
	// if len(string(res)) > 64 {
	// 	res = res[len(res)-64:]
	// 	res = strings.TrimPrefix(string(res), "0")
	// 	return res
	// }
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
}
