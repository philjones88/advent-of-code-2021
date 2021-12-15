package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	example = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

	possibleDirections = []CoOrd{
		// up
		{x: 0, y: 1},
		// right
		{x: 1, y: 0},
		// down
		{x: 0, y: -1},
		// left
		{x: -1, y: 0},
	}
)

func main() {
	inputLines, _ := readFileLines("input")

	grid := makeGrid(inputLines)

	part1Result := findShortestPath(grid)

	fmt.Printf("Part 1 = %d\n", part1Result)

	biggerGrid := expandGrid(grid)

	part2Result := findShortestPath(biggerGrid)

	fmt.Printf("Part 2 = %d\n", part2Result)
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
		inputLines = append(inputLines, strings.TrimSpace(buffer.Text()))
	}

	return inputLines, buffer.Err()
}

func stringToInt(str string) int {
	parsed, err := strconv.Atoi(str)
	if err != nil {
		panic("Rut roh! Not a number")
	}
	return parsed
}

func makeGrid(inputLines []string) [][]int {
	grid := make([][]int, len(inputLines))
	for i, line := range inputLines {
		lineSplit := strings.Split(strings.TrimSpace(line), "")
		row := make([]int, len(lineSplit))
		for ii, num := range lineSplit {
			row[ii] = stringToInt(num)
		}
		grid[i] = row
	}
	return grid
}

func findShortestPath(grid [][]int) int {
	gridWidth := len(grid[0]) - 1
	gridHeight := len(grid) - 1

	startCoOrd := CoOrd{x: 0, y: 0}
	endCoOrd := CoOrd{x: gridWidth, y: gridHeight}

	pq := make(PriorityQueue, 0)

	heap.Init(&pq)
	heap.Push(&pq, &CoOrdHeap{
		position: startCoOrd,
		priority: grid[0][0],
	})

	coOrdCosts := map[CoOrd]int{
		startCoOrd: grid[0][0],
	}

	visitedCoOrds := map[CoOrd]CoOrd{
		startCoOrd: startCoOrd,
	}

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*CoOrdHeap)

		if current.position == endCoOrd {
			break
		}
		for _, neighbourPoint := range neighbourCoOrds(current.position, grid) {

			newCost := coOrdCosts[current.position] + grid[neighbourPoint.y][neighbourPoint.x]

			if val, ok := coOrdCosts[neighbourPoint]; !ok || newCost < val {
				visitedCoOrds[neighbourPoint] = current.position
				coOrdCosts[neighbourPoint] = newCost
				heap.Push(&pq, &CoOrdHeap{
					position: neighbourPoint,
					priority: newCost,
				})
			}
		}
	}

	costSum := 0
	currentCoOrd := endCoOrd

	for currentCoOrd != startCoOrd {
		costSum += grid[currentCoOrd.y][currentCoOrd.x]
		currentCoOrd = visitedCoOrds[currentCoOrd]
	}

	return costSum
}

func expandGrid(grid [][]int) [][]int {
	// make a 5 time bigger grid
	biggerGrid := make([][]int, len(grid)*5)

	for i := 0; i < len(biggerGrid); i++ {
		biggerGrid[i] = make([]int, len(grid[0])*5)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			biggerGrid[y][x] = grid[y][x]
		}
	}
	for y := 0; y < len(grid); y++ {
		for x := len(grid[y]); x < len(biggerGrid[y]); x++ {
			newValue := (biggerGrid[y][x-len(grid[y])] + 1) % 10
			if newValue == 0 {
				newValue = 1
			}
			biggerGrid[y][x] = newValue
		}
	}

	for y := len(grid); y < len(biggerGrid); y++ {
		for x := 0; x < len(biggerGrid[y]); x++ {
			newValue := (biggerGrid[y-len(grid)][x] + 1) % 10
			if newValue == 0 {
				newValue = 1
			}
			biggerGrid[y][x] = newValue
		}
	}

	return biggerGrid
}

func neighbourCoOrds(p CoOrd, grid [][]int) []CoOrd {
	result := make([]CoOrd, 0)

	for _, d := range possibleDirections {
		newCoOrd := CoOrd{x: p.x + d.x, y: p.y + d.y}

		if newCoOrd.x >= 0 && newCoOrd.x < len(grid[p.y]) && newCoOrd.y >= 0 && newCoOrd.y < len(grid) {
			result = append(result, newCoOrd)
		}
	}

	return result
}
