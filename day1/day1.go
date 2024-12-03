package main

import (
	"container/heap"
	"fmt"

	"github.com/advent-of-code-2024/lib"
)

func main() {
	input, err := lib.ParseInts("day1.in")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input [][]int) int {
	// Put input into two heaps.
	list1 := &lib.IntHeap{}
	list2 := &lib.IntHeap{}

	for _, line := range input {
		*list1 = append(*list1, line[0])
		*list2 = append(*list2, line[1])
	}

	heap.Init(list1)
	heap.Init(list2)

	res := 0
	for list1.Len() > 0 && list2.Len() > 0 {
		val1 := heap.Pop(list1).(int)
		val2 := heap.Pop(list2).(int)

		if val1 > val2 {
			res += val1 - val2
		} else {
			res += val2 - val1
		}
	}

	return res
}

func part2(input [][]int) int {
	freqMap := make(map[int]int)

	for _, line := range input {
		num := line[1]
		freqMap[num] += 1
	}

	var result int
	for _, line := range input {
		num := line[0]
		result += num * freqMap[num]
	}

	return result
}
