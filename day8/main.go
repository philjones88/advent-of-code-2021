package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

const (
	example = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
	edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
	fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
	fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
	aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
	fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
	dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
	bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
	egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
	gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`
)

// 0:      1:      2:      3:      4:
// aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
// ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
// gggg    ....    gggg    gggg    ....

//  5:      6:      7:      8:      9:
// aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
// dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
// gggg    gggg    ....    gggg    gggg

// segments =
// 0 = 6
// 1 = 2 (unique)
// 2 = 5
// 3 = 5
// 4 = 4 (unique)
// 5 = 5
// 6 = 6
// 7 = 3 (unique)
// 8 = 7 (unique)
// 9 = 6

func stringToInt(str string) int {
	parsed, err := strconv.Atoi(str)
	if err != nil {
		panic("Rut roh! Not a number")
	}
	return parsed
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

func part1(inputLines []string) int {
	counter := 0

	for _, inputLine := range inputLines {
		inputLineSplit := strings.Split(inputLine, " ")

		for _, segment := range inputLineSplit {
			if len(segment) == 2 || len(segment) == 4 || len(segment) == 3 || len(segment) == 7 {
				counter++
			}
		}
	}

	return counter
}

func part2(inputLines []string) int {
	totalSum := 0

	for _, inputLine := range inputLines {
		inputPartsSplit := strings.Split(inputLine, " | ")

		signalPatterns := strings.Fields(strings.TrimSpace(inputPartsSplit[0]))

		// Sort them
		for i, signalPattern := range signalPatterns {
			temp := strings.Split(signalPattern, "")
			sort.Strings(temp)
			signalPatterns[i] = strings.Join(temp, "")
		}

		outputPatterns := strings.Fields(strings.TrimSpace(inputPartsSplit[1]))

		for i, outputPattern := range outputPatterns {
			temp := strings.Split(outputPattern, "")
			sort.Strings(temp)
			outputPatterns[i] = strings.Join(temp, "")
		}

		signalPatternMap := make(map[int]string)

		// First find the easy ones, these help diff the harder numbers
		for _, signalPattern := range signalPatterns {
			if len(signalPattern) == 2 {
				signalPatternMap[1] = signalPattern
			}

			if len(signalPattern) == 4 {
				signalPatternMap[4] = signalPattern
			}

			if len(signalPattern) == 3 {
				signalPatternMap[7] = signalPattern
			}

			if len(signalPattern) == 7 {
				signalPatternMap[8] = signalPattern
			}
		}

		one := strings.Split(signalPatternMap[1], "")
		four := strings.Split(signalPatternMap[4], "")

		// 0 == 1 = 0
		// 6 == 1 = 1
		// 9 == 1 = 0

		// find 6
		for _, signalPattern := range signalPatterns {
			if len(signalPattern) == 6 {
				signalPatternSplit := strings.Split(signalPattern, "")
				if o, _ := funk.DifferenceString(one, signalPatternSplit); len(o) == 1 {
					signalPatternMap[6] = signalPattern
					break
				}
			}
		}

		// 0 == 4 = 1
		// 9 == 4 = 2

		// find 0
		for _, signalPattern := range signalPatterns {
			if len(signalPattern) == 6 {
				if signalPattern == signalPatternMap[6] {
					continue
				}

				signalPatternSplit := strings.Split(signalPattern, "")

				if o, _ := funk.DifferenceString(four, signalPatternSplit); len(o) == 1 {
					signalPatternMap[0] = signalPattern
					break
				}
			}
		}

		// find 9
		for _, signalPattern := range signalPatterns {
			if len(signalPattern) == 6 {
				if signalPattern == signalPatternMap[0] || signalPattern == signalPatternMap[6] {
					continue
				}
				signalPatternMap[9] = signalPattern
				break
			}
		}

		// 2 == 6 = 3
		// 3 == 6 = 3
		// 5 == 6 = 1

		// find 5
		six := strings.Split(signalPatternMap[6], "")
		for _, signalPattern := range signalPatterns {
			if len(signalPattern) == 5 {
				signalPatternSplit := strings.Split(signalPattern, "")

				if o, _ := funk.DifferenceString(six, signalPatternSplit); len(o) == 1 {
					signalPatternMap[5] = signalPattern
					break
				}
			}
		}

		// 3 == 5 = 2
		// 2 == 5 = 2

		// find 3
		five := strings.Split(signalPatternMap[5], "")
		for _, signalPattern := range signalPatterns {
			if len(signalPattern) == 5 {
				signalPatternSplit := strings.Split(signalPattern, "")

				if o, _ := funk.DifferenceString(five, signalPatternSplit); len(o) == 1 {
					signalPatternMap[3] = signalPattern
					break
				}
			}
		}

		// find 2
		for _, signalPattern := range signalPatterns {
			if len(signalPattern) == 5 {
				signalPatternSplit := strings.Split(signalPattern, "")

				if o, _ := funk.DifferenceString(five, signalPatternSplit); len(o) == 2 {
					signalPatternMap[2] = signalPattern
					break
				}
			}
		}

		rowSum := ""

		for _, outputPattern := range outputPatterns {
			temp := ""
			for signalKey, signalVal := range signalPatternMap {
				if outputPattern == signalVal {
					temp = temp + strconv.Itoa(signalKey)
				}
			}
			rowSum = rowSum + temp
		}

		totalSum += stringToInt(rowSum)
	}

	return totalSum
}

func main() {
	exampleLines, _ := readFileLines("./day8/input")
	var inputLines []string

	for _, inputLine := range exampleLines {
		inputLines = append(inputLines, strings.Split(inputLine, " | ")[1])
	}

	result := part1(inputLines)

	fmt.Printf("Part 1: %d\n", result)

	resultPart2 := part2(exampleLines)

	fmt.Printf("Part 2: %d\n", resultPart2)
}
