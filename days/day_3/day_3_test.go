package day3

import "testing"

const PART1_INPUT = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestPart1(t *testing.T) {
	result := Part1(PART1_INPUT)

	if result != 161 {
		t.Errorf("Expected 161, got %d", result)
	}
}

const PART2_INPUT = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart2(t *testing.T) {
	result := Part2(PART2_INPUT)

	if result != 48 {
		t.Errorf("Expected 48, got %d", result)
	}
}
