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
}

func part1(octoFloor [10][10]int) int {
	total := 0
	for step := 0; step < 100; step++ {
		//the energy of each octo inc by 1
		octoFloor = energyInc(octoFloor)

		//octo flashes
		for true {
			//is there any flashes?
			result, row, col := isFlash(octoFloor)
			//if not then dont spread the flashing
			if !result {
				break
			}

			//if there is then increase energy of surrounding octo
			//top left corner checks right and below
			if row == 0 && col == 0 {
				octoFloor[row][col+1] = increase(octoFloor[row][col+1])
				octoFloor[row+1][col+1] = increase(octoFloor[row+1][col+1])
				octoFloor[row+1][col] = increase(octoFloor[row+1][col])
			} else if row == 0 && col == 9 {
				//top right corner
				octoFloor[row][col-1] = increase(octoFloor[row][col-1])
				octoFloor[row+1][col-1] = increase(octoFloor[row+1][col-1])
				octoFloor[row+1][col] = increase(octoFloor[row+1][col])
			} else if row == 0 {
				//top row
				octoFloor[row][col-1] = increase(octoFloor[row][col-1])
				octoFloor[row+1][col-1] = increase(octoFloor[row+1][col-1])
				octoFloor[row+1][col] = increase(octoFloor[row+1][col])
				octoFloor[row][col+1] = increase(octoFloor[row][col+1])
				octoFloor[row+1][col+1] = increase(octoFloor[row+1][col+1])
			} else if row == 9 && col == 0 {
				//bottom left corner
				octoFloor[row-1][col] = increase(octoFloor[row-1][col])
				octoFloor[row-1][col+1] = increase(octoFloor[row-1][col+1])
				octoFloor[row][col+1] = increase(octoFloor[row][col+1])
			} else if col == 0 {
				//left side
				octoFloor[row-1][col] = increase(octoFloor[row-1][col])
				octoFloor[row-1][col+1] = increase(octoFloor[row-1][col+1])
				octoFloor[row][col+1] = increase(octoFloor[row][col+1])
				octoFloor[row+1][col] = increase(octoFloor[row+1][col])
				octoFloor[row+1][col+1] = increase(octoFloor[row+1][col+1])
			} else if col == 9 && row == 9 {
				//bottom right corner
				octoFloor[row-1][col] = increase(octoFloor[row-1][col])
				octoFloor[row-1][col-1] = increase(octoFloor[row-1][col-1])
				octoFloor[row][col-1] = increase(octoFloor[row][col-1])
			} else if row == 9 {
				//bottom side
				octoFloor[row-1][col] = increase(octoFloor[row-1][col])
				octoFloor[row-1][col+1] = increase(octoFloor[row-1][col+1])
				octoFloor[row][col+1] = increase(octoFloor[row][col+1])
				octoFloor[row-1][col-1] = increase(octoFloor[row-1][col-1])
				octoFloor[row][col-1] = increase(octoFloor[row][col-1])
			} else if col == 9 {
				//right side
				octoFloor[row-1][col] = increase(octoFloor[row-1][col])
				octoFloor[row-1][col-1] = increase(octoFloor[row-1][col-1])
				octoFloor[row][col-1] = increase(octoFloor[row][col-1])
				octoFloor[row+1][col-1] = increase(octoFloor[row+1][col-1])
				octoFloor[row+1][col] = increase(octoFloor[row+1][col])
			} else {
				octoFloor[row-1][col] = increase(octoFloor[row-1][col])
				octoFloor[row-1][col+1] = increase(octoFloor[row-1][col+1])
				octoFloor[row][col+1] = increase(octoFloor[row][col+1])
				octoFloor[row+1][col] = increase(octoFloor[row+1][col])
				octoFloor[row+1][col+1] = increase(octoFloor[row+1][col+1])
				octoFloor[row][col-1] = increase(octoFloor[row][col-1])
				octoFloor[row+1][col-1] = increase(octoFloor[row+1][col-1])
				octoFloor[row-1][col-1] = increase(octoFloor[row-1][col-1])
			}
			octoFloor[row][col] = -1
			total++
		}

		//reset floor and take another step
		octoFloor = reset(octoFloor)
	}
	return total
}

func reset(octoFloor [10][10]int) [10][10]int {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if octoFloor[row][col] == -1 {
				octoFloor[row][col] = 0
			}
		}
	}
	return octoFloor
}
func increase(octo int) int {
	if octo == -1 {
		return octo
	} else {
		return octo + 1
	}
}

func isFlash(octoFloor [10][10]int) (bool, int, int) {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if octoFloor[row][col] > 9 {
				return true, row, col
			}
		}
	}
	return false, -1, -1
}

func energyInc(octoFloor [10][10]int) [10][10]int {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			octoFloor[row][col] += 1
		}
	}
	return octoFloor
}

func readFile() [10][10]int {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	var text [10][10]int

	row := 0
	for scanner.Scan() {
		for col, char := range scanner.Text() {
			str := string(char)
			energyLevel := convert(str)

			text[row][col] = energyLevel
		}
		row++
	}
	return text
}

func convert(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		//handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}
