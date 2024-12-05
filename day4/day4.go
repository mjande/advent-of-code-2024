package main

import (
	"fmt"

	"github.com/advent-of-code-2024/lib"
)

const QUERY = "XMAS"

func main() {
	input, err := lib.ParseStrings("day4.in")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []string) int {
	res := 0

	for r, row := range input {
		for c, letter := range row {
			if letter == 'X' {
				res += checkDirections(input, r, c)
			}
		}
	}

	return res
}

func part2(input []string) int {
	res := 0

	for r, row := range input {
		for c, letter := range row {
			if letter == 'A' && checkX(input, r, c) {
				res += 1
			}
		}
	}

	return res
}

// Look for QUERY in all direction from starting position at [row, col].
// Return number of occurences of QUERY in any direction.
func checkDirections(wordSearch []string, row, col int) int {
	res := 0
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	for _, direction := range directions {
		if checkDirection(wordSearch, row, col, direction[0], direction[1]) {
			res += 1
		}
	}

	return res
}

// Checks if the query is present in the direction provided by the values
// in dx and dy from starting point at [row, col].
func checkDirection(wordSearch []string, row, col, dx, dy int) bool {
	index := 0
	rows, cols := len(wordSearch), len(wordSearch[0])

	for index < len(QUERY) {
		// If position is out of bounds or it does not match next character
		// in QUERY, return false
		if row < 0 || row >= rows || col < 0 || col >= cols ||
			wordSearch[row][col] != QUERY[index] {

			return false
		}

		// Check next position
		row += dx
		col += dy
		index += 1
	}

	return true
}

// Check in X (each diagonal) around starting point [row, col] for word MAS.
func checkX(wordSearch []string, row, col int) bool {
	// Check that positions in X around origin are all in bounds
	rows, cols := len(wordSearch), len(wordSearch[0])
	if row < 1 || row >= rows-1 || col < 1 || col >= cols-1 {
		return false
	}

	// Check that MAS occurs in both diagonals
	return checkDiag1ForMas(wordSearch, row, col) && checkDiag2ForMas(wordSearch, row, col)
}

// Check diagonal from top left to bottom right for the word MAS.
func checkDiag1ForMas(wordSearch []string, row, col int) bool {
	if wordSearch[row-1][col-1] == 'M' {
		return wordSearch[row+1][col+1] == 'S'
	} else if wordSearch[row-1][col-1] == 'S' {
		return wordSearch[row+1][col+1] == 'M'
	}

	return false
}

// Check diagonal from bottom left to top right for the word MAS.
func checkDiag2ForMas(wordSearch []string, row, col int) bool {
	if wordSearch[row+1][col-1] == 'M' {
		return wordSearch[row-1][col+1] == 'S'
	} else if wordSearch[row+1][col-1] == 'S' {
		return wordSearch[row-1][col+1] == 'M'
	}

	return false
}
