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
	board := readBoard()
	bingoCalls := readBingoNums()
	part2 := part2(board, bingoCalls)
	fmt.Println(part2)
}

func part2(board []string, bingoCalls []string) int {
	//assume that it might take all bingoCalls
	maxCalls := 0
	winScore := 0
	//a board is 25
	for i:=0; i<len(board); i += 25{
		calls,score := contains(board[i:i+25],bingoCalls,maxCalls)
		//I dont think I need to check for max again because I check in contains (thats a bad name)
		if(calls != -1){
			maxCalls = calls
			winScore = score
		}
	}
	return winScore
}

func contains(board []string, bingoCalls []string, maxCalls int) (int,int){
	//for each call check if it is on board
	var tempBoard [25]int
	for i:=0;i<len(bingoCalls);i++{
		for j:=0;j<len(board);j++{
			//if contains
			if board[j] == bingoCalls[i]{
				//then that position is marked
				tempBoard[j] = 1
				//cant be a winner before 5 calls
				if i>=5{
					//if there is a winner
					if isWinner(tempBoard){
						//check if this winner is greater than any other maximum calls
						if i>maxCalls{
							maxCalls = i
							score := getScore(board,tempBoard,convert(bingoCalls[i]))
							return maxCalls, score
						}else{return -1,-1} 
					}
				}
			}
		}
		
	}
	return -1,-1
}

func isWinner(tempBoard [25]int) bool{
	//horizontals
	for i:=0;i<25;i+=5{
		if (tempBoard[i] + tempBoard[i+1] + tempBoard[i+2] + tempBoard[i+3] + tempBoard[i+4])==5{
			return true
		}
	}

	//verticals
	for i:=0;i<5;i++{
		if (tempBoard[i] + tempBoard[i+5] + tempBoard[i+10] + tempBoard[i+15] + tempBoard[i+20])==5{
			return true
		} 
	}
	return false
}

func getScore(board []string,tempBoard [25]int,winningNumber int)int{
	sum :=0
	for i:=0;i<len(board);i++{
		if tempBoard[i]==0{
			sum += convert(board[i])
		}
	}
	return sum*winningNumber
}

func readBoard() []string {
	file, err := os.Open("inputboard.txt")

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

func readBingoNums() []string {
	file, err := os.Open("inputnumbers.txt")

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
