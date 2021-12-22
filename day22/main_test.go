package main

import (
	"strings"
	"testing"
)

func TestPart1Example0(t *testing.T) {
	input := strings.Split(example0, "\n")

	steps := processInputToSteps(input)

	result := part1(steps)

	if result != 27 {
		t.Fatalf("Expected 27 but got %d", result)
	}
}

func TestPart1Example1(t *testing.T) {
	input := strings.Split(example1, "\n")

	steps := processInputToSteps(input)

	result := part1(steps)

	if result != 46 {
		t.Fatalf("Expected 46 but got %d", result)
	}
}

func TestPart1Example2(t *testing.T) {
	input := strings.Split(example2, "\n")

	steps := processInputToSteps(input)

	result := part1(steps)

	if result != 38 {
		t.Fatalf("Expected 38 but got %d", result)
	}
}

func TestPart1Example3(t *testing.T) {
	input := strings.Split(example3, "\n")

	steps := processInputToSteps(input)

	result := part1(steps)

	if result != 39 {
		t.Fatalf("Expected 39 but got %d", result)
	}
}

func TestPart1Example4(t *testing.T) {
	input := strings.Split(example4, "\n")

	steps := processInputToSteps(input)

	result := part1(steps)

	if result != 590784 {
		t.Fatalf("Expected 590784 but got %d", result)
	}
}

func TestPart2Example5(t *testing.T) {
	input := strings.Split(example5, "\n")

	steps := processInputToSteps(input)

	result := part2(&steps)

	if result != 2758514936282235 {
		t.Fatalf("Expected 2758514936282235 but got %d", result)
	}
}
