package main

import "testing"

var input = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestDay2Part1(t *testing.T) {
	got := part1(input)
	want := 2

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestDay2Part2(t *testing.T) {
	got := part2(input)
	want := 4

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestEdgeCase1(t *testing.T) {
	edgeCase := [][]int{
		{54, 50, 48, 47, 46},
	}

	got := part2(edgeCase)
	want := 1

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestEdgeCase2(t *testing.T) {
	edgeCase := [][]int{
		{48, 46, 47, 49, 51, 54, 56},
	}

	got := part2(edgeCase)
	want := 1

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
