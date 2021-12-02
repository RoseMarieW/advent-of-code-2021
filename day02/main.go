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
	part2 := part2(text)
	fmt.Println(part2)
}

func part1(text []string) int {
	horizontal := 0
	depth := 0

	for i := 0; i < len(text)-1; i = i + 2 {
		x := convert(text[i+1])
		if text[i] == "forward" {
			horizontal += x
		} else if text[i] == "down" {
			depth += x
		} else {
			depth -= x
		}
	}

	return depth * horizontal
}

func part2(text []string) int {
	horizontal := 0
	depth := 0
	aim := 0

	for i := 0; i < len(text)-1; i = i + 2 {
		x := convert(text[i+1])
		if text[i] == "forward" {
			horizontal += x
			depth += (x * aim)
		} else if text[i] == "down" {
			aim += x
		} else {
			aim -= x
		}
	}

	return depth * horizontal
}

func readFile() []string {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
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
