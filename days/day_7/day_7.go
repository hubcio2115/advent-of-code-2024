package day7

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Main(input string, canConcat bool) int {
	result := 0
	testCases := parseInput(input)
	for _, testCase := range testCases {
		calibrationResult := testCase[0]
		calibrationValues := testCase[1:]

		if canMatch(calibrationResult, 0, calibrationValues, canConcat) {
			result += calibrationResult
		}
	}

	return result
}

func Part2(input string) int {
	result := 0
	testCases := parseInput(input)
	for _, testCase := range testCases {
		calibrationResult := testCase[0]
		calibrationValues := testCase[1:]

		if canMatch(calibrationResult, 0, calibrationValues, true) {
			result += calibrationResult
		}
	}

	return result
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")

	testCases := [][]int{}
	for _, line := range lines {
		split := strings.Split(line, ": ")

		testValue, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatalf("Couldn't parse %s", split[0])
		}

		testCase := []int{testValue}
		for _, value := range strings.Split(split[1], " ") {
			calibration, err := strconv.Atoi(value)

			if err != nil {
				log.Fatalf("Couldn't parse %s", value)
			}

			testCase = append(testCase, calibration)
		}

		testCases = append(testCases, testCase)
	}

	return testCases
}

// Try all possible combinations recursively.
func canMatch(target, current int, numbers []int, canConcat bool) bool {
	if len(numbers) == 0 {
		return current == target
	}

	next, remaining := numbers[0], numbers[1:]

	// Try addition
	if canMatch(target, current+next, remaining, canConcat) {
		return true
	}

	// Try multiplication
	if canMatch(target, current*next, remaining, canConcat) {
		return true
	}

	// For part 2, also try concatenation
	if canConcat {
		// Concatenate current and next (e.g. 1 and 2 become 12)
		concatVal := concat(current, next)
		if canMatch(target, concatVal, remaining, canConcat) {
			return true
		}
	}

	return false
}

func concat(a, b int) int {
	result, _ := strconv.Atoi(fmt.Sprint(a) + fmt.Sprint(b))
	return result
}
