package main

import (
	"reflect"
	"testing"
)

func TestSlopeCheck(t *testing.T) {

	input := "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"

	slope := MakeSlope(input)

	got := slope.Check()
	want := 7

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

func TestMakeSlope(t *testing.T) {
	got := MakeSlope("..##.......\n#...#...#..")
	want := Slope{{'.', '.', '#', '#', '.', '.', '.', '.', '.', '.', '.'}, {'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.'}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %c want %c", got, want)
	}
}
