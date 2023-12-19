package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	dat, e := os.ReadFile("./input.txt")
	if e != nil {
		panic(e)
	}
	lines := strings.Split(string(dat), "\n")

	var result int

	// loop over each line
	for _, line := range lines {
		sum := 0
		// find first digit, start from first char
		for i := 0; i <= len(line)-1; i++ {
			if unicode.IsDigit(rune(line[i])) {
				sum += int(line[i]-'0') * 10 // multiply by 10 because first digit
				break                        // stop loop if found
			}
		}

		// find last digit, start from last char
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				sum += int(line[i] - '0')
				break // stop loop if found
			}
		}
		// append to result
		result += sum
	}
	fmt.Println(result)
}
