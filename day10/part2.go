package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	text := readFile()
	part1 := part1(text)
	fmt.Println(part1)
}

func part1(text []string) int {
	score := 0
	for i := 0; i < len(text); i++ {
		score += checkLine(text[i])
	}
	return score
}

func checkLine(line string) int {
	//hint was helpful thank you
	var openBracketStack []string

	for _, char := range line {
		bracket := string(char)
		stackIndex := len(openBracketStack) - 1
		//if opening bracket push it onto stack
		if bracket == "(" || bracket == "[" || bracket == "{" || bracket == "<" {
			openBracketStack = append(openBracketStack, bracket)
		} else {
			//its a closing bracket
			if bracket == ")" {
				if openBracketStack[stackIndex] == "(" {
					//yay pop it
					openBracketStack = openBracketStack[:stackIndex]
				} else {
					//this is the mismatch
					return 3
				}
			}
			if bracket == "]" {
				if openBracketStack[stackIndex] == "[" {
					//yay pop it
					openBracketStack = openBracketStack[:stackIndex]
				} else {
					//this is the mismatch
					return 57
				}
			}
			if bracket == "}" {
				if openBracketStack[stackIndex] == "{" {
					//yay pop it
					openBracketStack = openBracketStack[:stackIndex]
				} else {
					//this is the mismatch
					return 1197
				}
			}
			if bracket == ">" {
				if openBracketStack[stackIndex] == "<" {
					//yay pop it
					openBracketStack = openBracketStack[:stackIndex]
				} else {
					//this is the mismatch
					return 25137
				}
			}
		}
	}
	return 0
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

//func convert(str string) int {
//	i, err := strconv.Atoi(str)
//	if err != nil {
//		// handle error
//		fmt.Println(err)
//		os.Exit(2)
//	}
//	return i
//}
