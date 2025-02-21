package day2

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	result := make([][]int, 0)

	for i, line := range lines {
		values := strings.Split(line, " ")
		row := make([]int, 0)

		for _, v := range values {
			intValue, err := strconv.Atoi(v)

			if err != nil {
				log.Fatalf("Couldn't parse: %s into an int on line: %d", v, i)
			}

			row = append(row, intValue)
		}

		result = append(result, row)
	}

	return result
}

func Part1(input string) int {
	rows := parseInput(input)
	result := 0

	for _, row := range rows {
		isSafe := true
		isIncreasing := false

		for i, current := range row {
			if !isSafe {
				break
			}

			if i == 0 {
				continue
			}

			prev := row[i-1]

			if i == 1 {
				isIncreasing = current > prev
			}

			isSafe = ((isIncreasing && current > prev) || (!isIncreasing && current < prev)) && diff(current, prev) <= 3
		}

		if isSafe {
			result += 1
		}
	}

	return result
}

// BUG: Passes the test but doesn't work on actual input
func Part2(input string) int {
	rows := parseInput(input)
	result := 0

	for _, row := range rows {
		isSafe := true
		isIncreasing := false
		failSafe := true

	loop:
		for i, current := range row {
			if !isSafe {
				break
			}

			if i == 0 {
				continue
			}

			prev := row[i-1]

			if i == 1 {
				isIncreasing = current > prev
			}

			isCurrentSafe := ((isIncreasing && current > prev) || (!isIncreasing && current < prev)) && diff(current, prev) <= 3

			if !isCurrentSafe && failSafe {
				failSafe = false
				row = slices.Delete(row, i, i+1)
				goto loop
			}

			isSafe = isCurrentSafe
		}

		if isSafe {
			result += 1
		}
	}

	return result
}
