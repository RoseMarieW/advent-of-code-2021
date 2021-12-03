package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	 "github.com/davecgh/go-spew/spew"
)

func main() {
	text := readFile()
	part1 := part1(text)
	fmt.Println(part1)
	part2 := part2(text)
	fmt.Println(part2)
}

func part1(text []int) int {
	numIncreases := 0
	for i := 1; i < len(text); i++ {
		if text[i] > text[i-1] {
			numIncreases++
			spew.Dump(numIncreases)
		}
	}
	return numIncreases
}

func part2(text []int) int {
	numIncreases := 0
	for i := 2; i < len(text)-1; i++ {
		if (text[i] + text[i+1] + text[i-1]) > (text[i] + text[i-1] + text[i-2]) {
			numIncreases++
		}
	}
	return numIncreases
}

func readFile() []int {
	file, err := os.Open("testinput.txt")

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
