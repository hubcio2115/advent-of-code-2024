package main

import (
	"fmt"
	"log"
	"os"

	day1 "example.com/m/v2/days/day_1"
)

func main() {
	data, err := os.ReadFile("./days/day_1/part1_input.txt")

	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	content := string(data)

	result := day1.Part1(content)
	result2 := day1.Part2(content)

	fmt.Printf("Part1: %d\n", result)
	fmt.Printf("Part2: %d", result2)
}
