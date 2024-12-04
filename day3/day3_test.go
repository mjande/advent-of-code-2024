package main

import "testing"

func TestDay3Part1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	got, err := part1(input)
	if err != nil {
		t.Error(err)
		return
	}

	want := 161

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestDay3Part2(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	got, err := part2(input)
	if err != nil {
		t.Error(err)
		return
	}

	want := 48

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
