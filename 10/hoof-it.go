package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	y, x int
}

func main() {
	hikingMap := getHikingMap()
	sum1, sum2 := sumTrailheadScore(hikingMap)
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func sumTrailheadScore(hikingMap [][]int) (sum1, sum2 int) {
	starts := findStarts(hikingMap)
	for _, start := range starts {
		score1, score2 := getTrailheadScore(start, hikingMap)
		sum1 += score1
		sum2 += score2
	}
	return
}

func getTrailheadScore(start Position, hikingMap [][]int) (int, int) {
	moveStack := []Position{start}
	heightsReachedSet := map[Position]bool{}
	heightsReachedList := []Position{}
	for len(moveStack) > 0 {
		var move Position
		move, moveStack = moveStack[len(moveStack)-1], moveStack[:len(moveStack)-1] // pop
		if getValueForPosition(move, hikingMap) == 9 {
			heightsReachedSet[move] = true
			heightsReachedList = append(heightsReachedList, move)
		} else {
			moveStack = append(moveStack, getPossibleNextMoves(move, hikingMap)...)
		}
	}
	return len(heightsReachedSet), len(heightsReachedList)
}

func getPossibleNextMoves(move Position, hikingMap [][]int) (nextMoves []Position) {
	currentPositionValue := getValueForPosition(move, hikingMap)
	//up
	if move.y > 0 {
		nextMove := Position{y: move.y - 1, x: move.x}
		if currentPositionValue+1 == getValueForPosition(nextMove, hikingMap) {
			nextMoves = append(nextMoves, nextMove)
		}
	}
	//right
	if move.x < len(hikingMap[0])-1 {
		nextMove := Position{y: move.y, x: move.x + 1}
		if currentPositionValue+1 == getValueForPosition(nextMove, hikingMap) {
			nextMoves = append(nextMoves, nextMove)
		}
	}
	//down
	if move.y < len(hikingMap)-1 {
		nextMove := Position{y: move.y + 1, x: move.x}
		if currentPositionValue+1 == getValueForPosition(nextMove, hikingMap) {
			nextMoves = append(nextMoves, nextMove)
		}
	}
	//left
	if move.x > 0 {
		nextMove := Position{y: move.y, x: move.x - 1}
		if currentPositionValue+1 == getValueForPosition(nextMove, hikingMap) {
			nextMoves = append(nextMoves, nextMove)
		}
	}
	return
}

func findStarts(hikingMap [][]int) (starts []Position) {
	for y := range hikingMap {
		for x := range hikingMap[y] {
			if hikingMap[y][x] == 0 {
				starts = append(starts, Position{y: y, x: x})
			}
		}
	}
	return
}

func getValueForPosition(position Position, hikingMap [][]int) int {
	return hikingMap[position.y][position.x]
}

func getHikingMap() (hikingMap [][]int) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		y := strings.Split(line, "")
		yInt := []int{}
		for _, value := range y {
			valueInt, _ := strconv.Atoi(value)
			yInt = append(yInt, valueInt)
		}
		hikingMap = append(hikingMap, yInt)
	}
	return
}
