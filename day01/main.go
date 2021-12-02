package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	text := readFile()
	part1 := part1(text)
	fmt.Println(part1)
}

func part1(text []int) int {
	numIncreases := 0
	for i := 1; i < len(text); i++ {
		if text[i] > text[i-1] {
			numIncreases++
		}
	}
	return numIncreases
}

func readFile() []int {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []int

	for scanner.Scan() {
		text = append(text, convert(scanner.Text()))
	}
	return text
}

func convert(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}
