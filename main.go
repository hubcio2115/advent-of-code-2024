package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	day4 "example.com/m/v2/days/day_4"
)

func main() {
	name := flag.String("input", "", "Path to the input file")
	flag.Parse()

	data, err := os.ReadFile(*name)

	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	content := string(data)

	result := day4.Part1(content)
	fmt.Printf("Part1: %d\n", result)

	result2 := day4.Part2(content)
	fmt.Printf("Part2: %d", result2)
}
