package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	caves = readFile()
	part1 := part1("start", "end", []string{})
	fmt.Println(part1)
}

var caves map[string][]string

//thank you for recursion tip
func part1(from string, to string, paths []string) int {
	//base case reach end
	if from == to {
		completePath := append(paths, "end")
		fmt.Println(completePath)

		return 1
	}
	//calculate as we go
	total := 0
	for _, path := range caves[from] {
		//can only go through small caves once
		if pathInCaveSmall(path, paths) {
			continue
		}
		//recurse
		updatedPaths := append(paths, from)
		total += part1(path, to, updatedPaths)
	}

	return total
}

func pathInCaveSmall(cave string, paths []string) bool {
	for _, path := range paths {
		if path == cave && strings.ToLower(cave) == cave {
			return true
		}
	}
	return false
}

func readFile() map[string][]string {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	text := make(map[string][]string)

	for scanner.Scan() {
		var splitString = strings.Split(scanner.Text(), "-")

		if _, ok := text[splitString[0]]; ok {
			text[splitString[0]] = append(text[splitString[0]], splitString[1])
		} else {
			text[splitString[0]] = []string{splitString[1]}
		}

		if _, ok := text[splitString[1]]; ok {
			text[splitString[1]] = append(text[splitString[1]], splitString[0])
		} else {
			text[splitString[1]] = []string{splitString[0]}
		}
	}
	return text
}

//func convert(str string) int {
//	i, err := strconv.Atoi(str)
//	if err != nil {
//		//handle error
//		fmt.Println(err)
//		os.Exit(2)
//	}
//	return i
//}
