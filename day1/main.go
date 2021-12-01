package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(path string) ([]string, error) {
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

		current, cErr := strconv.Atoi(lines[i])
		if cErr != nil {
			fmt.Fprintln(os.Stderr, "Line of input is not a number", cErr)
			return
		}

		prev, _ := strconv.Atoi(lines[i-1])

		if current > prev {
			fmt.Printf("Found increase from %d to %d at index %d\n", current, prev, i)
			increases++
		}
	}

	fmt.Printf("Part 1 result: %d increases!\n", increases)
}
