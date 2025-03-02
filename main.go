package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	day3 "example.com/m/v2/days/day_3"
)

func main() {
	name := flag.String("input", "", "Path to the input file")
	flag.Parse()

	data, err := os.ReadFile(*name)

	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	content := string(data)

	result := day3.Part1(content)
	result2 := day3.Part2(content)

	fmt.Printf("Part1: %d\n", result)
	fmt.Printf("Part2: %d", result2)
}
