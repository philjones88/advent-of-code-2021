package main

import (
	"strings"
	"testing"
)

func TestParseIntoNumbersAndBoards(t *testing.T) {
	numbers, boards, _ := parseIntoNumbersAndBoards(strings.Split(example, "\n"))
	expectedNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	if len(numbers) != len(expectedNumbers) {
		t.Fatalf("Expected number of numbers to match but was %d vs %d", len(numbers), len(expectedNumbers))
	}

	if len(boards) != 3 {
		t.Fatalf("Expected 3 boards but got %d", len(boards))
	}
}

func TestPart1(t *testing.T) {
	numbers, boards, _ := parseIntoNumbersAndBoards(strings.Split(example, "\n"))
	winningNumbers, winningBoard := play(numbers, boards)
	score := calculateScore(winningNumbers, winningBoard)

	if score != 4512 {
		t.Fatalf("Expected result to be 4512 but got %d", score)
	}
}
