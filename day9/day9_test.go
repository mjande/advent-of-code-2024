package main

import "testing"

var day9TestInput = "2333133121414131402"

func TestDay9Part1(t *testing.T) {
	got := part1(day9TestInput)
	want := 1928

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestDay9Part2(t *testing.T) {
	got := part2(day9TestInput)
	want := 2858

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
