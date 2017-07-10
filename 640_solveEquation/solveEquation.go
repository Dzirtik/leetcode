package solveEquation

import (
	"strconv"
	"unicode"
)

func solveEquation(equation string) string {
	var x, nonX int
	var afterEq bool
	var op rune
	digit := -1
	for _, c := range equation + "=" {
		if unicode.IsDigit(c) {
			if digit > 0 {
				digit = digit*10 + int(c-'0')
			} else {
				digit = int(c - '0')
			}
		} else {
			if c == 'x' {
				if digit < 0 {
					digit = 1
				}

				if op == '-' && !afterEq || op == '+' && afterEq {
					x -= digit
				} else {
					x += digit
				}
			} else {
				if digit > 0 {
					if op == '-' && !afterEq || op == '+' && afterEq {
						nonX += digit
					} else {
						nonX -= digit
					}
				}
				if c == '=' {
					afterEq = true
					op = '+'
				} else if c == '+' || c == '-' {
					op = c
				}
			}
			digit = -1
		}
	}

	if x == 0 {
		if nonX == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}

	return "x=" + strconv.Itoa(int(nonX/x))
}
