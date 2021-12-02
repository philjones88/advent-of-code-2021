package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	undefined direction = iota
	up
	down
	forward
)

type command struct {
	dir    direction
	amount int
}

func directionToString(input direction) (string, error) {
	switch input {
	case up:
		return "up", nil
	case down:
		return "down", nil
	case forward:
		return "forward", nil
	default:
		return "undefined", nil
	}
}

func stringToDirection(input string) (direction, error) {
	switch input {
	case "up":
		return 1, nil
	case "down":
		return 2, nil
	case "forward":
		return 3, nil
	default:
		return 0, nil
	}
}

func readInput(path string) ([]command, error) {
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

	var lines []command

	for _, s := range inputLines {
		split := strings.Split(s, " ")

		if len(split) != 2 {
			return nil, os.ErrInvalid
		}

		dir, _ := stringToDirection(split[0])
		amount, _ := strconv.Atoi(split[1])

		newCommand := command{dir: dir, amount: amount}

		lines = append(lines, newCommand)
	}

	return lines, buffer.Err()
}

func part1(lines []command) {
	horizontal := 0
	depth := 0

	for i, line := range lines {
		switch line.dir {
		case up:
			depth = depth - line.amount
			fmt.Printf("Line %d command up on depth %d by %d\n", i, depth, line.amount)
			break
		case down:
			depth = depth + line.amount
			fmt.Printf("Line %d command down on depth %d by %d\n", i, depth, line.amount)
			break
		case forward:
			horizontal = horizontal + line.amount
			fmt.Printf("Line %d command forward on horizontal %d by %d\n", i, horizontal, line.amount)
			break
		}
	}

	fmt.Printf("Depth: %d\n", depth)
	fmt.Printf("Horizontal: %d\n", horizontal)
	fmt.Printf("Result: %d\n", depth*horizontal)
}

func part2(lines []command) {
	horizontal := 0
	depth := 0
	aim := 0

	for i, line := range lines {
		switch line.dir {
		case up:
			aim = aim - line.amount
			fmt.Printf("Line %d command up on aim %d\n", i, aim)
			break
		case down:
			aim = aim + line.amount
			fmt.Printf("Line %d command down on aim %d\n", i, aim)
			break
		case forward:
			horizontal = horizontal + line.amount
			depth = depth + (line.amount * aim)
			fmt.Printf("Line %d command forward on horizontal %d and depth %d\n", i, horizontal, depth)
		}
	}

	fmt.Printf("Depth: %d\n", depth)
	fmt.Printf("Horizontal: %d\n", horizontal)
	fmt.Printf("Result: %d\n", depth*horizontal)
}

func main() {
	lines, err := readInput("./day2/input")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input", err)
	}

	part1(lines)
	part2(lines)
}
