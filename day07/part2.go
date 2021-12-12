package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	text := readFile()
	part1:=part1(text)
	fmt.Println(part1)
}

func part1(text []int) int{
	minFuel := float64(99999999)
	//figure out possible positions
	min,max := minMax(text)
	//itterate through those positions
	for i:=min;i<=max;i++{
		//calculate total fuel cost for each position
		fuel := fuelCost(text, i)
		if fuel < minFuel{
			minFuel=fuel
		}
	}	
	return int(minFuel)
}

func fuelCost(positions []int,explode int) float64{
	var fuelCost float64
	for _,position := range positions {
		distance := float64(abs(position,explode))
		fuelCost += (math.Pow(distance, 2) + distance) / 2
	}
	return fuelCost
}

func abs(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }

func minMax(positions []int) (int, int) {
    var max int = positions[0]
    var min int = positions[0]
    for _, value := range positions {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}

func readFile() []int {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	var text []string

	for scanner.Scan() {
		text = strings.Split(scanner.Text(),",")
	}

	var positions []int
	for _, a := range text{
        positions = append(positions,convert(a))
    }
	return positions
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