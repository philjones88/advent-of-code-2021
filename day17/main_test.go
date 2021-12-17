package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := strings.TrimSpace(example)

	result := part1(input)

	if result != 45 {
		t.Fatalf("Expected 45 but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := strings.TrimSpace(example)

	result := part2(input)

	if result != 112 {
		t.Fatalf("Expected 112 but got %d", result)
	}
}
