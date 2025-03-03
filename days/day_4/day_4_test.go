package day4

import "testing"

const testInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1(t *testing.T) {
	result := Part1(testInput)

	if result != 18 {
		t.Errorf("Expected 18, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(testInput)

	if result != 9 {
		t.Errorf("Expected 9, got %d", result)
	}
}
