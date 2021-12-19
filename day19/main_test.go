package main

import (
	"strings"
	"testing"
)

func TestRotationOfPoint1(t *testing.T) {
	point := Point{1, 1, 1}

	result := point.Rotate(0)

	if point != result {
		t.Fatalf("Expected rotation 0 to return same x,y,z")
	}
}
func TestRotationOfPoint2(t *testing.T) {
	point := Point{1, 1, 1}
	expected := Point{-1, -1, -1}

	result := point.Rotate(23)

	if result != expected {
		t.Fatalf("Expected rotation 0 to return same x,y,z")
	}
}

func TestPart1(t *testing.T) {
	scanners := computeScanners(strings.Split(example, "\n\n"))
	beacons, _ := compute(scanners)
	result := part1(beacons)

	if result != 79 {
		t.Fatalf("Expected 79 but got %d", result)
	}
}

func TestPart2(t *testing.T) {
	rawScanners := computeScanners(strings.Split(example, "\n\n"))
	_, scanners := compute(rawScanners)
	result := part2(scanners)

	if result != 3621 {
		t.Fatalf("Expected 3621 but got %d", result)
	}
}
