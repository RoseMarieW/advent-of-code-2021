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
	part1:=part1(text)
	fmt.Println(part1)
}

func part1(text []string) int{
	riskLevel :=0
	for i:=0;i<len(text);i++{
		for j:=0;j<len(text[i]);j++{
			//top left corner checks right and below
			if i==0 && j==0{
				if text[i][j] < text[i][j+1] && text[i][j] < text[i+1][j]{
					riskLevel += convert(string(text[i][j]))+1
				}
			}else if i==0 && j==len(text[i])-1{
				//top right corner
				if text[i][j] < text[i][j-1] && text[i][j] < text[i+1][j]{
					riskLevel += convert(string(text[i][j]))+1
				}
			}else if i==0{
				//top row
				if text[i][j] < text[i][j-1] && text[i][j] < text[i+1][j] && text[i][j] < text[i][j+1]{
					riskLevel += convert(string(text[i][j]))+1
				}
			}else if i==len(text)-1 && j==0{
				//bottom left corner
				if text[i][j] < text[i-1][j] && text[i][j] < text[i][j+1]{
					riskLevel += convert(string(text[i][j]))+1
				}
			}else if j==0{
				//left side
				if text[i][j] < text[i-1][j] && text[i][j] < text[i+1][j] && text[i][j] < text[i][j+1]{
					riskLevel += convert(string(text[i][j]))+1
				}
			}else if j==len(text[i])-1 && i==len(text)-1{
				//bottom right corner
				if text[i][j] < text[i-1][j] && text[i][j] < text[i][j-1]{
					riskLevel += convert(string(text[i][j]))+1
				}
			}else if i==len(text)-1{
				//bottom side
				if text[i][j] < text[i-1][j] && text[i][j] < text[i][j-1] && text[i][j] < text[i][j+1]{
					riskLevel += convert(string(text[i][j]))+1
				}
			}else if j==len(text[i])-1{
				//right side
				if text[i][j] < text[i-1][j] && text[i][j] < text[i][j-1] && text[i][j] < text[i+1][j]{
					riskLevel += convert(string(text[i][j]))+1
				}
			}else{
				if text[i][j] < text[i-1][j] && text[i][j] < text[i][j-1] && text[i][j] < text[i+1][j] && text[i][j] < text[i][j+1]{
					riskLevel += convert(string(text[i][j]))+1
				}
			}
		}
		
	}
	return riskLevel
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

func convert(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}