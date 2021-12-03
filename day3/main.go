package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(path string) ([]string, error) {
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

func part1(lines []string) (int64, int64, int64, error) {
	numberOfLines := len(lines)
	numberOfBitsPerLine := len(lines[0])

	fmt.Printf("number of lines %d\n", numberOfLines)
	fmt.Printf("per line %d\n\n", numberOfBitsPerLine)

	gamma := ""
	epsilon := ""

	for i := 0; i < numberOfBitsPerLine; i++ {
		on := 0
		off := 0
		for ii := 0; ii < numberOfLines; ii++ {
			bit, _ := strconv.Atoi(string(string(lines[ii])[i]))

			if bit == 1 {
				on++
			} else {
				off++
			}
		}

		fmt.Printf("column %d = %d vs %d\n", i, on, off)

		if on > off {
			gamma += "1"
		} else {
			gamma += "0"
		}

		if on < off {
			epsilon += "1"
		} else {
			epsilon += "0"
		}

		fmt.Printf("column %d = gamma = %s = eplison = %s\n", i, gamma, epsilon)
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("Gamma = %d\n", gammaInt)
	fmt.Printf("Epsilon = %d\n", epsilonInt)
	fmt.Printf("Power = %d\n", gammaInt*epsilonInt)

	return gammaInt, epsilonInt, gammaInt * epsilonInt, nil
}

func part2(lines []string) (int64, int64, int64, error) {
	numberOfBitsPerLine := len(lines[0])

	o2Lines := lines
	co2Lines := lines

	for i := 0; i < numberOfBitsPerLine; i++ {
		// recalculate each iteration as we're changing the array length
		numberOfLines := len(o2Lines)

		if numberOfLines == 1 {
			continue
		}

		on := 0
		off := 0
		for ii := 0; ii < numberOfLines; ii++ {
			bit, _ := strconv.Atoi(string(string(o2Lines[ii])[i]))

			if bit == 1 {
				on++
			} else {
				off++
			}
		}

		keep := 0

		if on >= off {
			keep = 1
		}

		var newO2Lines []string

		for ii := 0; ii < numberOfLines; ii++ {
			bit, _ := strconv.Atoi(string(string(o2Lines[ii])[i]))

			if bit == keep {
				newO2Lines = append(newO2Lines, o2Lines[ii])
			}
		}

		o2Lines = newO2Lines
	}

	for i := 0; i < numberOfBitsPerLine; i++ {
		// recalculate each iteration as we're changing the array length
		numberOfLines := len(co2Lines)

		if numberOfLines == 1 {
			continue
		}

		on := 0
		off := 0
		for ii := 0; ii < numberOfLines; ii++ {
			bit, _ := strconv.Atoi(string(string(co2Lines[ii])[i]))

			if bit == 1 {
				on++
			} else {
				off++
			}
		}

		keep := 1

		if on >= off {
			keep = 0
		}

		var newCO2Lines []string

		for ii := 0; ii < numberOfLines; ii++ {
			bit, _ := strconv.Atoi(string(string(co2Lines[ii])[i]))

			if bit == keep {
				newCO2Lines = append(newCO2Lines, co2Lines[ii])
			}
		}

		co2Lines = newCO2Lines
	}

	if len(o2Lines) > 1|len(co2Lines) {
		return 0, 0, 0, os.ErrInvalid
	}

	fmt.Printf("02 = %s\n", o2Lines[0])
	o2, _ := strconv.ParseInt(o2Lines[0], 2, 64)

	fmt.Printf("CO2 = %s\n", co2Lines[0])
	co2, _ := strconv.ParseInt(co2Lines[0], 2, 64)

	fmt.Printf("O2 = %d\n", o2)
	fmt.Printf("CO2 = %d\n", co2)
	fmt.Printf("Power = %d\n", o2*co2)

	return o2, co2, o2 * co2, nil
}

func main() {
	lines, _ := readInput("./day3/input")

	// part1(lines)
	part2(lines)
}
