package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
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

func get_val(prev_line string, line string, next_line string, start_index int, end_index int) int {
	buffer := string(line[max(start_index-1, 0)])
	buffer += string(line[max(start_index-1, 0):min(end_index+1, len(line))])
	buffer += string(prev_line[max(start_index-1, 0):min(end_index+1, len(line))])
	buffer += string(next_line[max(start_index-1, 0):min(end_index+1, len(line))])
	if end_index == len(line)-1 {
		end_index = len(line)
	}
	// Hacky workaround for edge case
	// TODO: Find a better way to solve this
	num, _ := strconv.Atoi(strings.ReplaceAll((line[start_index:end_index]), ".", ""))
	if start_index == end_index {
		num, _ = strconv.Atoi(string(line[start_index]))
	}
	if !strings.ContainsAny(buffer, "* | / | = | $ | @ | # | % | + | - | &") {
		return 0
	}
	return num
}

func main() {
	input, _ := readLines()
	sum := 0
	for i, v := range input {
		line := string(v)
		prev_line := input[max(0, i-1)]
		next_line := input[min(len(input)-1, i+1)]

		is_in_number := false
		start_index := 0

		for i, v := range line {
			// No number (i.e. number end) or end of line
			if !unicode.IsDigit(v) || (i == len(line)-1 && is_in_number) {
				if is_in_number {
					is_in_number = false
					sum += get_val(prev_line, line, next_line, start_index, i)
				}
				continue
			}
			// Is in number? Continue till end
			if is_in_number {
				continue
			}

			is_in_number = true
			start_index = i

			if i == len(line)-1 && is_in_number {
				sum += get_val(prev_line, line, next_line, start_index, i)
			}
		}
	}
	fmt.Println(sum)
}
