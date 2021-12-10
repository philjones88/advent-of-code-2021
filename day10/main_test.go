package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	inputLines := strings.Split(example, "\n")
	result, _ := part1(inputLines)

	if result != 26397 {
		t.Fatalf("Expected 26397 but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	inputLines := strings.Split(example, "\n")
	_, result := part1(inputLines)

	if result != 288957 {
		t.Fatalf("Expected 288957 but got %d", result)
	}
}
