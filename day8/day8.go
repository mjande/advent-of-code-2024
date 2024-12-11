package main

import (
	"fmt"

	"github.com/advent-of-code-2024/lib"
)

type coord struct {
	r, c int
}

func main() {
	input, err := lib.ParseStrings("day8.in")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []string) int {
	antennaMap := mapAntennas(input)
	locations := map[coord]bool{}

	for _, coords := range antennaMap {
		findAntinodes(coords, len(input), len(input[0]), locations)
	}

	return len(locations)
}

func part2(input []string) int {
	antennaMap := mapAntennas(input)
	locations := map[coord]bool{}

	for _, coords := range antennaMap {
		findAntinodesWithResonance(coords, len(input), len(input[0]), locations)
	}

	return len(locations)
}

// Take a slice of strings representing the entire city map and
// return a map whose keys represent each type of antennae and whose values
// are a list of the coordinates for every antenna of that type.
func mapAntennas(cityMap []string) map[rune][]coord {
	antennaMap := map[rune][]coord{}

	for r, row := range cityMap {
		for c, cell := range row {
			if lib.IsAlphanumeric(cell) {
				_, ok := antennaMap[cell]
				if !ok {
					antennaMap[cell] = []coord{}
				}
				antennaMap[cell] = append(antennaMap[cell], coord{r, c})
			}
		}
	}

	return antennaMap
}

// Calculate the distance between each pair of antennas in the coordinate list,
// and use that distance to find the two antinodes associated with each pair.
// Then, add every antinode to the locations set.
func findAntinodes(coords []coord, rows, cols int, locations map[coord]bool) {
	if len(coords) == 1 {
		return
	}

	for i, antenna1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			antenna2 := coords[j]

			// Calculate distance from antenna1 to antenna2
			dr := antenna2.r - antenna1.r
			dc := antenna2.c - antenna1.c

			// Calculate position of both antinodes
			antinode1 := coord{antenna1.r - dr, antenna1.c - dc}
			antinode2 := coord{antenna2.r + dr, antenna2.c + dc}

			// Add each antinode to locations set if it is in bounds
			if isValidLocation(antinode1, rows, cols) {
				locations[antinode1] = true
			}

			if isValidLocation(antinode2, rows, cols) {
				locations[antinode2] = true
			}
		}
	}
}

// Returns true if the given coordinate in within bounds as defined by the
// given max rows and columns.
func isValidLocation(coordinate coord, rows, cols int) bool {
	return coordinate.r >= 0 && coordinate.r < rows && coordinate.c >= 0 && coordinate.c < cols
}

// Calculate the distance between each pair of antennas in the coordinate list,
// and use that distance to find the all antinodes associated with each pair.
// Then, add every antinode to the locations set.
func findAntinodesWithResonance(coords []coord, rows, cols int, locations map[coord]bool) {
	for i, antenna1 := range coords {
		locations[antenna1] = true
		for j := i + 1; j < len(coords); j++ {
			antenna2 := coords[j]

			// Calculate distance from antenna1 to antenna2
			dr := antenna2.r - antenna1.r
			dc := antenna2.c - antenna1.c

			// Keep adding antinodes out from antenna1 until out of bounds
			antinode1 := coord{antenna1.r - dr, antenna1.c - dc}
			for isValidLocation(antinode1, rows, cols) {
				locations[antinode1] = true
				antinode1 = coord{antinode1.r - dr, antinode1.c - dc}
			}

			// Keep adding antinodes out from antenna2 until out of bounds
			antinode2 := coord{antenna2.r + dr, antenna2.c + dc}
			for isValidLocation(antinode2, rows, cols) {
				locations[antinode2] = true
				antinode2 = coord{antinode2.r + dr, antinode2.c + dc}
			}
		}
	}
}
