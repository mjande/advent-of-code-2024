package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/advent-of-code-2024/lib"
)

func main() {
	input, err := lib.ParseString("day3.in")
	if err != nil {
		fmt.Println(err)
	}

	part1Ans, err := part1(input)
	if err != nil {
		fmt.Printf("Part 1 Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1: %v\n", part1Ans)

	part2Ans, err := part2(input)
	if err != nil {
		fmt.Printf("Part 2 Error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2: %v\n", part2Ans)
}

func part1(input string) (int, error) {
	res, err := evalEnabledMemory(input)
	if err != nil {
		return 0, err
	}

	return res, nil
}

// Find every substring representing a mul() operation in this
// section of memory, and return the sum of the results of all mul() calls.
func evalEnabledMemory(memory string) (int, error) {
	// Compile regex for finding mul() calls
	pattern, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		return 0, err
	}

	// For every matching substring, find the product and add to result
	matches := pattern.FindAllStringSubmatch(memory, -1)
	res := 0
	for _, match := range matches {
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}

		num2, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}

		res += num1 * num2
	}

	return res, nil
}

func part2(input string) (int, error) {
	// Compile regex for finding enabled sections of the program.
	pattern, err := regexp.Compile(`(?s)(?:^|do\(\))(.*?)(?:don\'t\(\)|$)`)
	if err != nil {
		return 0, err
	}

	// For each enabled section, evaluate the sum of all mul calls
	// using evalEnabledMemory() and add to result.
	matches := pattern.FindAllStringSubmatch(input, -1)
	res := 0
	for _, match := range matches {
		sum, err := evalEnabledMemory(match[1])
		if err != nil {
			return 0, err
		}

		res += sum
	}

	return res, nil
}
