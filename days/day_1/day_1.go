package day1

import (
	"cmp"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

const SEPARATOR = "   "

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		values := strings.Split(line, SEPARATOR)

		leftValue, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatalf("Couldn't convert value to an int: %s", values[0])
		}

		rightValue, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatalf("Couldn't convert value to an int: %s", values[1])
		}

		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	slices.SortFunc(left, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	slices.SortFunc(right, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	result := 0
	for i, leftValue := range left {
		result += int(math.Abs(float64(leftValue - right[i])))
	}

	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		values := strings.Split(line, SEPARATOR)

		leftValue, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatalf("Couldn't convert value to an int: %s", values[0])
		}

		rightValue, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatalf("Couldn't convert value to an int: %s", values[1])
		}

		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	result := 0
	for _, leftValue := range left {
		occurences := 0

		for _, rightValue := range right {
			if leftValue == rightValue {
				occurences += 1
			}
		}

		result += leftValue * occurences
	}

	return result
}
