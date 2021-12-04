package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	example = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

	 22 13 17 11  0
	  8  2 23  4 24
	 21  9 14 16  7
	  6 10  3 18  5
	  1 12 20 15 19
	 
	  3 15  0  2 22
	  9 18 13 17  5
	 19  8  7 25 23
	 20 11 10 24  4
	 14 21 16 12  6
	 
	 14 21 17 24  4
	 10 16 15  9 19
	 18  8 23 26 20
	 22 11 13  6  5
	  2  0 12  3  7`
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

func parseIntoNumbersAndBoards(lines []string) ([]int, map[int][][]int, error) {
	var numbers []int
	boards := make(map[int][][]int)

	boardIndex := 0
	boardRowIndex := 0

	for i, line := range lines {
		if i == 0 {
			fmt.Printf("Parsing first line of numbers %s\n", line)

			stringNumbers := strings.Split(strings.TrimSpace(line), ",")

			for _, stringNumber := range stringNumbers {
				parsedStringNumber, _ := strconv.Atoi(stringNumber)
				fmt.Printf("Parsing string to number %d\n", parsedStringNumber)
				numbers = append(numbers, parsedStringNumber)
			}
			continue
		}

		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			boardRowIndex = 0
			boardIndex++
			fmt.Print("\n")
			continue
		}

		if boardRowIndex == 0 {
			boards[boardIndex] = make([][]int, 5, 5)
		}

		board := boards[boardIndex]

		stringNumbers := strings.Fields(trimmedLine)

		for _, stringNumber := range stringNumbers {
			parsedStringNumber, _ := strconv.Atoi(strings.TrimSpace(stringNumber))
			fmt.Printf("%d,", parsedStringNumber)
			board[boardRowIndex] = append(board[boardRowIndex], parsedStringNumber)
		}

		boards[boardIndex] = board

		fmt.Print("\n")
		boardRowIndex++
	}
	return numbers, boards, nil
}

func isInIntArray(array []int, toFind int) bool {
	for _, item := range array {
		if item == toFind {
			return true
		}
	}
	return false
}

func hasBoardWon(numbers []int, board [][]int) bool {
	for r := 0; r < 5; r++ {
		rowWinningCount := 0
		colWinningCount := 0

		for c := 0; c < 5; c++ {
			if isInIntArray(numbers, board[r][c]) {
				rowWinningCount++
			}
			if isInIntArray(numbers, board[c][r]) {
				colWinningCount++
			}
		}

		if rowWinningCount == 5 || colWinningCount == 5 {
			return true
		}
	}
	return false
}

func play(numbersToPlay []int, boards map[int][][]int) ([]int, [][]int) {
	for i := range numbersToPlay {
		// todo: could we skip the first 4 as you need 5 to win?
		numbersToPlaySoFar := numbersToPlay[:i]

		for _, board := range boards {
			if hasBoardWon(numbersToPlaySoFar, board) {
				return numbersToPlaySoFar, board
			}
		}
	}

	return nil, nil
}

func calculateScore(winningNumbers []int, winningBoard [][]int) int {
	score := 0

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !isInIntArray(winningNumbers, winningBoard[r][c]) {
				score = score + winningBoard[r][c]
			}
		}
	}
	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Winning Number: %d\n", winningNumbers[len(winningNumbers)-1])
	return score * winningNumbers[len(winningNumbers)-1]
}

func part1() {
	lines, _ := readFileLines("./day4/input")
	
	numbers, boards, _ := parseIntoNumbersAndBoards(lines)

	fmt.Printf("Numbers: %d\n", len(numbers))
	fmt.Printf("Boards count: %d\n", len(boards))

	winningNumbers, winningBoard := play(numbers, boards)

	fmt.Printf("Winning numbers: %+v\n", winningNumbers)
	fmt.Printf("Winning board: %+v\n", winningBoard)

	score := calculateScore(winningNumbers, winningBoard)

	fmt.Printf("Score: %d\n", score)
}

func main() {
	part1()
}
