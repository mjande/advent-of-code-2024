package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Open the input file and convert the input into a 3D slice of
// integers.
func ParseInts(filename string) ([][]int, error) {
	var result [][]int

	file, err := os.Open("inputs/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var nums []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("failed to parse number %q: %w", num, err)
			}

			nums = append(nums, num)
		}

		result = append(result, nums)
	}

	return result, nil
}

// Open the input file and return the input as a single string.
func ParseString(filename string) (string, error) {
	bytes, err := os.ReadFile("inputs/" + filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Open the input file and convert the input into a slice of strings.
func ParseStrings(filename string) ([]string, error) {
	var result []string

	file, err := os.Open("inputs/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

// Open the input file and convert the input into a 3D slice of bytes.
func ParseBytes(filename string) ([][]byte, error) {
	var result [][]byte

	file, err := os.Open("inputs/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)

		result = append(result, bytes)
	}

	return result, nil
}
