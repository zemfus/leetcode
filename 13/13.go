package main

import "fmt"

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var sum, buf, pre int
	for i := len(s) - 1; i >= 0; i-- {
		buf = roman[string(s[i])]
		if buf >= pre {
			sum += buf
		} else {
			sum -= buf
		}
		pre = buf
	}
	return sum
}

func main() {
	var str string

	for _, err := fmt.Scan(&str); err == nil; _, err = fmt.Scan(&str) {
		fmt.Println(romanToInt(str))
	}
}
