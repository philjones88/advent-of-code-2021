package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	example = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`
)

type SnailFishNumber struct {
	value int
	depth int
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

func parseToSnailFishNumberPairts(stringSnailfish string) []SnailFishNumber {
	raw := []rune(stringSnailfish)
	result := []SnailFishNumber{}

	depth := 0

	for i := 0; i < len(raw); i++ {
		switch raw[i] {
		case '[':
			depth++
			continue
		case ']':
			depth--
			continue
		case ',':
			continue
		default:
			result = append(result, SnailFishNumber{
				value: stringToInt(string(raw[i])),
				depth: depth,
			})
		}
	}
	return result
}

func computeFinal(raw []string) []SnailFishNumber {
	current := parseToSnailFishNumberPairts(raw[0])
	raw = raw[1:]

	hasReduced := true
	for hasReduced {
		hasReduced = false

		for i := range current {
			if current[i].depth == 5 {
				if i > 0 {
					current[i-1].value += current[i].value
				}
				if i+2 < len(current) {
					current[i+2].value += current[i+1].value
				}

				current[i].value = 0
				current[i].depth--
				current = append(current[:i+1], current[i+2:]...)
				hasReduced = true
				break
			}
		}
		if hasReduced {
			continue
		}

		for i := range current {
			if current[i].value >= 10 {
				current[i].depth++

				new := SnailFishNumber{
					depth: current[i].depth,
					value: current[i].value/2 + current[i].value%2,
				}

				current[i].value /= 2

				appended := append([]SnailFishNumber{}, current[i+1:]...)
				current = append(current[:i+1], new)
				current = append(current, appended...)
				hasReduced = true
				break
			}
		}
		if hasReduced {
			continue
		}

		if len(raw) > 0 {
			hasReduced = true
			current = append(current, parseToSnailFishNumberPairts(raw[0])...)
			raw = raw[1:]

			for i := range current {
				current[i].depth++
			}
		}
	}

	return current
}

func computeSum(rawSnailFish []string) int {
	result := computeFinal(rawSnailFish)

	for depth := 4; depth >= 0; {
		found := false
		if len(result) == 1 {
			break
		}
		for i := range result {
			if result[i].depth == depth {
				result[i].depth--
				result[i].value = 3*result[i].value + 2*result[i+1].value

				found = true

				new := append([]SnailFishNumber{}, result[:i+1]...)

				if len(result) > i+2 {
					new = append(new, result[i+2:]...)
				}

				result = new
				break
			}
		}

		if !found {
			depth--
		}
	}

	return result[0].value
}

func computeMagnitude(rawSnailFish []string) int {
	max := 0

	for i := 0; i < len(rawSnailFish)-1; i++ {
		for ii := i + 1; ii < len(rawSnailFish); ii++ {
			magnitude := computeSum([]string{rawSnailFish[i], rawSnailFish[ii]})

			if magnitude > max {
				max = magnitude
			}

			magnitude = computeSum([]string{rawSnailFish[ii], rawSnailFish[i]})

			if magnitude > max {
				max = magnitude
			}
		}
	}

	return max
}

func part1(inputLines []string) int {
	return computeSum(inputLines)
}

func part2(inputLines []string) int {
	return computeMagnitude(inputLines)
}

func main() {
	inputLines, _ := readFileLines("input")

	part1Result := part1(inputLines)
	fmt.Printf("Part 1 = %d\n", part1Result)

	part2Result := part2(inputLines)
	fmt.Printf("Part 2 = %d\n", part2Result)
}
