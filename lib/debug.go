package lib

import "fmt"

func PrintGrid(grid [][]byte) {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			fmt.Printf("%c ", grid[row][column])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
