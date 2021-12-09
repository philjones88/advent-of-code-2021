package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	example = `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`
)

//       North
// West        East
//       South

func stringToInt(str string) int {
	parsed, err := strconv.Atoi(str)
	if err != nil {
		panic("Rut roh! Not a number")
	}
	return parsed
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

func convertToIntLines(inputLines []string) [][]int {
	outputLines := make([][]int, len(inputLines))

	for i, line := range inputLines {
		tempLines := strings.Split(strings.TrimSpace(line), "")
		for _, tempLine := range tempLines {
			temp := stringToInt(tempLine)
			outputLines[i] = append(outputLines[i], temp)
		}
	}

	return outputLines
}

func part1(heightMapLines [][]int) int {
	risk := 0

	for row := 0; row < len(heightMapLines); row++ {

		for col := 0; col < len(heightMapLines[row]); col++ {

			results := []bool{true, true, true, true}

			// Check N, E, S, W

			// North but handle beight on the edge (row is 0)
			if row != 0 {
				results[0] = heightMapLines[row-1][col] > heightMapLines[row][col]
			}

			// East, but handle on the edge (col is last in row)
			if col != len(heightMapLines[row])-1 {
				results[1] = heightMapLines[row][col+1] > heightMapLines[row][col]
			}

			// South, but handle on the edge (last row)
			if row != len(heightMapLines)-1 {
				results[2] = heightMapLines[row+1][col] > heightMapLines[row][col]
			}

			// West, but handle on the edge (col is 0)
			if col != 0 {
				results[3] = heightMapLines[row][col-1] > heightMapLines[row][col]
			}

			// If N + E + S + W are all "true" (lowest)
			if results[0] && results[1] && results[2] && results[3] {
				// "risk level" is 1 plus its height
				risk += heightMapLines[row][col] + 1
			}
		}

	}

	return risk
}

func main() {
	inputLines, _ := readFileLines("./day9/input")

	heightMap := convertToIntLines(inputLines)

	part1Result := part1(heightMap)

	fmt.Printf("Part 1 = %d\n", part1Result)
}
