package main

import "testing"

func TestPart1(t *testing.T) {
	part1Grid, part1Instructions := processInput(example)

	part1Result := part1(part1Grid, part1Instructions)

	if part1Result != 17 {
		t.Fatalf("Expected 17 but got %d", part1Result)
	}
}
