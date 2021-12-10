package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	example = `[({(<(())[]>[[{[]{<()<>>
		[(()[<>])]({[<{<<[]>>(
		{([(<{}[<>[]}>{[]{[(<()>
		(((({<>}<{<{<>}{[]{[]{}
		[[<[([]))<([[{}[[()]]]
		[{[{({}]{}}([{[{{{}}([]
		{<[[]]>}<{[{[{[]{()[[[]
		[<(<(<(<{}))><([]([]()
		<{([([[(<>()){}]>(<<{{
		<{([{{}}[<[[[<>{}]]]>[]]`
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

// ( to )
// [ to ]
// { to }
// < to >

func part1(inputLines []string) (int, int) {

	invalidChunkScore := 0
	var invalidChunkScore2 []int

	for _, line := range inputLines {
		splitLine := strings.Split(strings.TrimSpace(line), "")

		validChunks := map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
		validChunksReversed := map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}
		invalidChunkScoresClosers := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
		invalidChunkScoresOpeners := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}

		isInvalidLine := false
		invalidScore1 := 0

		var chunkQueue []string

		for _, char := range splitLine {
			if _, isAStartChar := validChunks[char]; isAStartChar {
				chunkQueue = append(chunkQueue, char)
			} else {
				if chunkQueue[len(chunkQueue)-1] != validChunksReversed[char] {
					invalidScore1 += invalidChunkScoresClosers[char]
					isInvalidLine = true
					break
				}
				chunkQueue = chunkQueue[0 : len(chunkQueue)-1]
			}
		}

		if isInvalidLine {
			invalidChunkScore += invalidScore1
		} else {
			closeScore := 0

			for i := len(chunkQueue) - 1; i >= 0; i-- {
				closeScore = closeScore*5 + invalidChunkScoresOpeners[chunkQueue[i]]
			}
			invalidChunkScore2 = append(invalidChunkScore2, closeScore)
		}
	}

	sort.Ints(invalidChunkScore2)

	return invalidChunkScore, invalidChunkScore2[len(invalidChunkScore2)/2]
}

func main() {
	inputLines, _ := readFileLines("./day10/input")

	part1, part2 := part1(inputLines)

	fmt.Printf("Part 1 = %d\n", part1)
	fmt.Printf("Part 2 = %d\n", part2)
}
