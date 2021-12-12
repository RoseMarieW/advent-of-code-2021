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
	var digits [10]int
	for i:=11;i<len(text);i+=15{
		digits = digitValue(text,digits,i)
		digits = digitValue(text,digits,i+1)
		digits = digitValue(text,digits,i+2)
		digits = digitValue(text,digits,i+3)
	}

	count :=0
	for i:=0;i<len(digits);i++{
		count += digits[i]
	}
	return count
}

func digitValue(text []string, digits [10]int, index int)[10]int{
	if (len(text[index])== 2){
		digits[1] ++
	}else if (len(text[index])== 4){
		digits[4] ++
	}else if (len(text[index])== 3){
		digits[7] ++
	}else if (len(text[index])== 7){
		digits[8] ++
	}
	return digits
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