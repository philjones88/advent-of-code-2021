package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CoOrd struct {
	x int
	y int
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

// Pen and paper like previous day to figure this out...

func compute(code []string) (min int, max int) {
	p1 := make([]int, 14)
	p2 := make([]int, 14)
	p3 := make([]int, 14)

	for it := 0; it < 14; it++ {
		p1[it] = stringToInt(code[it*18+4][6:])
		p2[it] = stringToInt(code[it*18+5][6:])
		p3[it] = stringToInt(code[it*18+15][6:])
	}

	coOrds := make([]CoOrd, 0)
	for len(coOrds) < 7 {
		ll := -1
		for i, p := range p1 {
			if p == 1 {
				ll = i
			} else if p == 26 {
				coOrds = append(coOrds, CoOrd{ll, i})
				p1[ll] = 2
				p1[i] = 27
				break
			}

		}
	}
	mMin := make([]int, 14)
	mMax := make([]int, 14)
	for _, coOrd := range coOrds {
		diff := p3[coOrd.x] + p2[coOrd.y]
		if diff > 0 {
			mMax[coOrd.x] = 9 - diff
			mMin[coOrd.x] = 1
			mMax[coOrd.y] = 9
			mMin[coOrd.y] = 1 + diff
		} else {
			mMax[coOrd.x] = 9
			mMin[coOrd.x] = 1 - diff
			mMax[coOrd.y] = 9 + diff
			mMin[coOrd.y] = 1
		}
	}

	for i := 0; i < 14; i++ {
		min *= 10
		max *= 10
		min += mMin[i]
		max += mMax[i]
	}

	return
}

func main() {
	inputLines, _ := readFileLines("input")

	min, max := compute(inputLines)

	fmt.Printf("Part 1 = %d\n", max)
	fmt.Printf("Part 2 = %d\n", min)
}
