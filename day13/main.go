package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	example = `6,10
	0,14
	9,10
	0,3
	10,4
	4,11
	6,0
	6,12
	4,1
	0,13
	10,12
	3,4
	3,0
	8,4
	1,10
	2,14
	8,10
	9,0

	fold along y=7
	fold along x=5`
)

type coord struct {
	x int
	y int
}

func stringToInt(str string) int {
	parsed, err := strconv.Atoi(str)
	if err != nil {
		panic("Rut roh! Not a number")
	}
	return parsed
}

func readFileLines(path string) (string, error) {
	input, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer input.Close()

	var inputLines []string
	buffer := bufio.NewScanner(input)
	for buffer.Scan() {
		inputLines = append(inputLines, buffer.Text())
	}

	return strings.Join(inputLines, "\n"), buffer.Err()
}

func processInput(input string) (map[coord]bool, []string) {
	inputSplit := strings.Split(input, "\n\n")
	dots := strings.Split(inputSplit[0], "\n")

	grid := make(map[coord]bool)

	for _, dot := range dots {
		splitDotIntoCoOrds := strings.Split(strings.TrimSpace(dot), ",")
		x := stringToInt(splitDotIntoCoOrds[0])
		y := stringToInt(splitDotIntoCoOrds[1])
		grid[coord{x, y}] = true
	}

	instructions := strings.Split(strings.TrimSpace(inputSplit[1]), "\n")

	return grid, instructions
}

func compute(grid map[coord]bool, instructions []string, howManyInstructionsToDo int) map[coord]bool {

	for instructionIndex, instruction := range instructions {
		// split on =
		instructionSplit := strings.Split(instruction, "=")

		// byte to string stuff to get to a single char
		axis := string(instructionSplit[0][len(instructionSplit[0])-1])

		// get the amount of fold
		foldAmount := stringToInt(strings.TrimSpace(instructionSplit[1]))

		// make a new grid to hold the newly folded grid
		newGrid := make(map[coord]bool)

		fmt.Printf("Folding along %s=%d\n", axis, foldAmount)

		// y is across the grid horizontally, folding UPWARDS
		if axis == "y" {
			for newY := 0; newY < foldAmount; newY++ {
				for position := range grid {
					if position.y == newY || position.y == (foldAmount+foldAmount-newY) {
						newGrid[coord{x: position.x, y: newY}] = true
					}
				}
			}
		}

		// x is down the grid vertically, folding LEFTWARDS
		if axis == "x" {
			for newX := 0; newX < foldAmount; newX++ {
				for position := range grid {
					if position.x == newX || position.x == (foldAmount+foldAmount-newX) {
						newGrid[coord{x: newX, y: position.y}] = true
					}
				}
			}
		}

		// make sure we save the new grid onto the old grid
		grid = newGrid

		// handle part1 wanting just 1 of 2 instructions executed
		if instructionIndex+1 == howManyInstructionsToDo {
			return grid
		}
	}

	return grid
}

func part1(grid map[coord]bool, instructions []string) int {
	grid = compute(grid, instructions, 1)
	// Use the len of the map to count the dots
	return len(grid)
}

func part2(grid map[coord]bool, instructions []string) {
	grid = compute(grid, instructions, len(instructions)+1)

	printGrid(grid)
}

func printGrid(grid map[coord]bool) {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	// find the grid bounds
	for position := range grid {
		if position.x < minX {
			minX = position.x
		}
		if position.x > maxX {
			maxX = position.x
		}
		if position.y < minY {
			minY = position.y
		}
		if position.y > maxY {
			maxY = position.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[coord{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

}

func main() {
	inputLines, _ := readFileLines("./day13/input")

	grid, instructions := processInput(inputLines)

	part1Result := part1(grid, instructions)

	fmt.Printf("Part 1 = %d\n\n", part1Result)

	part2(grid, instructions)
}
