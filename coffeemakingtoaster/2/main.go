package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MaxCubes struct {
	Red   int
	Blue  int
	Green int
}

func readLines() ([]string, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func isPossible(input string, maxValues MaxCubes) bool {
	v := strings.Split(input, " ")
	// Skip first because input starts with a blank space
	num, _ := strconv.Atoi(v[1])
	color := v[2]
	if color == "red" {
		return num <= maxValues.Red
	}
	if color == "green" {
		return num <= maxValues.Green
	}
	return num <= maxValues.Blue
}

func main() {
	lines, _ := readLines()
	sum := 0
	maxCubes := MaxCubes{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	for _, v := range lines {
		input := strings.Split(v, ":")
		game_number, _ := strconv.Atoi(strings.Split(input[0], " ")[1])

		cubeValues := strings.ReplaceAll(input[1], ";", ",")

		data := strings.Split(cubeValues, ",")
		hasError := false
		for _, a := range data {
			if !isPossible(a, maxCubes) {
				hasError = true
				break
			}
		}
		if hasError {
			continue
		}
		sum += game_number
	}
	fmt.Println(sum)
}
