package main

import (
	"strings"
	"testing"
)

func TestPart1Example1(t *testing.T) {
	inputLines := strings.Split(example1, "\n")

	result := part1(inputLines)

	if result != 10 {
		t.Fatalf("Expected 10 but got %d", result)
	}
}

func TestPart1Example2(t *testing.T) {
	inputLines := strings.Split(example2, "\n")

	result := part1(inputLines)

	if result != 19 {
		t.Fatalf("Expected 19 but got %d", result)
	}
}

func TestPart1Example3(t *testing.T) {
	inputLines := strings.Split(example3, "\n")

	result := part1(inputLines)

	if result != 226 {
		t.Fatalf("Expected 226 but got %d", result)
	}
}

func TestPart2Example1(t *testing.T) {
	inputLines := strings.Split(example1, "\n")

	result := part2(inputLines)

	if result != 36 {
		t.Fatalf("Expected 36 but got %d", result)
	}
}

func TestPart2Example2(t *testing.T) {
	inputLines := strings.Split(example2, "\n")

	result := part2(inputLines)

	if result != 103 {
		t.Fatalf("Expected 103 but got %d", result)
	}
}
func TestPart2Example3(t *testing.T) {
	inputLines := strings.Split(example3, "\n")

	result := part2(inputLines)

	if result != 3509 {
		t.Fatalf("Expected 3509 but got %d", result)
	}
}
