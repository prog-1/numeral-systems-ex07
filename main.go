package main

import (
	"fmt"
	"strings"
)

func NumToBin(num int64) string {
	var s string
	var l int64 = 1
	var count, zeros int
	SliceAlreadyExists := false
	var m []byte
	if num == 0 {
		s = "0"
		return s
	}
	NumberIsNegative := false
	if num < 0 {
		NumberIsNegative = true
		num = num * -1
	}
	for num > 0 {
		if l >= num && !SliceAlreadyExists {
			if l > num {
				count = count - 1
			}
			for zeros < count {
				m = append(m, 0)
				zeros++
			}
			count, zeros = 0, 0
			m = append(m, 1)
			if l == num {
				break
			} else {
				l = l / 2
			}
			num = num - l
			if num == 0 {
				break
			}
			l = 1
			SliceAlreadyExists = true
		}
		if l >= num && SliceAlreadyExists {
			if l > num {
				count = count - 1
			}
			m[count] = 1
			count, zeros = 0, 0
			if l == num {
				break
			} else {
				l = l / 2
			}
			num = num - l
			if num == 0 {
				break
			}
			l = 1
		}
		l = l * 2
		count++
	}
	if NumberIsNegative {
		First1Exists := false
		for i, b := range m {
			if b == 1 && !First1Exists {
				First1Exists = true
			} else if b == 0 && First1Exists {
				m[i] = 1
			} else if b == 1 && First1Exists {
				m[i] = 0
			}
		}
		for len(m) < 64 {
			m = append(m, 1)
		}
	}
	for i := len(m) - 1; i >= 0; i-- {
		s = s + fmt.Sprint(m[i])
	}
	return s
}

func AddZeros(s string, count int) string {
	for ; count > 0; count-- {
		s = "0" + s
	}
	return s
}

func MakeNormal(a, b string) (string, string) {
	if len(a) > len(b) {
		b = AddZeros(b, len(a)-len(b))
	} else {
		a = AddZeros(a, len(b)-len(a))
	}
	return a, b
}

func NumberExistsInBase2(num string) bool {
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, v := range num {
		if !strings.Contains(base36[:2], string(v)) {
			return false
		}
	}
	return true
}

func BinNumSum(a, b string) string {
	var s string
	var add uint8 // uint8 = byte
	a, b = MakeNormal(a, b)

	if NumberExistsInBase2(a) && NumberExistsInBase2(b) {
		for i := len(a) - 1; i >= 0; i-- {
			num := (a[i] - '0') + (b[i] - '0') + add // '0' = 48
			add = num / 2
			s = string((num%2)+'0') + s // '0' = 48
		}
		if add != 0 {
			s = "1" + s
		}
		if len(s) > 64 {
			s = s[len(s)-64:]
		}
		l := 0
		for l < len(s) { // to remove all unnecessary 0
			if s[l] == '0' {
				s = s[l:]
				l = 0
			} else {
				s = s[l:]
				break
			}
			l++
		}
	} else {
		s = "-1"
	}
	return s
}

func MainMenu() (choice int) {
	fmt.Println("Please choose your action:")
	fmt.Println("Enter 1 if you want to convert an integer to binary")
	fmt.Println("Enter 2 if you want to add two binary numbers")
	fmt.Println("Enter 3 if you want to exit")
	fmt.Scan(&choice)
	return
}

func main() {
	choice := MainMenu()
	for {
		if choice == 1 {
			var num int64
			fmt.Print("Enter number: ")
			fmt.Scan(&num)
			fmt.Println(NumToBin(num))
			break
		} else if choice == 2 {
			var a, b string
			fmt.Print("Enter a: ")
			fmt.Scan(&a)
			fmt.Print("Enter b: ")
			fmt.Scan(&b)
			if BinNumSum(a, b) != "-1" {
				fmt.Println(BinNumSum(a, b))
			} else {
				fmt.Println("ERROR: At least one number is not in base 2")
			}
			break
		} else if choice == 3 {
			fmt.Println("Program exited")
			break
		} else {
			fmt.Println("ERROR: Please enter 1, 2 or 3")
			fmt.Scan(&choice)
		}
	}
}
