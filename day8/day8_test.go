package main

import "testing"

var input8 = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

func TestDay8Part1(t *testing.T) {
	got := part1(input8)
	want := 14

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestDay8Part2(t *testing.T) {
	got := part2(input8)
	want := 34

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
