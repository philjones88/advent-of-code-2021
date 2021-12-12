package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	example1 = `start-A
	start-b
	A-c
	A-b
	b-d
	A-end
	b-end`

	example2 = `dc-end
	HN-start
	start-kj
	dc-start
	dc-HN
	LN-dc
	HN-end
	kj-sa
	kj-HN
	kj-dc`

	example3 = `fs-end
	he-DX
	fs-he
	start-DX
	pj-DX
	end-zg
	zg-sl
	zg-pj
	pj-he
	RW-he
	fs-DX
	pj-RW
	zg-RW
	start-pj
	he-WI
	zg-he
	pj-fs
	start-RW`
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

func makeCaveMapInstructions(inputLines []string) [][]string {
	var caveMap [][]string

	for _, inputLine := range inputLines {
		caveMap = append(caveMap, strings.Split(strings.TrimSpace(inputLine), "-"))
	}

	return caveMap
}

func iteratePart1(caveGraph map[string]map[string]bool, currentPosition string, visitedCaveMap map[string]bool) int {
	if currentPosition == "end" {
		return 1
	}

	pathsCount := 0

	for nextPath := range caveGraph[currentPosition] {
		if visitedCaveMap[nextPath] && strings.ToUpper(nextPath) != nextPath {
			continue
		}

		visitedCaveMap[currentPosition] = true

		pathsCount += iteratePart1(caveGraph, nextPath, visitedCaveMap)

		visitedCaveMap[nextPath] = false
	}

	return pathsCount
}

func part1(inputLines []string) int {
	instructions := makeCaveMapInstructions(inputLines)

	caveGraph := make(map[string]map[string]bool)

	for _, pair := range instructions {
		if _, exists := caveGraph[pair[0]]; !exists {
			caveGraph[pair[0]] = make(map[string]bool)
		}
		if _, exists := caveGraph[pair[1]]; !exists {
			caveGraph[pair[1]] = make(map[string]bool)
		}
		caveGraph[pair[0]][pair[1]] = true
		caveGraph[pair[1]][pair[0]] = true
	}

	currentPosition := "start"
	visitedCaveMap := map[string]bool{"start": true}

	return iteratePart1(caveGraph, currentPosition, visitedCaveMap)
}

func iteratePart2(caveGraph map[string]map[string]bool, currentPosition string, visitedCaveMap map[string]int, doubleVisit bool) int {
	if currentPosition == "end" {
		return 1
	}

	visitedCaveMap[currentPosition]++

	pathsCount := 0

	for nextPath := range caveGraph[currentPosition] {
		if nextPath == "start" {
			continue
		}

		if strings.ToUpper(nextPath) != nextPath && visitedCaveMap[nextPath] > 0 {
			if doubleVisit {
				continue
			}
			doubleVisit = true
		}

		pathsCount += iteratePart2(caveGraph, nextPath, visitedCaveMap, doubleVisit)

		visitedCaveMap[nextPath]--

		if strings.ToUpper(nextPath) != nextPath && visitedCaveMap[nextPath] == 1 {
			doubleVisit = false
		}
	}

	return pathsCount
}

func part2(inputLines []string) int {
	instructions := makeCaveMapInstructions(inputLines)

	caveGraph := make(map[string]map[string]bool)

	for _, pair := range instructions {
		if _, exists := caveGraph[pair[0]]; !exists {
			caveGraph[pair[0]] = make(map[string]bool)
		}
		if _, exists := caveGraph[pair[1]]; !exists {
			caveGraph[pair[1]] = make(map[string]bool)
		}
		caveGraph[pair[0]][pair[1]] = true
		caveGraph[pair[1]][pair[0]] = true
	}

	currentPosition := "start"
	visitedCaveMap := map[string]int{"start": 0}
	doubleVisit := false

	return iteratePart2(caveGraph, currentPosition, visitedCaveMap, doubleVisit)
}

func main() {
	inputLines, _ := readFileLines("./day12/input")

	resultPart1 := part1(inputLines)

	fmt.Printf("Part 1 = %d\n", resultPart1)

	resultPart2 := part2(inputLines)

	fmt.Printf("Part 2 = %d\n", resultPart2)
}
