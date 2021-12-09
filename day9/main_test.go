package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	inputLines := strings.Split(example, "\n")

	heightMap := convertToIntLines(inputLines)

	part1Result := part1(heightMap)

	if part1Result != 15 {
		t.Fatalf("Expected 15 but got %d", part1Result)
	}
}
