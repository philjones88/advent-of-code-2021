package main

import "testing"

// undefined 		0
// up						1
// down					2
// forward			3

func TestPart1(t *testing.T) {
	input := []command{
		{
			dir:    3,
			amount: 5,
		},
		{
			dir:    2,
			amount: 5,
		},
		{
			dir:    3,
			amount: 8,
		},
		{
			dir:    1,
			amount: 3,
		},
		{
			dir:    2,
			amount: 8,
		},
		{
			dir:    3,
			amount: 2,
		},
	}

	actual := part1(input)

	if actual != 150 {
		t.Fatalf("Expected %d to be 150!", actual)
	}
}

func TestPart2(t *testing.T) {
	input := []command{
		{
			dir:    3,
			amount: 5,
		},
		{
			dir:    2,
			amount: 5,
		},
		{
			dir:    3,
			amount: 8,
		},
		{
			dir:    1,
			amount: 3,
		},
		{
			dir:    2,
			amount: 8,
		},
		{
			dir:    3,
			amount: 2,
		},
	}

	actual := part2(input)

	if actual != 900 {
		t.Fatalf("Expected %d to be 900", actual)
	}
}
