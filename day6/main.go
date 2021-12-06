package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	example = "3,4,3,1,2"
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

func prettyPrintFishs(fishes map[int]int) {
	var keys []int
	for key := range fishes {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for key := range keys {
		fmt.Printf("Stage %d = %d\n", key, fishes[key])
	}
}

func computeStartingFishStages(numbers []int) map[int]int {
	fish := make(map[int]int)

	for _, number := range numbers {
		fish[number]++
	}

	return fish
}

func countFish(fish map[int]int) int {
	fishCount := 0
	for stage := range fish {
		fishCount = fishCount + fish[stage]
	}
	return fishCount
}

func part1(fish map[int]int) map[int]int {
	for day := 1; day <= 80; day++ {

		preStage0 := fish[0]

		for stage := 0; stage <= 8; stage++ {
			if stage == 0 {
				continue
			}
			fish[stage-1] = fish[stage-1] + fish[stage]
			fish[stage] = 0

		}

		fish[6] = fish[6] + preStage0
		fish[8] = fish[8] + preStage0
		fish[0] = fish[0] - preStage0
	}
	return fish
}

func main() {
	numbers, _ := readFileNumbers("./day6/input")

	fish := computeStartingFishStages(numbers)

	prettyPrintFishs(fish)

	resultFish := part1(fish)

	fmt.Print("------\n")

	prettyPrintFishs(resultFish)

	result := countFish(fish)

	fmt.Printf("Part 1 = %d\n", result)
}
