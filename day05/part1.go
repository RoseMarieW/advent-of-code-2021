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
	from,to:=readtext()
	part1 := part1(from, to)
	fmt.Println(part1)
}

func part1(from []string, to []string) int {
	var ocean [1000][1000]int
	//fond scary spots
	for i:=0;i<len(from);i=i+2{
		if (from[i] == to[i]) || (from[i+1]==to[i+1]){
			ocean = draw(ocean,convert(from[i]),convert(from[i+1]),convert(to[i]),convert(to[i+1]))
		}
	}
	//count scary spots
	nopeZone :=0
	for i:=0;i<1000;i++{
		for j:=0;j<1000;j++{
			if ocean[i][j]>=2{
				nopeZone++
			}
		}
	}
	return nopeZone
}

func draw(ocean [1000][1000]int,fromX int,fromY int,toX int,toY int) [1000][1000]int{
	if(fromX==toX){
		if fromY < toY{
			for i:=fromY;i<=toY;i++{
				ocean[fromX][i]++
			}
		}else{
			for i:=toY;i<=fromY;i++{
				ocean[fromX][i]++
			}	
		}
	}else{
		if fromX < toX{
			for i:=fromX;i<=toX;i++{
				ocean[i][fromY]++
			}
		}else{
			for i:=toX;i<=fromX;i++{
				ocean[i][fromY]++
			}	
		}
	}
	return ocean
}

func readtext() ([]string,[]string) {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)

	
	var from []string
	var to []string
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		//ew I dont like this
		from = append(from,strings.Split(split[0],",")[0],strings.Split(split[0],",")[1])
		to = append(to,strings.Split(split[2],",")[0],strings.Split(split[2],",")[1])
	}
	return from,to
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
