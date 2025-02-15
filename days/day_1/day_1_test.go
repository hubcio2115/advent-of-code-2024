package day1

import "testing"

const INPUT = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	result := Part1(INPUT)

	if result != 11 {
		t.Errorf("Expected 11, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(INPUT)

	if result != 31 {
		t.Errorf("Expected 31, got %d", result)
	}
}
