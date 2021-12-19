package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := strings.Split(example, "\n")
	result := part1(input)

	if result != 4140 {
		t.Fatalf("Expected 4140 but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := strings.Split(example, "\n")
	result := part2(input)

	if result != 3993 {
		t.Fatalf("Expected 3993 but got %d", result)
	}
}
