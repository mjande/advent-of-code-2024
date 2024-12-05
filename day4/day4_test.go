package main

import "testing"

var input = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestDay4Part1(t *testing.T) {
	got := part1(input)
	want := 18

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestDay4Part2(t *testing.T) {
	got := part2(input)
	want := 9

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
