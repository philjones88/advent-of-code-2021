package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	inputLines := strings.Split(example, "\n")
	grid := makeGrid(inputLines)
	result := findShortestPath(grid)
	if result != 40 {
		t.Fatalf("Expected 40 but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	inputLines := strings.Split(example, "\n")
	grid := makeGrid(inputLines)
	grid = expandGrid(grid)
	result := findShortestPath(grid)
	if result != 315 {
		t.Fatalf("Expected 315 but got %d", result)
	}
}
