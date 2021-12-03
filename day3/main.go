package main

import (
	"bufio"
	"fmt"
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

func part1(lines []string) (int64, int64, int64, error) {
	numberOfLines := len(lines)
	numberOfBitsPerLine := len(lines[0])

	fmt.Printf("number of lines %d\n", numberOfLines)
	fmt.Printf("per line %d\n\n", numberOfBitsPerLine)

	gamma := ""
	epsilon := ""

	for i := 0; i < numberOfBitsPerLine; i++ {
		on := 0
		off := 0
		for ii := 0; ii < numberOfLines; ii++ {
			bit, _ := strconv.Atoi(string(string(lines[ii])[i]))

			if bit == 1 {
				on++
			} else {
				off++
			}
		}

		fmt.Printf("column %d = %d vs %d\n", i, on, off)

		if on > off {
			gamma += "1"
		} else {
			gamma += "0"
		}

		if on < off {
			epsilon += "1"
		} else {
			epsilon += "0"
		}

		fmt.Printf("column %d = gamma = %s = eplison = %s\n", i, gamma, epsilon)
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("Gamma = %d\n", gammaInt)
	fmt.Printf("Epsilon = %d\n", epsilonInt)
	fmt.Printf("Power = %d\n", gammaInt * epsilonInt)

	return gammaInt, epsilonInt, gammaInt * epsilonInt, nil
}

func main() {
	lines, _ := readInput("./day3/input")

	part1(lines)
}
