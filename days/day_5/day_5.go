package day5

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) int {
	split := strings.Split(input, "\n\n")
	rules := parseRules(split[0])
	queues := strings.Split(split[1], "\n")

	return calculateCorrectValues(rules, queues)
}

func parseRules(s string) map[string][]string {
	rules := map[string][]string{}

	for rule := range strings.SplitSeq(s, "\n") {
		split := strings.Split(rule, "|")

		if rules[split[0]] == nil {
			rules[split[0]] = make([]string, 0)
		}

		rules[split[0]] = append(rules[split[0]], split[1])
	}

	return rules
}

func calculateCorrectValues(rules map[string][]string, queues []string) int {
	result := 0

	for _, queue := range queues {
		pages := strings.Split(queue, ",")
		slices.Reverse(pages)

		isRight := true
		for i, page := range pages {
			if !isRight {
				break
			}

			rule := rules[page]
			rest := pages[i:]

			for _, page := range rest {
				if slices.Contains(rule, page) {
					isRight = false
				}
			}
		}

		if !isRight {
			continue
		}

		middleElement := pages[len(pages)/2]
		value, err := strconv.Atoi(middleElement)

		if err != nil {
			log.Fatalf("Couldn't parse %s into an int, during addition", middleElement)
		}

		result += value
	}

	return result
}

func Part2(input string) int {
	split := strings.Split(input, "\n\n")
	rules := parseRules(split[0])
	queues := strings.Split(split[1], "\n")

	wrongQueues := make([]string, 0)
	for _, queue := range queues {
		pages := strings.Split(queue, ",")
		slices.Reverse(pages)

		isRight := true
		for i, page := range pages {
			if !isRight {
				break
			}

			rule := rules[page]
			rest := pages[i:]

			for _, page := range rest {
				if slices.Contains(rule, page) {
					isRight = false
				}
			}
		}

		if isRight {
			continue
		}

		slices.Reverse(pages)
		slices.SortFunc(pages, func(a, b string) int {
			aRules := rules[a]
			bRules := rules[b]

			bContainedInARules := slices.Contains(aRules, b)
			aContainedInBRules := slices.Contains(bRules, a)

			switch true {
			case bContainedInARules && aContainedInBRules:
				return 0
			case aContainedInBRules:
				return 1
			default:
				return -1
			}
		})

		wrongQueues = append(wrongQueues, strings.Join(pages, ","))
	}

	return calculateCorrectValues(rules, wrongQueues)
}
