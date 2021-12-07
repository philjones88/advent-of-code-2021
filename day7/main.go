package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

const (
	example = "16,1,2,0,4,2,7,1,2,14"
)

func readFileNumbers(path string) ([]int, error) {
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

	var numbers []int

	for _, line := range inputLines {
		stringNumbers := strings.Split(line, ",")
		for _, stringNumber := range stringNumbers {
			parsed, _ := strconv.Atoi(stringNumber)
			numbers = append(numbers, parsed)
		}
	}

	return numbers, buffer.Err()
}

func stringToInt(str string) int {
	parsed, err := strconv.Atoi(str)
	if err != nil {
		panic("Rut roh! Not a number")
	}
	return parsed
}

func part1(crabPositions []int) int {
	sort.Ints(crabPositions)

	// Could brute force? but find the most common crab "line", this will save time looping
	average := funk.SumInt(crabPositions)/len(crabPositions) + 1

	minCrabPosition := funk.MinInt(crabPositions)

	bestPossibleCrabPositions := make(map[int]int)

	for i := minCrabPosition; i <= average; i++ {
		cost := 0
		for _, crabPosition := range crabPositions {
			cost += int(math.Abs(float64(crabPosition - i)))
		}
		bestPossibleCrabPositions[i] = cost
	}

	bestFuelSolution := bestPossibleCrabPositions[0]

	for _, fuel := range bestPossibleCrabPositions {
		if fuel < bestFuelSolution {
			bestFuelSolution = fuel
		}
	}

	return bestFuelSolution
}

func part2(crabPositions []int) int {
	sort.Ints(crabPositions)

	// Could brute force? but find the most common crab "line", this will save time looping
	average := funk.SumInt(crabPositions)/len(crabPositions) + 1

	minCrabPosition := funk.MinInt(crabPositions)

	bestPossibleCrabPositions := make(map[int]int)

	for i := minCrabPosition; i <= average; i++ {
		cost := 0
		for _, crabPosition := range crabPositions {
			moveAmount := int(math.Abs(float64(crabPosition - i)))
			cost += (moveAmount * (moveAmount + 1)) / 2
		}
		bestPossibleCrabPositions[i] = cost
	}

	bestFuelSolution := bestPossibleCrabPositions[0]

	for _, fuel := range bestPossibleCrabPositions {
		if fuel < bestFuelSolution {
			bestFuelSolution = fuel
		}
	}

	return bestFuelSolution
}

func main() {
	stringNumbers, _ := readFileNumbers("./day7/input")
	part1Result := part1(stringNumbers)

	fmt.Printf("Part 1 = %d\n", part1Result)

	part2Result := part2(stringNumbers)

	fmt.Printf("Part 2 = %d\n", part2Result)
}
