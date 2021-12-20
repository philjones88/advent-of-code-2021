package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	example = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`
)

type CoOrd struct {
	x int
	y int
}

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

func binaryToInt(b string) int {
	v, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic("Rut roh not a number")
	}
	n, err := strconv.Atoi(fmt.Sprintf("%d", v))
	if err != nil {
		panic(err)
	}
	return n
}

func parseAlgorithAndImage(input string) (string, map[CoOrd]bool) {
	inputSplit := strings.Split(input, "\n\n")

	image := make(map[CoOrd]bool)

	for x, line := range strings.Split(strings.TrimSpace(inputSplit[1]), "\n") {
		for y, char := range strings.Split(line, "") {
			image[CoOrd{x, y}] = char == "#"
		}
	}

	return strings.TrimSpace(inputSplit[0]), image
}

func countImage(image map[CoOrd]bool) int {
	result := 0
	for _, val := range image {
		if val {
			result++
		}
	}
	return result
}

func printGrid(image map[CoOrd]bool) {
	var minX int = math.MaxInt
	var minY int = math.MaxInt
	var maxX int = math.MinInt
	var maxY int = math.MinInt

	for key := range image {
		if key.x < minX {
			minX = key.x
		}
		if key.y < minY {
			minY = key.y
		}
		if key.x > maxX {
			maxX = key.x
		}
		if key.y > maxY {
			maxY = key.y
		}
	}
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			if val, exists := image[CoOrd{x, y}]; exists && val {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func enhance(algorithm string, image map[CoOrd]bool, iteration int) map[CoOrd]bool {
	var minX int = math.MaxInt
	var minY int = math.MaxInt
	var maxX int = math.MinInt
	var maxY int = math.MinInt

	for key := range image {
		if key.x < minX {
			minX = key.x
		}
		if key.y < minY {
			minY = key.y
		}
		if key.x > maxX {
			maxX = key.x
		}
		if key.y > maxY {
			maxY = key.y
		}
	}

	newGrid := make(map[CoOrd]bool)

	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			result := ""
			for _, xx := range []int{x - 1, x, x + 1} {
				for _, yy := range []int{y - 1, y, y + 1} {

					if lit, exists := image[CoOrd{xx, yy}]; exists {
						if lit {
							result += "1"
						} else {
							result += "0"
						}
					} else {
						// This handles example != input algorithms
						// maybe a better way but this works...
						if algorithm[0] == '#' && algorithm[511] == '.' && iteration%2 == 1 {
							result += "1"
						} else {
							result += "0"
						}
					}
				}
			}
			resultInt := binaryToInt(result)
			if algorithm[resultInt] == '#' {
				newGrid[CoOrd{x, y}] = true
			} else {
				newGrid[CoOrd{x, y}] = false
			}
		}
	}

	return newGrid
}

func part1(algorithm string, image map[CoOrd]bool) int {
	newImage := enhance(algorithm, image, 0)
	newImage = enhance(algorithm, newImage, 1)
	return countImage(newImage)
}

func part2(algorithm string, image map[CoOrd]bool) int {
	newImage := image
	for i := 0; i < 50; i++ {
		newImage = enhance(algorithm, newImage, i)
	}
	// printGrid(newImage)
	return countImage(newImage)
}

func main() {
	input, _ := readFile("input")
	// input := example

	algorithm, image := parseAlgorithAndImage(input)

	part1Result := part1(algorithm, image)

	fmt.Printf("Part 1 = %d\n", part1Result)

	part2Result := part2(algorithm, image)

	fmt.Printf("Part 2 = %d\n", part2Result)
}
