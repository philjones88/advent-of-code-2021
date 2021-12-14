package main

import "testing"

func TestPart1(t *testing.T) {
	template, pairRules := processInput(example)

	result := calculate(template, pairRules, 10)

	if result != 1588 {
		t.Fatalf("Expected 1588 but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	template, pairRules := processInput(example)

	result := calculate(template, pairRules, 40)

	if result != 2188189693529 {
		t.Fatalf("Expected 2188189693529 but got %d", result)
	}
}
