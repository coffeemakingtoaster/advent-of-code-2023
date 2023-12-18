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

func wordToInt(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	default:
		return "9"
	}
}

func main() {
	// Ignore error
	input, _ := readLines()
	fmt.Println(input[0])

	sum := 0
	re := regexp.MustCompile("([0-9]|(one|two|three|four|five|six|seven|eight|nine))")
	rev_re := regexp.MustCompile("([0-9]|(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin))")

	for _, v := range input {
		first_digit := re.FindString(v)
		if len(first_digit) > 1 {
			first_digit = wordToInt(first_digit)
		}
		second_digit := reverse(rev_re.FindString(reverse(v)))
		if len(second_digit) > 1 {
			second_digit = wordToInt(second_digit)
		}
		n, _ := strconv.Atoi(first_digit + second_digit)
		sum = sum + n
	}
	fmt.Println(sum)
}
