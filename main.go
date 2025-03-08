package main

import (
	day6 "advent-of-code-2024/days/day_6"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	name := flag.String("input", "", "Path to the input file")
	flag.Parse()

	data, err := os.ReadFile(*name)

	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	content := string(data)

	result := day6.Part1(content)
	fmt.Printf("Part1: %d\n", result)

	// result2 := day5.Part2(content)
	// fmt.Printf("Part2: %d", result2)
}
