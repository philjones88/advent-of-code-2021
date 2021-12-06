package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	
	numberStrings := strings.Split(example, ",")
	var numbers []int
	for _, numberString := range numberStrings {
		parsed, _ := strconv.Atoi(numberString)
		numbers = append(numbers, parsed)
	}
	fish := computeStartingFishStages(numbers)
	resultFish := part1(fish)
	result := countFish(resultFish)

	if result != 5934 {
		t.Fatalf("Expected 5934 but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	
	numberStrings := strings.Split(example, ",")
	var numbers []int
	for _, numberString := range numberStrings {
		parsed, _ := strconv.Atoi(numberString)
		numbers = append(numbers, parsed)
	}
	fish := computeStartingFishStages(numbers)
	resultFish := part2(fish)
	result := countFish(resultFish)

	if result != 26984457539 {
		t.Fatalf("Expected 26984457539 but got %d", result)
	}
}