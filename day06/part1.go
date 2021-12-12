package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	text := readFile()
	part1:=part1(text)
	fmt.Println(part1)
}

func part1(text []string) int{
	//count fishies at each stage of life
	var age [9]int
	age = countFishAge(text,age)
	//itterate through the days to multiply in fish
	for i:=1;i<=256;i++{
		age = makeBabies(age)
	}
	count := 0
	for i:=0;i<9;i++{
		count += age[i]
	}
	return count
}

func makeBabies(age [9]int) [9]int{
	temp := age[0]
	age[0]=0
	for i:=1;i<9;i++{
		age[i-1]=age[i]
	}
	age[6]+=temp
	age[8]=temp

	return age
}

func countFishAge(fish []string,age [9]int) [9]int{
	for i:=0;i<len(fish);i++{
		age[convert(fish[i])] ++
	}
	return age
}

func readFile() []string {
	file, err := os.Open("testinput.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	var text []string

	for scanner.Scan() {
		text = strings.Split(scanner.Text(),",")
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