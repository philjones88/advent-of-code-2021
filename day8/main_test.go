package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	exampleLines := strings.Split(example, "\n")
	var inputLines []string

	for _, inputLine := range exampleLines {
		inputLines = append(inputLines, strings.Split(inputLine, " | ")[1])
	}

	result := part1(inputLines)

	if result != 26 {
		t.Fatalf("Expected 26 but got %d\n", result)
	}
}
