package main

import (
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestComputeGrid(t *testing.T) {
	actual := computeSeaCucumberGrid([]string{"...>>>>>..."})

	if !reflect.DeepEqual(actual, [][]int{{-1, -1, -1, 1, 1, 1, 1, 1, -1, -1, -1}}) {
		t.Fatalf("Compute grid is wrong! Got: %+v", actual)
	}
}

func TestMoveExample0Step1(t *testing.T) {
	actual := computeSeaCucumberGrid([]string{"...>>>>>..."})
	expected := computeSeaCucumberGrid([]string{"...>>>>.>.."})

	gridResult := step(actual)

	if !reflect.DeepEqual(gridResult, expected) {
		t.Fatalf("Step is wrong! Got: %+v", gridResult)
	}
}

func TestMoveExample0Step2(t *testing.T) {
	actual := computeSeaCucumberGrid([]string{"...>>>>.>.."})
	expected := computeSeaCucumberGrid([]string{"...>>>.>.>."})

	gridResult := step(actual)

	if !reflect.DeepEqual(gridResult, expected) {
		t.Fatalf("Step is wrong! Got: %+v", gridResult)
	}
}

func TestMoveExample0Step3(t *testing.T) {
	actual := computeSeaCucumberGrid([]string{"...>>>.>.>."})
	expected := computeSeaCucumberGrid([]string{"...>>.>.>.>"})

	gridResult := step(actual)

	if !reflect.DeepEqual(gridResult, expected) {
		t.Fatalf("Step is wrong! Got: %+v", gridResult)
	}
}

func TestMoveExample0Step4(t *testing.T) {
	actual := computeSeaCucumberGrid([]string{"...>>.>.>.>"})
	expected := computeSeaCucumberGrid([]string{">..>.>.>.>."})

	gridResult := step(actual)

	if !reflect.DeepEqual(gridResult, expected) {
		t.Fatalf("Step is wrong! Got: %+v", gridResult)
	}
}

func TestPart1Step1(t *testing.T) {
	actual := computeSeaCucumberGrid(strings.Split(`v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`, "\n"))

	expected := computeSeaCucumberGrid(strings.Split(`....>.>v.>
v.v>.>v.v.
>v>>..>v..
>>v>v>.>.v
.>v.v...v.
v>>.>vvv..
..v...>>..
vv...>>vv.
>.v.v..v.v`, "\n"))

	gridResult, _ := part1(actual, 1)

	if !reflect.DeepEqual(gridResult, expected) {
		t.Fatalf("They don't match, got %+v", gridResult)
	}
}

func TestPart1Step2(t *testing.T) {
	actual := computeSeaCucumberGrid(strings.Split(`v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`, "\n"))

	expected := computeSeaCucumberGrid(strings.Split(`>.v.v>>..v
v.v.>>vv..
>v>.>.>.v.
>>v>v.>v>.
.>..v....v
.>v>>.v.v.
v....v>v>.
.vv..>>v..
v>.....vv.`, "\n"))

	gridResult, _ := part1(actual, 2)

	if reflect.DeepEqual(gridResult, expected) {
		t.Fatalf("They don't match, got %+v", gridResult)
	}
}

func TestPart1Step3(t *testing.T) {
	actual := computeSeaCucumberGrid(strings.Split(`v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`, "\n"))

	expected := computeSeaCucumberGrid(strings.Split(`v>v.v>.>v.
v...>>.v.v
>vv>.>v>..
>>v>v.>.v>
..>....v..
.>.>v>v..v
..v..v>vv>
v.v..>>v..
.v>....v..`, "\n"))

	gridResult, _ := part1(actual, 3)

	if !reflect.DeepEqual(gridResult, expected) {
		t.Fatalf("They don't match, got %+v", gridResult)
	}
}

func TestPart1Example1(t *testing.T) {
	grid := computeSeaCucumberGrid(strings.Split(example1, "\n"))

	_, result := part1(grid, math.MaxInt)

	if result != 58 {
		t.Fatalf("Expected 58 but got %d", result)
	}
}
