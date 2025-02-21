package day2

import (
	"testing"
)

const INPUT = `1 3 2 4 5
1 2 7 8 9
9 7 6 2 1
7 6 4 2 1
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	result := Part1(INPUT)

	if result != 2 {
		t.Errorf("Expected 11, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(INPUT)

	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}
