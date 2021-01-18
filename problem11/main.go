package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	grid = `08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48`

	gridRowLen = 20
	adjacent   = 4
)

type Direction int

const (
	Down Direction = iota
	Up
	Left
	Right
	DiagonalUpLeft
	DiagonalUpRight
	DiagonalDownLeft
	DiagonalDownRight
)

func calculate(direction Direction, row, col, adjacent int, grid *[][]int) int {
	switch direction {
	case Right:
		p := 1
		for c := col; c < (col + adjacent); c++ {
			p *= (*grid)[row][c]
		}
		return p
	case Left:
		p := 1
		for c := col; c > (col - adjacent); c-- {
			p *= (*grid)[row][c]
		}
		return p
	case Up:
		p := 1
		for r := row; r > (row - adjacent); r-- {
			p *= (*grid)[r][col]
		}
		return p
	case Down:
		p := 1
		for r := row; r < (row + adjacent); r++ {
			p *= (*grid)[r][col]
		}
		return p
	case DiagonalUpLeft:
		p := 1
		c := col
		for r := row; (r > (row - adjacent)) && (c > (c - adjacent)); {
			p *= (*grid)[r][c]
			r--
			c--
		}
		return p
	case DiagonalUpRight:
		p := 1
		r := row
		c := col
		for (r > (row - adjacent)) && (c < (col + adjacent)) {
			p *= (*grid)[r][c]
			r--
			c++
		}
		return p
	case DiagonalDownLeft:
		p := 1
		r := row
		c := col
		for r < (row+adjacent) && (c > (col - adjacent)) {
			p *= (*grid)[r][c]
			r++
			c--
		}
		return p
	case DiagonalDownRight:
		p := 1
		r := row
		c := col
		for (r < (row + adjacent)) && (c < (col + adjacent)) {
			p *= (*grid)[r][c]
			r++
			c++
		}
		return p
	default:
		return 0
	}
}

func canMove(row, col, adjacent int, direction Direction, grid *[][]int) bool {
	switch direction {
	case Right:
		if ((col + adjacent) > len(*grid)) || (row >= len(*grid)) {
			return false
		}
		return true
	case Left:
		if (((col + 1) - adjacent) < 0) || (row >= len(*grid)) {
			return false
		}
		return true
	case Up:
		if ((col < 0) || (col >= len(*grid))) || (((row + 1) - adjacent) < 0) {
			return false
		}
		if (row + 1) > len(*grid) {
			return false
		}
		return true
	case Down:
		if ((col < 0) || (col >= len(*grid))) || ((row + adjacent) > len(*grid)) {
			return false
		}
		return true
	case DiagonalUpLeft:
		if ((col + 1) - adjacent) < 0 {
			return false
		}
		if ((row + 1) - adjacent) < 0 {
			return false
		}
		if (row + 1) > len(*grid) {
			return false
		}
		return true
	case DiagonalUpRight:
		if ((row + 1) - adjacent) < 0 {
			return false
		}
		if (col + adjacent) > len(*grid) {
			return false
		}
		if row >= len(*grid) {
			return false
		}
		return true
	case DiagonalDownLeft:
		if (col < 0) || (col >= len(*grid)) {
			return false
		}
		if (row + adjacent) > len(*grid) {
			return false
		}
		if ((col + 1) - adjacent) < 0 {
			return false
		}
		return true
	case DiagonalDownRight:
		if (col + adjacent) > len(*grid) {
			return false
		}
		if row >= len(*grid) {
			return false
		}
		if (col < 0) || (col >= len(*grid)) {
			return false
		}
		if (row + adjacent) > len(*grid) {
			return false
		}
		return true
	default:
		return false
	}
}

func availableMoves(row, col int, grid *[][]int) []Direction {
	directions := []Direction{
		Down,
		Up,
		Left,
		Right,
		DiagonalUpLeft,
		DiagonalUpRight,
		DiagonalDownLeft,
		DiagonalDownRight,
	}

	available := make([]Direction, 0)

	for _, direction := range directions {
		if canMove(row, col, adjacent, direction, grid) {
			available = append(available, direction)
		}
	}

	return available
}

func gridTextToMatrix2(grid string, matrixSize int) [][]int {
	matrix := make([][]int, matrixSize)

	for i := 0; i < matrixSize; i++ {
		matrix[i] = make([]int, matrixSize)
	}

	textRows := strings.Split(grid, "\n")
	for i := 0; i < len(textRows); i++ {
		row := textRows[i]
		cols := strings.Split(row, " ")
		for c := 0; c < len(cols); c++ {
			x, _ := strconv.ParseInt(cols[c], 10, 64)
			matrix[i][c] = int(x)
		}
	}

	return matrix
}

func main() {
	matrix := gridTextToMatrix2(grid, gridRowLen)
	max := 0
	for row := 0; row < gridRowLen; row++ {
		for col := 0; col < gridRowLen; col++ {
			availableMoves := availableMoves(row, col, &matrix)
			for _, availableMove := range availableMoves {
				product := calculate(availableMove, row, col, adjacent, &matrix)
				if product > max {
					max = product
				}
			}
		}
	}

	fmt.Printf("Max product is: %d\n", max)
}
