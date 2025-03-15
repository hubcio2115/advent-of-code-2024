package day7

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func Main(input string, canConcat bool) int {
	var result atomic.Int64
	testCases := parseInput(input)
	wg := sync.WaitGroup{}

	for _, testCase := range testCases {
		calibrationResult := testCase[0]
		calibrationValues := testCase[1:]

		wg.Add(1)
		go func() {
			if canMatch(calibrationResult, 0, calibrationValues, canConcat) {
				result.Add(int64(calibrationResult))
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return int(result.Load())
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
		for value := range strings.SplitSeq(split[1], " ") {
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
	result, _ := strconv.Atoi(fmt.Sprint(a, b))
	return result
}
