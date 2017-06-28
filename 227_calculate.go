package main

import (
	"fmt"
	"unicode"
)

var numbers []int
var ops []rune

func operate() {
	var op rune
	fmt.Println(numbers, ops)
	op, ops = ops[len(ops)-1], ops[0:len(ops)-1]

	var result int
	len := len(numbers)
	if op == '+' {
		result = numbers[len-2] + numbers[len-1]
	}
	if op == '-' {
		result = numbers[len-2] - numbers[len-1]
	}
	if op == '*' {
		result = numbers[len-2] * numbers[len-1]
	}
	if op == '/' {
		result = numbers[len-2] / numbers[len-1]
	}
	numbers = append(numbers[0:len-2], result)
}

func calculate(s string) int {
	numbers = []int{}
	ops = []rune{}
	var number int
	for _, c := range s {
		if unicode.IsDigit(c) {
			number = number*10 + int(c-'0')
		}
		if !unicode.IsDigit(c) && c != ' ' {
			numbers = append(numbers, number)
			number = 0
			if c == '+' || c == '-' {
				for len(ops) > 0 {
					operate()
				}
			}
			if len(ops) > 0 && (ops[len(ops)-1] == '*' || ops[len(ops)-1] == '/') {
				operate()
			}
			ops = append(ops, c)
		}
	}
	numbers = append(numbers, number)
	for len(ops) > 0 {
		operate()
	}
	return numbers[0]
}

func main() {
	fmt.Println(calculate("1+2*5/3+6/4*2"))
}
