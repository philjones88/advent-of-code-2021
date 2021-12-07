package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	exampleSplit := strings.Split(example, ",")

	var crabPositions []int

	for _, example := range exampleSplit {
		crabPositions = append(crabPositions, stringToInt(example))
	}

	result := part1(crabPositions)

	if result != 37 {
		t.Fatalf("Expected 37 but got %d\n", result)
	}
}

func TestPart2(t *testing.T) {
	exampleSplit := strings.Split(example, ",")

	var crabPositions []int

	for _, example := range exampleSplit {
		crabPositions = append(crabPositions, stringToInt(example))
	}

	result := part2(crabPositions)

	if result != 168 {
		t.Fatalf("Expected 168 but got %d\n", result)
	}
}
