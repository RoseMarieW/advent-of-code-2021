package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	text := readFile()
	part2 := part2(text)
	fmt.Println(part2)
}

func part2(text []string) int {
	var scores []int
	for i := 0; i < len(text); i++ {
		score := checkLine(text[i])
		if score != -1 {
			scores = append(scores, score)
		}
	}
	return median(scores)
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
					return -1
				}
			}
			if bracket == "]" {
				if openBracketStack[stackIndex] == "[" {
					//yay pop it
					openBracketStack = openBracketStack[:stackIndex]
				} else {
					//this is the mismatch
					return -1
				}
			}
			if bracket == "}" {
				if openBracketStack[stackIndex] == "{" {
					//yay pop it
					openBracketStack = openBracketStack[:stackIndex]
				} else {
					//this is the mismatch
					return -1
				}
			}
			if bracket == ">" {
				if openBracketStack[stackIndex] == "<" {
					//yay pop it
					openBracketStack = openBracketStack[:stackIndex]
				} else {
					//this is the mismatch
					return -1
				}
			}
		}
	}
	//calculate based on what is left in the stack
	return calcCompletionScore(openBracketStack)
}

func median(scores []int) int {
	sort.Ints(scores)

	mNumber := len(scores) / 2

	return scores[mNumber]
}

func calcCompletionScore(openBracketStack []string) int {
	score := 0

	for i := len(openBracketStack) - 1; i >= 0; i-- {
		score *= 5

		if openBracketStack[i] == "(" {
			score += 1
		} else if openBracketStack[i] == "[" {
			score += 2
		} else if openBracketStack[i] == "{" {
			score += 3
		} else if openBracketStack[i] == "<" {
			score += 4
		}
	}
	return score
}
func readFile() []string {
	file, err := os.Open("testinput.txt")

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
