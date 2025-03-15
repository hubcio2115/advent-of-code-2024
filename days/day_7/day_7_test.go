package day7

import "testing"

const INPUT = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestPart1(t *testing.T) {
	result := Main(INPUT, false)

	if result != 3749 {
		t.Errorf("Expected 3749, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := Main(INPUT, true)

	if result != 11387 {
		t.Errorf("Expected 11387, got %d", result)
	}
}
