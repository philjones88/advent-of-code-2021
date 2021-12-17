package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	example = `target area: x=20..30, y=-10..-5`
)

type CoOrd struct {
	x int
	y int
}

func readFile(path string) (string, error) {
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

func calculate(x1 int, x2 int, y1 int, y2 int, position CoOrd, velocity CoOrd) (int, bool) {
	maxY := 0

	for {
		position.x += velocity.x
		position.y += velocity.y
		if position.y > maxY {
			maxY = position.y
		}
		if position.x >= x1 && position.x <= x2 && position.y >= y1 && position.y <= y2 {
			return maxY, true
		}
		if velocity.x == 0 && velocity.y < y1 {
			return maxY, false
		}
		if velocity.x > 0 {
			velocity.x--
		}
		if velocity.x < 0 {
			velocity.x++
		}
		velocity.y--
	}
}

func part1(input string) int {
	fmt.Printf("Input: %s\n", input)

	x1 := 0
	x2 := 0
	y1 := 0
	y2 := 0

	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	fmt.Printf("x1: %d, x2: %d, y1: %d, y2: %d\n", x1, x2, y1, y2)

	maxY := 0

	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			position := CoOrd{x, y}

			if max, ok := calculate(x1, x2, y1, y2, CoOrd{0, 0}, position); ok {
				if max > maxY {
					maxY = max
				}
			}
		}
	}

	return maxY
}

func part2(input string) int {
	fmt.Printf("Input: %s\n", input)

	x1 := 0
	x2 := 0
	y1 := 0
	y2 := 0

	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	fmt.Printf("x1: %d, x2: %d, y1: %d, y2: %d\n", x1, x2, y1, y2)

	grid := make(map[CoOrd]bool)

	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			position := CoOrd{x, y}

			if _, ok := calculate(x1, x2, y1, y2, CoOrd{0, 0}, position); ok {
				grid[position] = true
			}
		}
	}

	return len(grid)
}

func main() {
	input, _ := readFile("input")

	part1Result := part1(input)

	fmt.Printf("Part 1 = %d\n", part1Result)

	part2Result := part2(input)

	fmt.Printf("Part 2 = %d\n", part2Result)
}
