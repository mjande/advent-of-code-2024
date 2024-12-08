package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/advent-of-code-2024/lib"
)

func main() {
	input, err := lib.ParseStrings("day7.in")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []string) int {
	res := 0

	for _, equation := range input {
		testValue, operands, err := destructureEquation(equation)
		if err != nil {
			fmt.Println(err)
			return -1
		}

		if isCalibrated(operands, testValue, false) {
			res += testValue
		}
	}

	return res
}

func part2(input []string) int {
	res := 0

	for _, equation := range input {
		testValue, operands, err := destructureEquation(equation)
		if err != nil {
			fmt.Println(err)
			return -1
		}

		if isCalibrated(operands, testValue, true) {
			res += testValue
		}
	}

	return res
}

// Take equation string and return test value as an integer and
// the operands as a list of integers.
func destructureEquation(equation string) (int, []int, error) {
	operands := []int{}
	parts := strings.Split(equation, ":")

	// Convert test value to integer
	testValue, err := strconv.Atoi(parts[0])
	if err != nil {
		return -1, operands, err
	}

	// Convert operands to list of integers
	parts[1] = strings.TrimSpace(parts[1])
	operandStrs := strings.Split(parts[1], " ")
	for _, operandStr := range operandStrs {
		num, err := strconv.Atoi(operandStr)
		if err != nil {
			return -1, operands, err
		}

		operands = append(operands, num)
	}

	return testValue, operands, nil
}

// Checks if all operands can be combined into target value using recursion.
func isCalibrated(operands []int, target int, concatOp bool) bool {
	// Base case: operands are combined into final value
	if len(operands) == 1 {
		return operands[0] == target
	}

	// Get result of adding or multiplying
	sum := operands[0] + operands[1]
	product := operands[0] * operands[1]

	// Create new operands lists with previous results
	operandsAfterAdd := append([]int{sum}, operands[2:]...)
	operandsAfterMul := append([]int{product}, operands[2:]...)

	// Use recursion to see if adding or multiplying first two operands
	// will yield target value
	res := isCalibrated(operandsAfterAdd, target, concatOp) || isCalibrated(operandsAfterMul, target, concatOp)

	// If concatOp is true, check to see if concat operation will yield
	// target value
	if concatOp {
		// Concatenate first two operands
		concatenated, err := concatNums(operands[0], operands[1])
		if err != nil {
			fmt.Println(err)
			return false
		}

		// Connect result of concat to res
		operandsAfterConcat := append([]int{concatenated}, operands[2:]...)
		res = res || isCalibrated(operandsAfterConcat, target, concatOp)
	}

	return res
}

// Takes two integers and concatenate them and return the resulting integer.
func concatNums(num1, num2 int) (int, error) {
	str := fmt.Sprintf("%v%v", num1, num2)
	return strconv.Atoi(str)
}
