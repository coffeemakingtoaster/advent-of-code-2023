package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Print("Error occured")
	}
	fmt.Println(string(dat))
}
