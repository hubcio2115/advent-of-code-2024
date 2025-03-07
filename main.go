package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	day5 "example.com/m/v2/days/day_5"
)

func main() {
	name := flag.String("input", "", "Path to the input file")
	flag.Parse()

	data, err := os.ReadFile(*name)

	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	content := string(data)

	result := day5.Part1(content)
	fmt.Printf("Part1: %d\n", result)

	result2 := day5.Part2(content)
	fmt.Printf("Part2: %d", result2)
}
