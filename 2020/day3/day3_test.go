package main

import "testing"

func TestSlopeCheck(t *testing.T) {

	input := "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"

	slope := MakeSlope(input)

	got := slope.Check()
	want := 7

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
