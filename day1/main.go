package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func readInput(path string) ([]int, error) {
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

	var lines []int

	for _, s := range inputLines {
		converted, err := strconv.Atoi(s)
		if err != nil {
			return lines, err
		}
		lines = append(lines, converted)
	}

	return lines, buffer.Err()
}

func main() {
	log.Println("hello!")
	lines, err := readInput("./day1/input")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file", err)
		return
	}

	var increases = 0

	for i := range lines {
		if i == 0 {
			continue
		}

		current := lines[i]
		prev := lines[i-1]

		if current > prev {
			fmt.Printf("Found increase from %d to %d at index %d\n", current, prev, i)
			increases++
		}
	}

	fmt.Printf("Part 1 result: %d increases!\n", increases)

	var windowIncreases = 0

	for i := range lines {
		if i == 0 {
			continue
		}
		window := sum(lines[i : i+3])
		prevWindow := sum(lines[i-1 : i+2])

		if window > prevWindow {
			fmt.Printf("Found increase from %d to %d at index %d\n", window, prevWindow, i)
			windowIncreases++
		}
	}

	fmt.Printf("Part 2 result: %d increases!\n", windowIncreases)
}
