package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result := computePart1(strings.Split(example, "\n"))

	if result != 5 {
		t.Fatalf("Expected 5 but got %d", result)
	}
}
