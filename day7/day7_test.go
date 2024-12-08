package main

import "testing"

var input = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

func TestDay7Part1(t *testing.T) {
	got := part1(input)
	want := 3749

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestDay7Part2(t *testing.T) {
	got := part2(input)
	want := 11387

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
