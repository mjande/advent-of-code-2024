package main

import (
	"fmt"
	"slices"

	"github.com/advent-of-code-2024/lib"
)

type direction struct {
	r, c int
}

type coord struct {
	r, c int
}

func main() {
	input1, err := lib.ParseBytes("day6.in")
	if err != nil {
		fmt.Println(err)
	}

	input2 := copyRoom(input1)
	fmt.Printf("Part 1: %v\n", part1(input1))
	fmt.Printf("Part 2: %v\n", part2(input2))
}

func part1(input [][]byte) int {
	r, c := findStart(input)
	dir := calculateDirection(input[r][c])

	return traverse(input, r, c, dir)
}

func part2(input [][]byte) int {
	r, c := findStart(input)
	dir := calculateDirection(input[r][c])

	candidates := findPotentialObstacles(input, r, c, dir)

	loops := 0
	for _, candidate := range candidates {
		room := copyRoom(input)

		row, col := candidate.r, candidate.c
		temp := room[row][col]

		if room[row][col] != '#' && checkForLoop(room, r, c, dir, row, col) {
			loops += 1
			fmt.Printf("[%v, %v]\n", row, col)
		}

		room[row][col] = temp
	}

	return loops
}

func findStart(room [][]byte) (int, int) {
	for r, row := range room {
		for c, cell := range row {
			if slices.Contains([]byte{'<', '>', '^', 'V'}, cell) {
				return r, c
			}
		}
	}

	return -1, -1
}

func calculateDirection(start byte) direction {
	switch start {
	case '>':
		return direction{0, 1}
	case '<':
		return direction{0, -1}
	case '^':
		return direction{-1, 0}
	case 'V':
		return direction{1, 0}
	default:
		return direction{0, 0}
	}
}

func traverse(room [][]byte, row, col int, dir direction) int {
	rows, cols := len(room), len(room[0])
	positions := 1
	room[row][col] = 'X'

	for row >= 0 && row < rows && col >= 0 && col < cols {
		if room[row][col] == '#' {
			row -= dir.r
			col -= dir.c
			dir = turn(dir)
			continue
		}

		if room[row][col] == '.' {
			positions += 1
			room[row][col] = 'X'
		}

		row += dir.r
		col += dir.c
	}

	return positions
}

func turn(dir direction) direction {
	if dir.r == -1 {
		return direction{0, 1}
	}
	if dir.c == 1 {
		return direction{1, 0}
	}
	if dir.r == 1 {
		return direction{0, -1}
	}
	if dir.c == -1 {
		return direction{-1, 0}
	}

	return direction{0, 0}
}

func copyRoom(original [][]byte) [][]byte {
	rows, cols := len(original), len(original[0])

	copiedRoom := make([][]byte, rows)

	for i, row := range original {
		copiedRoom[i] = make([]byte, cols)
		copy(copiedRoom[i], row)
	}

	return copiedRoom
}

func checkForLoop(room [][]byte, row, col int, dir direction, obRow, obCol int) bool {
	rows, cols := len(room), len(room[0])
	room[row][col] = 'X'
	newPath := false
	counter := 0

	for row >= 0 && row < rows && col >= 0 && col < cols {
		if row == obRow && col == obCol {
			newPath = true
			room[row][col] = '#'
		}

		if room[row][col] == '#' {
			row -= dir.r
			col -= dir.c
			dir = turn(dir)
			row += dir.r
			col += dir.c
			continue
		}

		if room[row][col] == 'X' && newPath {
			counter += 1
		}

		if room[row][col] == '.' {
			room[row][col] = 'X'
		}

		if counter >= max(rows, cols) {
			return true
		}

		row += dir.r
		col += dir.c
	}

	return false
}

func findPotentialObstacles(room [][]byte, row, col int, dir direction) []coord {
	rows, cols := len(room), len(room[0])
	candidates := map[coord]bool{}

	row += dir.r
	col += dir.c

	for row >= 0 && row < rows && col >= 0 && col < cols {
		if room[row][col] == '#' {
			row -= dir.r
			col -= dir.c
			dir = turn(dir)
		} else {
			candidates[coord{row, col}] = true
		}

		row += dir.r
		col += dir.c
	}

	candidatesSlice := make([]coord, len(candidates))
	for k := range candidates {
		candidatesSlice = append(candidatesSlice, k)
	}

	return candidatesSlice
}

/*
func findLoops(room [][]byte, row, col int, dir direction) int {
	backFillTraversal(room, row, col, dir)
	positions := map[coord]bool{}

	row += dir.r
	col += dir.c

	for isValid(row, col, len(room), len(room[0])) {
		for room[row][col] == '#' {
			row -= dir.r
			col -= dir.c
			dir = turn(dir)
			backFillTraversal(room, row, col, dir)
			row += dir.r
			col += dir.c
			continue
		}

		if slices.Contains(dir_symbols, room[row][col]) && resultsInLoop(room, row, col, dir) {
			positions[coord{row + dir.r, col + dir.c}] = true
			// fmt.Printf("Loop location found: [%v, %v]\n", row+dir.r, col+dir.c)
		}

		room[row][col] = dir.symbol
		// lib.PrintGrid(room)

		row += dir.r
		col += dir.c
	}

	return len(positions)
}

func resultsInLoop(room [][]byte, row, col int, dir direction) bool {
	newDir := turn(dir)
	return room[row][col] == newDir.symbol
}

var dir_symbols = []byte{'^', '>', '<', 'V'}

func backFillTraversal(room [][]byte, row, col int, dir direction) {
	for isValid(row, col, len(room), len(room[0])) && room[row][col] != '#' {
		room[row][col] = dir.symbol
		row -= dir.r
		col -= dir.c
	}
}

func isValid(row, col, rows, cols int) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols
}

*/
