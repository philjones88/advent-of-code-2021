package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	example = `0,9 -> 5,9
	8,0 -> 0,8
	9,4 -> 3,4
	2,2 -> 2,1
	7,0 -> 7,4
	6,4 -> 2,0
	0,9 -> 2,9
	3,4 -> 1,4
	0,0 -> 8,8
	5,5 -> 8,2`
)

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

func computePart1(inputLines []string) int {
	lines := make(map[string]int)

	for _, line := range inputLines {
		var start [2]int
		var end [2]int

		fmt.Sscanf(line, "%d,%d -> %d,%d", &(start[0]), &(start[1]), &(end[0]), &(end[1]))

		// Only do vertical + horizontal for part 1
		if start[0] == end[0] || start[1] == end[1] {
			// Push start manually
			lines[fmt.Sprintf("%d,%d", start[0], start[1])]++

			for start[0] != end[0] || start[1] != end[1] {
				if start[0] > end[0] {
					start[0]--
				} else if start[0] < end[0] {
					start[0]++
				}

				if start[1] > end[1] {
					start[1]--
				} else if start[1] < end[1] {
					start[1]++
				}

				lines[fmt.Sprintf("%d,%d", start[0], start[1])]++
			}
		}
	}

	count := 0

	for _, line := range lines {
		if line >= 2 {
			count++
		}
	}
	return count
}

func computePart2(inputLines []string) int {
	lines := make(map[string]int)

	for _, line := range inputLines {
		var start [2]int
		var end [2]int

		fmt.Sscanf(line, "%d,%d -> %d,%d", &(start[0]), &(start[1]), &(end[0]), &(end[1]))

		// Do vertical + horizontal + diagonal
		if start[0] != end[0] || start[1] != end[1] {
			// Push start manually
			lines[fmt.Sprintf("%d,%d", start[0], start[1])]++

			for start[0] != end[0] || start[1] != end[1] {
				if start[0] > end[0] {
					start[0]--
				} else if start[0] < end[0] {
					start[0]++
				}

				if start[1] > end[1] {
					start[1]--
				} else if start[1] < end[1] {
					start[1]++
				}

				lines[fmt.Sprintf("%d,%d", start[0], start[1])]++
			}
		}
	}

	count := 0

	for _, line := range lines {
		if line >= 2 {
			count++
		}
	}
	return count
}

func main() {
	inputLines, _ := readFileLines("./day5/input")
	part1 := computePart1(inputLines)

	fmt.Printf("Part 1 = %d\n", part1)

	part2 := computePart2(inputLines)

	fmt.Printf("Part 2 = %d\n", part2)
}
