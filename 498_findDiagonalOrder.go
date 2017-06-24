package main

import "fmt"

func findDiagonalOrder(matrix [][]int) []int {
	var x, y int
	var result []int
	var down bool

	if len(matrix) == 0 {
		return result
	}

	maxX, maxY := len(matrix[0])-1, len(matrix)-1

	for {
		result = append(result, matrix[y][x])

		if x == maxX && y == maxY {
			return result
		}

		if !down && x == maxX {
			y++
			down = down != true
		} else if down && y == maxY {
			x++
			down = down != true
		} else if down && x == 0 {
			y++
			down = down != true
		} else if !down && y == 0 {
			x++
			down = down != true
		} else {
			if down {
				x--
				y++
			} else {
				x++
				y--
			}
		}
	}
}

func main() {
	values := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	row3 := []int{7, 8, 9}

	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	fmt.Println(findDiagonalOrder(values))
}
