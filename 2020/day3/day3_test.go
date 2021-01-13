package main

import (
	"reflect"
	"testing"
)

func TestMakeSlope(t *testing.T) {
	got := MakeSlope("..##.......\n#...#...#..")
	want := Slope{{'.', '.', '#', '#', '.', '.', '.', '.', '.', '.', '.'}, {'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.'}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %c want %c", got, want)
	}
}

func TestSlopeCheck(t *testing.T) {

	input := "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"

	slope := MakeSlope(input)

	checkTests := []struct {
		right int
		down  int
		trees int
	}{
		{right: 1, down: 1, trees: 2},
		{right: 3, down: 1, trees: 7},
		{right: 5, down: 1, trees: 3},
		{right: 7, down: 1, trees: 4},
		{right: 1, down: 2, trees: 2},
	}

	for _, tt := range checkTests {
		t.Run("Slope check", func(t *testing.T) {
			got := slope.Check(tt.right, tt.down)
			if got != tt.trees {
				t.Errorf("right %d, down %d, got %d want %d", tt.right, tt.down, got, tt.trees)
			}
		})
	}

}
