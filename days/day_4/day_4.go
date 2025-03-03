package day4

import (
	"regexp"
	"slices"
	"strings"
)

const SEARCH_WORD = "XMAS"

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	result := 0

	// Horizontally
	result += findMatchesHorizontally(lines)
	result += findMatchesHorizontallyReversed(lines)

	// Vertically
	columns := getColumns(lines)

	result += findMatchesHorizontally(columns)
	result += findMatchesHorizontallyReversed(columns)

	// Diagonally
	topLeftToBottomRight, topRightToBottomLeft := getDiagonals(lines)

	result += findMatchesHorizontally(topLeftToBottomRight)
	result += findMatchesHorizontally(topRightToBottomLeft)

	result += findMatchesHorizontallyReversed(topLeftToBottomRight)
	result += findMatchesHorizontallyReversed(topRightToBottomLeft)

	return result
}

func Part2(input string) int {
	panic("Unimplemented!")
}

func findMatchesHorizontally(lines []string) int {
	result := 0
	re, _ := regexp.Compile(SEARCH_WORD)

	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		result += len(matches)
	}

	return result
}

func getCharacterAtIndex(s string, i int) string {
	return string([]rune(s)[i])
}

func findMatchesHorizontallyReversed(lines []string) int {
	result := 0
	re, _ := regexp.Compile(SEARCH_WORD)

	for _, line := range lines {
		characters := strings.Split(line, "")
		slices.Reverse(characters)
		reversed := strings.Join(characters, "")

		matches := re.FindAllString(reversed, -1)
		result += len(matches)
	}

	return result
}

func getColumns(lines []string) []string {
	columns := make([]string, 0)
	colNumber := len(lines[0])
	rowNumber := len(lines)
	for i := range colNumber {
		column := ""

		for j := range rowNumber {
			column += getCharacterAtIndex(lines[j], i)
		}

		columns = append(columns, column)
	}

	return columns
}

func getDiagonals(lines []string) ([]string, []string) {
	n := len(lines)

	// ↘️ Top-left to bottom-right topLeftToBottomRight
	topLeftToBottomRight := make([]string, 0)

	for start := range n {
		// Start from first row
		diag1 := ""
		for i, j := 0, start; i < n && j < n; i, j = i+1, j+1 {
			diag1 += getCharacterAtIndex(lines[i], j)
		}

		topLeftToBottomRight = append(topLeftToBottomRight, diag1)

		// Start from first column (except first row, already taken)
		if start > 0 {
			diag2 := ""
			for i, j := start, 0; i < n && j < n; i, j = i+1, j+1 {
				diag2 += getCharacterAtIndex(lines[i], j)
			}

			topLeftToBottomRight = append(topLeftToBottomRight, diag2)
		}
	}

	// ↙️ Top-right to bottom-left diagonals
	topRightToBottomLeft := make([]string, 0)

	for start := range n {
		// Start from first row
		diag1 := ""
		for i, j := 0, start; i < n && j >= 0; i, j = i+1, j-1 {
			diag1 += getCharacterAtIndex(lines[i], j)
		}

		topRightToBottomLeft = append(topRightToBottomLeft, diag1)

		// Start from last column (except first row, already taken)
		if start > 0 {
			diag2 := ""
			for i, j := start, n-1; i < n && j >= 0; i, j = i+1, j-1 {
				diag2 += getCharacterAtIndex(lines[i], j)
			}

			topRightToBottomLeft = append(topRightToBottomLeft, diag2)
		}
	}

	return topLeftToBottomRight, topRightToBottomLeft
}
