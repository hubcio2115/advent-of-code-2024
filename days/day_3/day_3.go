package day3

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const regexPatten = `mul\(\d{1,3},\d{1,3}\)`

func parseInput(input string) []string {
	re, _ := regexp.Compile(regexPatten)

	return re.FindAllString(input, -1)
}

func Part1(input string) int {
	matches := parseInput(input)

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

const doPattern = `do\(\)`
const dontPattern = `don\'t\(\)`

func parseInput2(input string) []string {
	re, _ := regexp.Compile(regexPatten)
	doRe, _ := regexp.Compile(doPattern)
	dontRe, _ := regexp.Compile(dontPattern)

	doMatches := doRe.FindAllStringIndex(input, -1)
	dontMatches := dontRe.FindAllStringIndex(input, -1)

	croppedInput := make([]string, 0)

	for i, dontIdx := range dontMatches {
		if i == 0 {
			croppedInput = append(croppedInput, input[:dontIdx[0]])
		} else {
			doIdx := doMatches[i-1]
			croppedInput = append(croppedInput, input[doIdx[0]:dontIdx[0]])
		}
	}

	if len(doMatches) == len(dontMatches) {
		croppedInput = append(croppedInput, input[doMatches[len(doMatches)-1][0]:])
	}

	return re.FindAllString(strings.Join(croppedInput, ""), -1)
}

func Part2(input string) int {
	matches := parseInput2(input)

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
