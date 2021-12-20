package main

import "testing"

func TestPart1(t *testing.T) {
	algorithm, image := parseAlgorithAndImage(example)

	result := part1(algorithm, image)

	if result != 35 {
		t.Fatalf("Expected 35 but got %d", result)
	}
}
func TestPart2(t *testing.T) {
	algorithm, image := parseAlgorithAndImage(example)

	result := part2(algorithm, image)

	if result != 3351 {
		t.Fatalf("Expected 3351 but got %d", result)
	}
}
