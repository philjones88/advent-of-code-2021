package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := strings.Split(example, "\n")

	players, positions := parsePlayersAndStartingPositions(input)

	result := part1(players, positions)

	if result != 739785 {
		t.Fatalf("Expected 739785 but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := strings.Split(example, "\n")
	_, positions := parsePlayersAndStartingPositions(input)
	cache = make(map[[5]int][2]int64)
	result := part2(positions)

	if result != 444356092776315 {
		t.Fatalf("Expected 444356092776315 but got %d", result)
	}
}
