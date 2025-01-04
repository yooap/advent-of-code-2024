package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	board := getWordSearchBoard()

	// part 1
	foundWordCount := 0
	for y := range board {
		for x, letter := range board[y] {
			if letter == 'X' {
				foundWordCount += searchForWordsAroundX(y, x, board)
			}
		}
	}
	fmt.Println(foundWordCount)

	// part 2
	foundWordCount = 0
	for y := range board {
		for x, letter := range board[y] {
			if letter == 'A' {
				if containsCrossMAS_AroundA(y, x, board) {
					foundWordCount++
				}
			}
		}
	}
	fmt.Println(foundWordCount)
}

func searchForWordsAroundX(y, x int, board [][]rune) (foundWordCount int) {
	foundWordCount += searchForWordsAroundX_Up(y, x, board)
	foundWordCount += searchForWordsAroundX_UpRight(y, x, board)
	foundWordCount += searchForWordsAroundX_Right(y, x, board)
	foundWordCount += searchForWordsAroundX_DownRight(y, x, board)
	foundWordCount += searchForWordsAroundX_Down(y, x, board)
	foundWordCount += searchForWordsAroundX_DownLeft(y, x, board)
	foundWordCount += searchForWordsAroundX_Left(y, x, board)
	foundWordCount += searchForWordsAroundX_UpLeft(y, x, board)
	return
}
func searchForWordsAroundX_Up(y, x int, board [][]rune) int {
	if y < 3 {
		return 0
	}

	for i, letter := range "MAS" {
		if board[y-1-i][x] != letter {
			return 0
		}
	}
	return 1
}
func searchForWordsAroundX_UpRight(y, x int, board [][]rune) int {
	if y < 3 || x > len(board[y])-4 {
		return 0
	}

	for i, letter := range "MAS" {
		if board[y-1-i][x+1+i] != letter {
			return 0
		}
	}
	return 1
}
func searchForWordsAroundX_Right(y, x int, board [][]rune) int {
	if x > len(board[y])-4 {
		return 0
	}

	for i, letter := range "MAS" {
		if board[y][x+1+i] != letter {
			return 0
		}
	}
	return 1
}
func searchForWordsAroundX_DownRight(y, x int, board [][]rune) int {
	if y > len(board)-4 || x > len(board[y])-4 {
		return 0
	}

	for i, letter := range "MAS" {
		if board[y+1+i][x+1+i] != letter {
			return 0
		}
	}
	return 1
}
func searchForWordsAroundX_Down(y, x int, board [][]rune) int {
	if y > len(board)-4 {
		return 0
	}

	for i, letter := range "MAS" {
		if board[y+1+i][x] != letter {
			return 0
		}
	}
	return 1
}
func searchForWordsAroundX_DownLeft(y, x int, board [][]rune) int {
	if y > len(board)-4 || x < 3 {
		return 0
	}

	for i, letter := range "MAS" {
		if board[y+1+i][x-1-i] != letter {
			return 0
		}
	}
	return 1
}
func searchForWordsAroundX_Left(y, x int, board [][]rune) int {
	if x < 3 {
		return 0
	}

	for i, letter := range "MAS" {
		if board[y][x-1-i] != letter {
			return 0
		}
	}
	return 1
}
func searchForWordsAroundX_UpLeft(y, x int, board [][]rune) int {
	if y < 3 || x < 3 {
		return 0
	}

	for i, letter := range "MAS" {
		if board[y-1-i][x-1-i] != letter {
			return 0
		}
	}
	return 1
}

func containsCrossMAS_AroundA(y, x int, board [][]rune) bool {
	if y == 0 || x == 0 || y == len(board)-1 || x == len(board[y])-1 {
		return false
	}

	return containsMAS_AroundA_UpRight(y, x, board) && containsMAS_AroundA_DownRight(y, x, board)
}
func containsMAS_AroundA_UpRight(y, x int, board [][]rune) bool {
	bottomLeftRune := board[y+1][x-1]
	topRightRune := board[y-1][x+1]
	return isM_orS(bottomLeftRune) && isM_orS(topRightRune) && bottomLeftRune != topRightRune
}
func containsMAS_AroundA_DownRight(y, x int, board [][]rune) bool {
	topLeftRune := board[y-1][x-1]
	bottomRightRune := board[y+1][x+1]
	return isM_orS(topLeftRune) && isM_orS(bottomRightRune) && topLeftRune != bottomRightRune
}
func isM_orS(r rune) bool {
	return r == 'M' || r == 'S'
}

func getWordSearchBoard() (board [][]rune) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		lineAsRunes := []rune(line)
		board = append(board, lineAsRunes)
	}
	return
}
