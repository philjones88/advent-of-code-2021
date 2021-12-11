package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	lines := strings.Split(example, "\n")

	octopuses := linesToOctopusMap(lines)

	result := part1(octopuses)

	if result != 1656 {
		t.Fatalf("Expected 1656 but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(example, "\n")

	octopuses := linesToOctopusMap(lines)

	result := part2(octopuses)

	if result != 195 {
		t.Fatalf("Expected 195 but got %d", result)
	}
}
