package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

const (
	example = `Player 1 starting position: 4
	Player 2 starting position: 8`
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

func parsePlayersAndStartingPositions(inputLines []string) (players [2]int, positions [2]int) {
	for i, line := range inputLines {
		var player int
		var position int
		fmt.Sscanf(strings.TrimSpace(line), "Player %d starting position: %d", &player, &position)
		player--
		players[i] = player
		position--
		positions[i] = position
	}
	return
}

func part1(players [2]int, positions [2]int) int {
	var scores []int

	for i := 0; i < len(players); i++ {
		scores = append(scores, 0)
	}

	var rolls = 0

	for {
		for _, player := range players {
			move := 0

			for i := 0; i < 3; i++ {
				rolls++
				move += rolls
			}

			positions[player] = (positions[player] + move) % 10

			scores[player] += int(positions[player]) + 1

			if scores[player] >= 1000 {
				if player == 0 {
					return rolls * scores[1]
				} else {
					return rolls * scores[0]
				}
			}
		}
	}
}

var cache map[[5]int][2]int64

func compute(positions [2]int, scores [2]int, currentPlayerIndex int) [2]int64 {
	if scores[0] >= 21 {
		return [2]int64{1, 0}
	}

	if scores[1] >= 21 {
		return [2]int64{0, 1}
	}

	// Use a cache to help speed it up
	// Generated a unique cache key based off the data we have
	// First check if we've seen this "play" before doing some work
	cacheKey := [5]int{positions[0], positions[1], scores[0], scores[1], currentPlayerIndex}
	if cachedWin, exists := cache[cacheKey]; exists {
		return cachedWin
	}

	var wins [2]int64

	for roll1 := 1; roll1 <= 3; roll1++ {
		for roll2 := 1; roll2 <= 3; roll2++ {
			for roll3 := 1; roll3 <= 3; roll3++ {

				newPositions := positions
				newPositions[currentPlayerIndex] = (positions[currentPlayerIndex] + roll1 + roll2 + roll3) % 10

				newScores := scores
				newScores[currentPlayerIndex] = scores[currentPlayerIndex] + newPositions[currentPlayerIndex] + 1

				otherPlayerIndex := 1
				if currentPlayerIndex == 1 {
					otherPlayerIndex = 0
				}

				newWins := compute(newPositions, newScores, otherPlayerIndex)

				wins[0] = wins[0] + newWins[0]
				wins[1] = wins[1] + newWins[1]
			}
		}
	}

	cache[cacheKey] = wins
	return wins
}

func part2(positions [2]int) int64 {
	// Init cache, TIL: "panic: assignment to entry in nil map" happens when you don't...
	cache = make(map[[5]int][2]int64)

	scores := [2]int{0, 0}

	results := compute(positions, scores, 0)
	fmt.Printf("Player 1 Wins in %d universes\n", results[0])
	fmt.Printf("Player 2 wins in %d universes\n", results[1])

	mostWins := funk.MaxInt64([]int64{results[0], results[1]})
	fmt.Printf("Most Player wins is %d\n", mostWins)

	return mostWins
}

func main() {
	// input := strings.Split(example, "\n")
	input, _ := readFileLines("input")

	fmt.Printf("Input: (%d)\n%+v\n", len(input), input)

	players, positions := parsePlayersAndStartingPositions(input)

	fmt.Printf("Players:\n%+v\n", players)
	fmt.Printf("Positions:\n%+v\n", positions)

	part1Result := part1(players, positions)

	fmt.Printf("Part 1 = %d\n", part1Result)

	players, positions = parsePlayersAndStartingPositions(input)

	part2Result := part2(positions)

	fmt.Printf("Part 2 = %d\n", part2Result)
}
