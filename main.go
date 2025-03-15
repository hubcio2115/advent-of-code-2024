package main

import (
	day7 "advent-of-code-2024/days/day_7"
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

	result := day7.Main(content, false)
	fmt.Printf("Part1: %d\n", result)

	result2 := day7.Main(content, true)
	fmt.Printf("Part2: %d", result2)
}
