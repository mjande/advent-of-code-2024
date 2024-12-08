package main

import "testing"

var input = [][]byte{
	[]byte("....#....."),
	[]byte(".........#"),
	[]byte(".........."),
	[]byte("..#......."),
	[]byte(".......#.."),
	[]byte(".........."),
	[]byte(".#..^....."),
	[]byte("........#."),
	[]byte("#........."),
	[]byte("......#..."),
}

func TestDay6Part1(t *testing.T) {
	got := part1(input)
	want := 41

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestDay6Part2(t *testing.T) {
	got := part2(input)
	want := 6

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
