package day6

import "testing"

const INPUT = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1(t *testing.T) {
	result := Part1(INPUT)

	if result != 41 {
		t.Errorf("Expected 41, got %d", result)
	}
}
