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
	example = `2199943210
	3987894921
	9856789892
	8767896789
	9899965678`
)

//       North
// West        East
//       South

type CoOrdinates struct {
	x int
	y int
}

func stringToInt(str string) int {
	parsed, err := strconv.Atoi(str)
	if err != nil {
		panic("Rut roh! Not a number")
	}
	return parsed
}

func containsCoOrdinates(s []CoOrdinates, c CoOrdinates) bool {
	for _, v := range s {
		if v.x == c.x && v.y == c.y {
			return true
		}
	}
	return false
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

func checkPosition(x int, y int, heightMapLines [][]int) (bool, []CoOrdinates) {
	var resultsMap []CoOrdinates
	results := []bool{true, true, true, true}
	// North but handle beight on the edge (row is 0)
	if x != 0 {
		results[0] = heightMapLines[x-1][y] > heightMapLines[x][y]
		resultsMap = append(resultsMap, CoOrdinates{x: x - 1, y: y})
	}
	// East, but handle on the edge (col is last in row)
	if y != len(heightMapLines[x])-1 {
		results[1] = heightMapLines[x][y+1] > heightMapLines[x][y]
		resultsMap = append(resultsMap, CoOrdinates{x: x, y: y + 1})
	}
	// South, but handle on the edge (last row)
	if x != len(heightMapLines)-1 {
		results[2] = heightMapLines[x+1][y] > heightMapLines[x][y]
		resultsMap = append(resultsMap, CoOrdinates{x: x + 1, y: y})
	}
	// West, but handle on the edge (col is 0)
	if y != 0 {
		results[3] = heightMapLines[x][y-1] > heightMapLines[x][y]
		resultsMap = append(resultsMap, CoOrdinates{x: x, y: y - 1})
	}
	return results[0] && results[1] && results[2] && results[3], resultsMap
}

func checkPositionIsBasinEdge(x int, y int, heightMapLines [][]int) []CoOrdinates {
	var resultsMap []CoOrdinates

	if heightMapLines[x][y] == 9 {
		return resultsMap
	}

	// North but handle beight on the edge (row is 0)
	if x != 0 && heightMapLines[x-1][y] != 9 {
		resultsMap = append(resultsMap, CoOrdinates{x: x - 1, y: y})
	}
	// East, but handle on the edge (col is last in row)
	if y != len(heightMapLines[x])-1 && heightMapLines[x][y+1] != 9 {
		resultsMap = append(resultsMap, CoOrdinates{x: x, y: y + 1})
	}
	// South, but handle on the edge (last row)
	if x != len(heightMapLines)-1 && heightMapLines[x+1][y] != 9 {
		resultsMap = append(resultsMap, CoOrdinates{x: x + 1, y: y})
	}
	// West, but handle on the edge (col is 0)
	if y != 0 && heightMapLines[x][y-1] != 9 {
		resultsMap = append(resultsMap, CoOrdinates{x: x, y: y - 1})
	}
	return resultsMap
}

func getBasinCoOrdinates(coords CoOrdinates, heightMapLines [][]int) []CoOrdinates {

	var finalBasinCoOrds []CoOrdinates

	var coOrdsToCheck []CoOrdinates

	_, resultMap := checkPosition(coords.x, coords.y, heightMapLines)

	for _, a := range resultMap {
		if !containsCoOrdinates(coOrdsToCheck, a) {
			coOrdsToCheck = append(coOrdsToCheck, a)
		}
	}

	finalBasinCoOrds = append(finalBasinCoOrds, coords)

	for {
		coOrdToCheck := coOrdsToCheck[0]

		resultMap := checkPositionIsBasinEdge(coOrdToCheck.x, coOrdToCheck.y, heightMapLines)
		if len(resultMap) > 0 {
			for _, a := range resultMap {
				if !containsCoOrdinates(coOrdsToCheck, a) && !containsCoOrdinates(finalBasinCoOrds, a) {
					coOrdsToCheck = append(coOrdsToCheck, a)
				}
			}

			finalBasinCoOrds = append(finalBasinCoOrds, coOrdToCheck)
		}

		coOrdsToCheck = coOrdsToCheck[1:]

		if len(coOrdsToCheck) == 0 {
			break
		}
	}

	return finalBasinCoOrds
}

func part1(heightMapLines [][]int) int {
	risk := 0
	for row := 0; row < len(heightMapLines); row++ {
		for col := 0; col < len(heightMapLines[row]); col++ {
			// If N + E + S + W are all "true" (lowest)
			if isLowPoint, _ := checkPosition(row, col, heightMapLines); isLowPoint {
				// "risk level" is 1 plus its height
				risk += heightMapLines[row][col] + 1
			}
		}
	}
	return risk
}

func part2(heightMapLines [][]int) int {
	var lowPointBasins []CoOrdinates

	for row := 0; row < len(heightMapLines); row++ {
		for col := 0; col < len(heightMapLines[row]); col++ {
			if isLowPoint, _ := checkPosition(row, col, heightMapLines); isLowPoint {
				lowPointBasins = append(lowPointBasins, CoOrdinates{x: row, y: col})
			}
		}
	}

	var basinSizes []int

	for _, lowPoint := range lowPointBasins {
		basinCoOrdinates := getBasinCoOrdinates(lowPoint, heightMapLines)
		basinSizes = append(basinSizes, len(basinCoOrdinates))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func main() {
	inputLines, _ := readFileLines("./day9/input")

	heightMap := convertToIntLines(inputLines)

	part1Result := part1(heightMap)

	fmt.Printf("Part 1 = %d\n", part1Result)

	part2Result := part2(heightMap)

	fmt.Printf("Part 2 = %d\n", part2Result)
}
