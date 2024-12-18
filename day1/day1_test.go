package main

import "testing"

func TestDay1Part1(t *testing.T) {
	input := [][]int{
		{3, 4},
		{4, 3},
		{2, 5},
		{1, 3},
		{3, 9},
		{3, 3},
	}

	got := part1(input)
	want := 11

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestDay1Part2(t *testing.T) {
	input := [][]int{
		{3, 4},
		{4, 3},
		{2, 5},
		{1, 3},
		{3, 9},
		{3, 3},
	}

	got := part2(input)
	want := 31

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
