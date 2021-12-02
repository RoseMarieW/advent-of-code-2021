package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	text := readFile()
	part1 := part1(text)
	println(part1)
}

func part1(text []string) int {
	numIncreases := 0
	for i := 1; i < len(text); i++ {
		if text[i] > text[i-1] {
			numIncreases++
		}
	}
	return numIncreases
}

func readFile() []string {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}
