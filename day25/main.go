package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const (
	examplePhil = `.>.>
v.v.
vv..`

	example0 = `...>>>>>...`

	example1 = `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`
)

// > == east
// v == south
// . == empty

// empty = -1
// north = 0
// east  = 1
// south = 2
// west  = 3

func readFileLines(path string) ([]string, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	var inputLines []string
	buffer := bufio.NewScanner(input)
	for buffer.Scan() {
		inputLines = append(inputLines, strings.TrimSpace(buffer.Text()))
	}

	return inputLines, buffer.Err()
}

func computeGridSize(grid [][]int) (int, int, int, int) {
	return 0, len(grid), 0, len(grid[0])
}

func prettyPrintGrid(grid [][]int) {
	minX, maxX, minY, maxY := computeGridSize(grid)
	herd1 := 0
	herd2 := 0
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			sc := grid[x][y]

			if sc == 1 {
				fmt.Print(">")
				herd1++
			} else if sc == 2 {
				fmt.Print("v")
				herd2++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Printf("Total = %d (1=%d, 2=%d)\n", len(grid), herd1, herd2)
}

func computeSeaCucumberGrid(inputLines []string) [][]int {
	grid := make([][]int, len(inputLines))
	for x, line := range inputLines {
		lineSplit := strings.Split(line, "")

		grid[x] = []int{}

		for _, char := range lineSplit {
			var herd int
			switch char {
			case ">":
				herd = 1
			case "v":
				herd = 2
			case ".":
				herd = -1
			default:
				panic("Rut Roh! Not a valid sea cucumber herd!")
			}
			grid[x] = append(grid[x], herd)
		}
	}
	return grid
}

func step(grid [][]int) [][]int {
	minX, maxX, minY, maxY := computeGridSize(grid)

	// Clone grid to newGridForEast
	newGridForEast := make([][]int, len(grid))
	for x := minX; x < maxX; x++ {
		newGridForEast[x] = []int{}
		for y := minY; y < maxY; y++ {
			newGridForEast[x] = append(newGridForEast[x], grid[x][y])
		}
	}

	// Move east first
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			if grid[x][y] == 2 || grid[x][y] == -1 {
				continue
			}

			if (y+1) >= maxY && grid[x][0] == -1 {
				newGridForEast[x][0] = 1
				newGridForEast[x][y] = -1
			} else if (y+1) < maxY && grid[x][y+1] == -1 {
				newGridForEast[x][y+1] = 1
				newGridForEast[x][y] = -1
			}
		}
	}

	// Clone newGridForEast to newGridForSouth
	newGridForSouth := make([][]int, len(newGridForEast))
	for x := minX; x < maxX; x++ {
		newGridForSouth[x] = []int{}
		for y := minY; y < maxY; y++ {
			newGridForSouth[x] = append(newGridForSouth[x], newGridForEast[x][y])
		}
	}

	// Move south next
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			if newGridForEast[x][y] == 1 || newGridForEast[x][y] == -1 {
				continue
			}

			if (x+1) >= maxX && newGridForEast[0][y] == -1 {
				newGridForSouth[0][y] = 2
				newGridForSouth[x][y] = -1
			} else if (x+1) < maxX && newGridForEast[x+1][y] == -1 {
				newGridForSouth[x+1][y] = 2
				newGridForSouth[x][y] = -1
			}
		}
	}

	return newGridForSouth
}

func part1(grid [][]int, maxIterations int) ([][]int, int) {
	stopsMovingAt := 1
	for i := 0; i < maxIterations; i++ {
		newGrid := step(grid)

		if reflect.DeepEqual(grid, newGrid) {
			return grid, stopsMovingAt
		}

		grid = newGrid
		stopsMovingAt++
	}
	return grid, stopsMovingAt
}

func main() {
	inputLines, _ := readFileLines("input")
	// inputLines := strings.Split(example1, "\n")

	grid := computeSeaCucumberGrid(inputLines)

	fmt.Print("----------------------\n")
	prettyPrintGrid(grid)

	newGrid, result := part1(grid, 1000)

	fmt.Print("----------------------\n")
	prettyPrintGrid(newGrid)
	fmt.Print("----------------------\n")

	fmt.Printf("Part 1 = %d\n", result)
}
