package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	 gamma, epsilon, power, _ := part1(lines)

	 if gamma != 22 {
		 t.Fatalf("Expected gamma to be 22 but got %d", gamma)
	 }

	 if epsilon != 9 {
		 t.Fatalf("Expected epsilon to be 9 but got %d", epsilon)
	 }

	 if power != 198 {
		 t.Fatalf("Expected power to be 198 but got %d", power)
	 }
}
