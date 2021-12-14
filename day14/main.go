package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thoas/go-funk"
)

const (
	example = `NNCB

	CH -> B
	HH -> N
	CB -> H
	NH -> C
	HB -> C
	HC -> B
	HN -> C
	NN -> C
	BH -> H
	NC -> B
	NB -> B
	BN -> B
	BB -> N
	BC -> B
	CC -> N
	CN -> C`
)

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

func processInput(input string) (string, map[string]string) {
	inputSplit := strings.Split(input, "\n\n")

	template := strings.TrimSpace(inputSplit[0])
	pairRuleStrings := strings.Split(inputSplit[1], "\n")

	pairRules := make(map[string]string)

	for _, line := range pairRuleStrings {
		lineSplit := strings.Split(strings.TrimSpace(line), " -> ")
		pairRules[lineSplit[0]] = lineSplit[1]
	}

	return template, pairRules
}

func computePairs(template string, pairRules map[string]string, steps int) map[string]int {
	pairs := make(map[string]int)

	// pre-populate from starting template
	for i := 0; i < len(template)-1; i++ {
		pairs[string(template[i])+string(template[i+1])]++
	}

	fmt.Printf("Starting pairs: %+v\n", pairs)

	for step := 1; step <= steps; step++ {
		newPairs := make(map[string]int)

		for key, count := range pairs {
			newPairs[string(key[0])+pairRules[key]] += count
			newPairs[pairRules[key]+string(key[1])] += count
		}

		fmt.Printf("New Pairs on step %d: %+v\n", step, newPairs)

		pairs = newPairs
	}

	fmt.Printf("New Pairs after %d steps: %+v\n", steps, pairs)
	return pairs
}

func countElements(pairs map[string]int) map[string]int {
	elements := make(map[string]int)

	for key, val := range pairs {
		elements[string(key[0])] += val
		elements[string(key[1])] += val
	}

	return elements
}

func calculate(template string, pairRules map[string]string, steps int) int {
	pairs := computePairs(template, pairRules, steps)

	fmt.Printf("Pairs: %+v\n", pairs)

	elementsCounts := countElements(pairs)

	fmt.Printf("Elements: %+v\n", elementsCounts)

	var counts []int

	for _, count := range elementsCounts {
		counts = append(counts, count)
	}

	fmt.Printf("Counts: %+v\n", counts)

	min := funk.MinInt(counts)
	max := funk.MaxInt(counts)

	return (max-min)/2 + 1
}

func main() {
	input, _ := readFile("./day14/input")
	template, pairRules := processInput(input)
	part1Result := calculate(template, pairRules, 10)

	fmt.Printf("Part 1 = %d\n", part1Result)

	part2Result := calculate(template, pairRules, 40)

	fmt.Printf("Part 2 = %d\n", part2Result)
}
