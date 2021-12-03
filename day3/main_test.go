package main

import (
	"fmt"
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

func TestPart2(t *testing.T) {
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

	oxygen, co2, life, _ := part2(lines)

	fmt.Printf("Oxygen = %d\n", oxygen)
	fmt.Printf("CO2 = %d\n", co2)
	fmt.Printf("Life = %d\n", life)

	if oxygen != 23 {
		t.Fatalf("Expected oxygen level to be 23 but got %d", oxygen)
	}

	if co2 != 10 {
		t.Fatalf("Expected co2 level to be 10 but got %d", co2)
	}

	if life != 230 {
		t.Fatalf("Expected life support to be 230 but got %d", life)
	}
}
