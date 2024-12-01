package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
