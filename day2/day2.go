package main

import (
	"fmt"

	"github.com/advent-of-code-2024/lib"
)

func main() {
	input, err := lib.ParseInts("day2.in")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input [][]int) int {
	res := 0

	for _, report := range input {
		if isSafe(report) {
			res += 1
		}
	}

	return res
}

func part2(input [][]int) int {
	res := 0

	for _, report := range input {
		unsafeLevelIndex := findUnsafeLevelIndex(report)

		if unsafeLevelIndex == -1 || isSafeWithRemovedLevel(report, unsafeLevelIndex) {
			res += 1
		}
	}

	return res
}

// Returns true if the report is safe with exactly one element removed.
func isSafeWithRemovedLevel(report []int, index int) bool {
	// Slice with the first element removed (needed for edge cases) where
	// first element ruining the increase/decrease check in isSafe
	firstElementRemoved := report[1:]

	// Slice with element at index removed
	firstIndexRemoved := make([]int, len(report))
	copy(firstIndexRemoved, report)
	firstIndexRemoved = append(firstIndexRemoved[:index], report[index+1:]...)

	// Slice with element at index + 1 removed
	secondIndexRemoved := make([]int, len(report))
	copy(secondIndexRemoved, report)
	secondIndexRemoved = append(secondIndexRemoved[:index+1], report[index+2:]...)

	// If any of the above slice variants is safe, the report is safe with
	// just one element removed
	return isSafe(firstElementRemoved) || isSafe(firstIndexRemoved) || isSafe(secondIndexRemoved)
}

// Returns true if the report is safe without modifications.
func isSafe(report []int) bool {
	if len(report) == 1 {
		return true
	}

	// Check if report is increasing by evaluating first two terms
	isIncreasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check for valid differences
		if (isIncreasing && (diff <= 0 || diff > 3)) ||
			(!isIncreasing && (diff >= 0 || diff < -3)) {
			return false
		}
	}

	return true
}

// Finds the first index at which a level is unsafe (either moves in the
// opposite direction or increases/decreases too quickly).
func findUnsafeLevelIndex(report []int) int {
	// Check if report is increasing by evaluating first two terms
	isIncreasing := report[1] > report[0]

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		// Check for valid differences
		if (isIncreasing && (diff <= 0 || diff > 3)) ||
			(!isIncreasing && (diff >= 0 || diff < -3)) {
			return i
		}
	}

	return -1
}
