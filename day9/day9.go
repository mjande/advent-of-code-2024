package main

import (
	"fmt"

	"github.com/advent-of-code-2024/lib"
)

func main() {
	input, err := lib.ParseString("day9.in")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	// fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input string) int {
	disk := simulateDisk(input)

	compressDisk(disk)

	return calculateChecksum(disk)
}

func part2(input string) int {
	disk := simulateDisk(input)

	compressDiskSafely(disk)

	return calculateChecksum(disk)
}

// Takes a disk map and converts it into an int slice where used space is
// represented by the id at that position and empty space by the number -1.
func simulateDisk(diskMap string) []int {
	disk := []int{}

	for i := 0; i < len(diskMap); i++ {
		numBlocks := int(diskMap[i] - '0')

		for j := 0; j < numBlocks; j++ {
			if i%2 == 0 {
				disk = append(disk, i/2)
			} else {
				disk = append(disk, -1)
			}
		}
	}

	return disk
}

// Takes an int slice representing a disk and compresses it by reading it from
// left to right and filling any free memory with used memory from the end of
// the disk.
func compressDisk(disk []int) {
	l := 0
	r := len(disk) - 1

	for disk[r] == -1 {
		r--
	}

	for l < r {
		if disk[l] == -1 {
			disk[l], disk[r] = disk[r], disk[l]

			for disk[r] == -1 {
				r--
			}
		}

		l++
	}
}

// Takes an int slice representing a disk and compresses it by reading it from
// left to right and filling any free memory with used memory from the end of
// the disk (without fragment any memory sections).
func compressDiskSafely(disk []int) {
	r := len(disk) - 1

	for r > 0 {
		if disk[r] != -1 {
			swapped := scanForEligibleSwaps(disk, r)

			if !swapped {
				id := disk[r]
				for disk[r] == id {
					r--
				}
			}
		}

		// Scan to next used memory space
		for r > 0 && disk[r] == -1 {
			r--
		}
	}
}

func scanForEligibleSwaps(disk []int, r int) bool {
	needed := calculateNeededSpace(disk, r)
	l := 0

	for l < r {
		if disk[l] != -1 {
			l++
			continue
		}

		free := calculateFreeSpace(disk, l)

		if free > needed {
			swapMemory(disk, l, r, needed)
			return true
		}

		l += free
	}

	return false
}

func calculateFreeSpace(disk []int, start int) int {
	free := 0

	for i := start; disk[i] == -1; i++ {
		free++
	}

	return free
}

func calculateNeededSpace(disk []int, end int) int {
	needed := 0
	id := disk[end]

	for i := end; disk[i] == id; i-- {
		needed++
	}

	return needed
}

func swapMemory(disk []int, dst, src, num int) {
	for ; num > 0; num-- {
		disk[src], disk[dst] = disk[dst], disk[src]

		src--
		dst++
	}
}

// Calculates a disk's checksum.
func calculateChecksum(disk []int) int {
	res := 0

	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			res += disk[i] * i
		}
	}

	return res
}
