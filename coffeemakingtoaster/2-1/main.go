package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeValues struct {
	Red   int
	Blue  int
	Green int
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
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

func (c *CubeValues) UpdateMin(a int, key string) {
	switch key {
	case "green":
		c.Green = getMax(c.Green, a)
		break
	case "red":
		c.Red = getMax(c.Red, a)
		break
	default:
		c.Blue = getMax(c.Blue, a)
	}
}

func getMinValues(input []string) CubeValues {
	// Set all to 0
	minValues := CubeValues{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
	for _, v := range input {
		split := strings.Split(v, " ")
		num, _ := strconv.Atoi(split[1])
		color := split[2]
		minValues.UpdateMin(num, color)
	}
	return minValues
}

func main() {
	input, _ := readLines()
	result := 0

	for _, v := range input {
		a := strings.Split(v, ":")[1]

		a = strings.ReplaceAll(a, ";", ",")

		values := strings.Split(a, ",")

		minValues := getMinValues(values)
		result += (minValues.Blue * minValues.Green * minValues.Red)
	}

	fmt.Println(result)
}
