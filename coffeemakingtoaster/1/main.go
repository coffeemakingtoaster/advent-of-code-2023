package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

func main() {
	// Ignore error
	input, _ := readLines()
	fmt.Println(input[0])

	sum := 0
	re := regexp.MustCompile("[0-9]")

	for _, v := range input {
		first_digit := re.FindString(v)
		second_digit := re.FindString(reverse(v))
		n, _ := strconv.Atoi(first_digit + second_digit)
		fmt.Println(sum)
		sum = sum + n
	}
	fmt.Println(sum)
}
