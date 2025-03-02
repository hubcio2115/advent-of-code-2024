package day3

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) int {
	matches := parseInput(input)

	return calculateResult(matches)
}

func Part2(input string) int {
	cropped := croppInput(input)
	matches := parseInput(cropped)

	return calculateResult(matches)
}

func parseInput(input string) []string {
	re, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	return re.FindAllString(input, -1)
}

func croppInput(input string) string {
	donts := strings.Split(input, "don't()")
	doInput := donts[0]
	donts = donts[1:]

	for _, dont := range donts {
		index := strings.Index(dont, "do()")

		if index != -1 {
			doInput += dont[index:]
		}
	}

	return doInput
}

func calculateResult(matches []string) int {
	result := 0

	for _, match := range matches {
		start := strings.Index(match, "(")
		end := strings.Index(match, ")")
		commaIndex := strings.Index(match, ",")

		first, err := strconv.Atoi(match[start+1 : commaIndex])
		if err != nil {
			log.Fatalf("Couldn't parse %d, Cause: %s", first, err.Error())
		}

		second, err := strconv.Atoi(match[commaIndex+1 : end])
		if err != nil {
			log.Fatalf("Couldn't parse %d, Cause: %s", second, err.Error())
		}

		result += first * second
	}

	return result
}
