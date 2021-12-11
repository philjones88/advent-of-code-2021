package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	example = `5483143223
	2745854711
	5264556173
	6141336146
	6357385478
	4167524645
	2176841721
	6882881134
	4846848554
	5283751526`
)

type coord struct {
	x int
	y int
}

func prettyPrintOctopuses(octopuses map[coord]int) {
	grid := make([][]int, 10, 10)

	for i := 0; i < 10; i++ {
		grid[i] = make([]int, 10)
	}

	for octopus, level := range octopuses {
		grid[octopus.x][octopus.y] = level
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			fmt.Print(grid[x][y])
		}
		fmt.Print("\n")
	}
}

func linesToOctopusMap(lines []string) map[coord]int {
	result := make(map[coord]int)

	for x, line := range lines {
		lineSplitIntoStrings := strings.Split(strings.TrimSpace(line), "")
		for y, char := range lineSplitIntoStrings {
			level, _ := strconv.Atoi(char)
			result[coord{x, y}] = level
		}
	}
	return result
}

func readFileLines(path string) ([]string, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	var inputLines []string
	buffer := bufio.NewScanner(input)
	for buffer.Scan() {
		inputLines = append(inputLines, buffer.Text())
	}

	return inputLines, buffer.Err()
}

func flashOctopus(flashes int, octopuses map[coord]int, flashedOctopuses map[coord]bool, octopus coord) (int, map[coord]int, map[coord]bool) {
	// If it's already flashed, ignore it
	if flashedOctopuses[octopus] {
		return flashes, octopuses, flashedOctopuses
	}
	flashes++
	flashedOctopuses[octopus] = true

	for _, xX := range []int{1, 0, -1} {
		for _, yY := range []int{1, 0, -1} {
			// check we're not at the same coord
			if xX == 0 && yY == 0 {
				continue
			}
			toCheck := coord{x: octopus.x + xX, y: octopus.y + yY}
			// I'm lazy and check it's a valid point
			if _, inMap := octopuses[toCheck]; !inMap {
				continue
			}
			// Don't flash a flashed octopus again!
			if flashedOctopuses[toCheck] {
				continue
			}

			octopuses[toCheck] = octopuses[toCheck] + 1
			if octopuses[toCheck] > 9 {
				// todo: is there a nicer way than storing + assigning like this?
				a, b, c := flashOctopus(flashes, octopuses, flashedOctopuses, toCheck)
				flashes = a
				octopuses = b
				flashedOctopuses = c
			}
		}
	}

	return flashes, octopuses, flashedOctopuses
}

func part1(octopuses map[coord]int) int {
	totalFlashedCount := 0

	for day := 1; day <= 100; day++ {
		flashedCount := 0
		flashedOctopuses := make(map[coord]bool)

		for octopus := range octopuses {
			octopuses[octopus]++
			if octopuses[octopus] > 9 {
				// todo: is there a nicer way than storing + assigning like this?
				a, b, c := flashOctopus(flashedCount, octopuses, flashedOctopuses, octopus)

				flashedCount = a
				octopuses = b
				flashedOctopuses = c
			}
		}

		totalFlashedCount += flashedCount

		// set all flashed octopuses to 0
		for flashedOctopus := range flashedOctopuses {
			octopuses[flashedOctopus] = 0
		}

		fmt.Printf("After Day %d:\n", day)
		prettyPrintOctopuses(octopuses)
		fmt.Print("\n")
	}

	return totalFlashedCount
}

func part2(octopuses map[coord]int) int {

	for day := 1; day <= 10000000; day++ {
		flashedOctopuses := make(map[coord]bool)

		for octopus := range octopuses {
			octopuses[octopus]++
			if octopuses[octopus] > 9 {
				_, b, c := flashOctopus(0, octopuses, flashedOctopuses, octopus)

				octopuses = b
				flashedOctopuses = c
			}
		}

		// we can tell if all are synced because flashed == octopuses map in len
		if len(flashedOctopuses) == len(octopuses) {
			return day
		}

		// set all flashed octopuses to 0
		for flashedOctopus := range flashedOctopuses {
			octopuses[flashedOctopus] = 0
		}

		fmt.Printf("After Day %d:\n", day)
		prettyPrintOctopuses(octopuses)
		fmt.Print("\n")
	}

	return -1
}

func main() {
	inputLines, _ := readFileLines("./day11/input")

	octopuses := linesToOctopusMap(inputLines)

	fmt.Print("Starting Octopuses:\n")
	prettyPrintOctopuses(octopuses)
	fmt.Print("\n")

	resultPart1 := part1(octopuses)

	fmt.Printf("Part 1 = %d\n", resultPart1)

	octopuses = linesToOctopusMap(inputLines)

	resultPart2 := part2(octopuses)

	fmt.Printf("Part 2 = %d\n", resultPart2)
}
