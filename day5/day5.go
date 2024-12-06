package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/advent-of-code-2024/lib"
)

func main() {
	input, err := lib.ParseStrings("day5.in")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []string) int {
	rules, updates := getRulesAndUpdates(input)

	graph := createDependencyGraph(rules)

	res := 0
	for _, update := range updates {
		updateSlice := strings.Split(update, ",")

		if isInOrder(updateSlice, graph) {
			mid := len(updateSlice) / 2
			num, err := strconv.Atoi(updateSlice[mid])
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			res += num
		}
	}

	return res
}

func part2(input []string) int {
	rules, updates := getRulesAndUpdates(input)

	graph := createDependencyGraph(rules)

	res := 0
	for _, update := range updates {
		updateSlice := strings.Split(update, ",")

		if !isInOrder(updateSlice, graph) {
			reorder(updateSlice, graph)

			mid := len(updateSlice) / 2
			num, err := strconv.Atoi(updateSlice[mid])
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			res += num
		}
	}

	return res
}

// Takes the problem input and separates it into rules and updates,
// returning both as the result.
func getRulesAndUpdates(input []string) ([]string, []string) {
	i := 0
	for ; input[i] != ""; i++ {
	}

	rules := input[:i]
	updates := input[i+1:]

	return rules, updates
}

// Creates a dependency graph for all the rules. The graph is map where the
// key is any number that appears on the left side of a rule, and the value
// is a list of all number that must appear after it (if they appear at all).
func createDependencyGraph(rules []string) map[string][]string {
	graph := map[string][]string{}

	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		before, after := parts[0], parts[1]

		deps, exists := graph[before]
		if !exists {
			graph[before] = []string{after}
		} else {
			graph[before] = append(deps, after)
		}
	}

	return graph
}

// Checks if the given update is in order using a dependency graph.
// This function will check each subsequence value, checking to see if
// each value violates the ordering rules defined in the graph.
func isInOrder(update []string, graph map[string][]string) bool {
	seen := map[string]bool{}

	for _, num := range update {
		deps, exists := graph[num]
		if exists {
			for _, dep := range deps {
				if seen[dep] {
					return false
				}
			}
		}

		seen[num] = true
	}

	return true
}

// Creates a comparison function using a closure over the dependency
// graph, and then uses slices.SortFunc to sort the update.
func reorder(update []string, graph map[string][]string) {
	// Create sorting function
	cmp := func(a string, b string) int {
		deps, exists := graph[a]
		if exists && slices.Contains(deps, b) {
			return -1
		}
		deps, exists = graph[b]
		if exists && slices.Contains(deps, a) {
			return 1
		}

		return 0
	}

	slices.SortFunc(update, cmp)
}
