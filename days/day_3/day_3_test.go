package day3

import "testing"

const part1Input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestPart1(t *testing.T) {
	result := Part1(part1Input)

	if result != 161 {
		t.Errorf("Expected 161, got %d", result)
	}
}

const part2Input = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart2(t *testing.T) {
	result := Part2(part2Input)

	if result != 48 {
		t.Errorf("Expected 48, got %d", result)
	}
}
