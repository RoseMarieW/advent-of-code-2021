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
	//part1 := part1(text)
	//fmt.Println(part1)
	part2 := part2(text)
	fmt.Println(part2)
}

func part2(text []string) string {
	//for those reading I did part1 a week ago, got sick, and now I hate how I did part1 sooo lets retry this
	
	//dont want to mess with original input
	oxygen := text
	oxygenRating := ""
	position :=0
	//Oxygen Generator Rating
	for len(oxygen) >1{
		//find maximum
		max := mostCommonValue(oxygen, position)
		//discard minimum
		oxygen = discard(oxygen, position, max)
		//add value to rating
		oxygenRating += max
		position++
	}

	//CO2 Scrubber Rating
	//dont want to mess with original input
	co := text
	coRating := ""
	position = 0
	//Oxygen Generator Rating
	for len(co) >1{
		//find min
		min := leastCommonValue(co, position)
		println(min)
		//discard min
		co = discard(co, position, min)
		//add value to rating
		coRating += min
		position++
	}

	println(coRating)
	return oxygenRating
}

func part1(text []string) int64 {
	var columnArray []int
	var split []string
	for bits := 0; bits < len(text); bits++{
		split = strings.Split(text[bits], "")
		for column := 0; column < len(split);column++{
			if bits==0{
				columnArray = append(columnArray,convert(split[column]))
			}else{columnArray[column] +=  convert(split[column])}
		}
	}

	gamma := ""
	epsilon := ""

	for i:=0;i<len(columnArray);i++{
		if columnArray[i]>len(text)/2{
			gamma += "1"
			epsilon += "0"
		}else{
			gamma += "0"
			epsilon += "1"
		}
	}

	var gammarate int64
	var epsilonrate int64

	if i, err := strconv.ParseInt(gamma, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		gammarate=i
	}
	
	if i, err := strconv.ParseInt(epsilon, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		epsilonrate=i
	}
	
	return gammarate*epsilonrate
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

func mostCommonValue(report []string, position int) string{
	one := 0
	zero := 0
	for _, diagnostic := range report {
		if diagnostic[position] == '1' {
			one++
		} else {
			zero++
		}
	}

	if one>=zero{
		return "1"
	} else {
		return "0"
	}

}

//this is redundant
func leastCommonValue(report []string, position int) string{
	one := 0
	zero := 0
	for _, diagnostic := range report {
		if diagnostic[position] == '1' {
			one++
		} else {
			zero++
		}
	}

	if one<zero{
		return "1"
	} else {
		return "0"
	}

}

func discard(report []string, position int, keep string) []string{
	var kept []string
	for index, diagnostic := range report {
		//this seems...
		if diagnostic[position] == []byte(keep)[0] {
			kept = append(kept,report[index])
		}
	}
	return kept
}