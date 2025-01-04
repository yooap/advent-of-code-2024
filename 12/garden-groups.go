package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type Position struct {
	y, x int
}

type Plot struct {
	positions []Position
	perimeter int
}

func main() {
	garden := getGarden()
	plots := []Plot{}
	allUsedPositions := map[Position]bool{}
	for y := range garden {
		for x := range garden[y] {
			position := Position{y: y, x: x}
			if _, ok := allUsedPositions[position]; ok {
				continue
			}
			plot := mapPlot(position, garden)
			plots = append(plots, plot)
			for _, plotPosition := range plot.positions {
				allUsedPositions[plotPosition] = true
			}
		}
	}

	costs := calculateCosts(plots)
	fmt.Println(costs)

	costs = calculateCosts2(plots)
	fmt.Println(costs)
}

func mapPlot(position Position, garden []string) Plot {
	value := garden[position.y][position.x]
	perimeter := 0
	handledPlotPositions := map[Position]bool{}
	moves := []Position{position}

	for len(moves) > 0 {
		move := moves[0]
		if _, ok := handledPlotPositions[move]; ok {
			moves = moves[1:]
			continue
		}
		// up
		if move.y > 0 && garden[move.y-1][move.x] == value {
			moves = append(moves, Position{y: move.y - 1, x: move.x})
		} else {
			perimeter++
		}
		// right
		if move.x < len(garden[0])-1 && garden[move.y][move.x+1] == value {
			moves = append(moves, Position{y: move.y, x: move.x + 1})
		} else {
			perimeter++
		}
		// down
		if move.y < len(garden)-1 && garden[move.y+1][move.x] == value {
			moves = append(moves, Position{y: move.y + 1, x: move.x})
		} else {
			perimeter++
		}
		// left
		if move.x > 0 && garden[move.y][move.x-1] == value {
			moves = append(moves, Position{y: move.y, x: move.x - 1})
		} else {
			perimeter++
		}
		handledPlotPositions[move] = true
		moves = moves[1:]
	}

	return Plot{positions: getKeys(handledPlotPositions), perimeter: perimeter}
}

func getKeys(_map map[Position]bool) (keys []Position) {
	for key := range _map {
		keys = append(keys, key)
	}
	return
}

func calculateCosts(plots []Plot) (costs int) {
	for _, plot := range plots {
		costs += plot.perimeter * len(plot.positions)
	}
	return
}

func calculateCosts2(plots []Plot) (costs int) {
	for _, plot := range plots {
		area := len(plot.positions)
		sides := countSides(plot.positions)
		costs += sides * area
	}
	return
}

func countSides(positions []Position) (sides int) {
	positionMap := map[Position]bool{}
	for _, position := range positions {
		positionMap[position] = true
	}

	// count horizontal sides
	sort.Slice(positions, func(a, b int) bool {
		aPos := positions[a]
		bPos := positions[b]
		if aPos.y < bPos.y || (aPos.y == bPos.y && aPos.x < bPos.x) {
			return true
		}
		return false
	})
	sides += countTopHorizontalSidesAfterSorting(positions, positionMap)
	slices.Reverse(positions)
	sides += countBottomHorizontalSidesAfterSorting(positions, positionMap)

	// count vertical sides
	sort.Slice(positions, func(a, b int) bool {
		aPos := positions[a]
		bPos := positions[b]
		if aPos.x < bPos.x || (aPos.x == bPos.x && aPos.y < bPos.y) {
			return true
		}
		return false
	})
	sides += countLeftVerticalSidesAfterSorting(positions, positionMap)
	slices.Reverse(positions)
	sides += countRightVerticalSidesAfterSorting(positions, positionMap)

	return
}

func countTopHorizontalSidesAfterSorting(positions []Position, positionMap map[Position]bool) (sides int) {
	sides++
	lastPosition := positions[0]
	sideBroken := false
	for _, position := range positions[1:] {
		_, topBlocked := positionMap[Position{y: position.y - 1, x: position.x}]
		if topBlocked {
			sideBroken = true
		} else {
			if lastPosition.y == position.y {
				if lastPosition.x+1 != position.x || sideBroken {
					// new side
					sides++
					sideBroken = false
				}
			} else {
				// new side
				sides++
				sideBroken = false
			}
		}
		lastPosition = position
	}
	return
}
func countBottomHorizontalSidesAfterSorting(positions []Position, positionMap map[Position]bool) (sides int) {
	sides++
	lastPosition := positions[0]
	sideBroken := false
	for _, position := range positions[1:] {
		_, bottomBlocked := positionMap[Position{y: position.y + 1, x: position.x}]
		if bottomBlocked {
			sideBroken = true
		} else {
			if lastPosition.y == position.y {
				if lastPosition.x-1 != position.x || sideBroken {
					// new side
					sides++
					sideBroken = false
				}
			} else {
				// new side
				sides++
				sideBroken = false
			}
		}
		lastPosition = position
	}
	return
}

func countLeftVerticalSidesAfterSorting(positions []Position, positionMap map[Position]bool) (sides int) {
	sides++
	lastPosition := positions[0]
	sideBroken := false
	for _, position := range positions[1:] {
		_, leftBlocked := positionMap[Position{y: position.y, x: position.x - 1}]
		if leftBlocked {
			sideBroken = true
		} else {
			if lastPosition.x == position.x {
				if lastPosition.y+1 != position.y || sideBroken {
					// new side
					sides++
					sideBroken = false
				}
			} else {
				// new side
				sides++
				sideBroken = false
			}
		}
		lastPosition = position
	}
	return
}
func countRightVerticalSidesAfterSorting(positions []Position, positionMap map[Position]bool) (sides int) {
	sides++
	lastPosition := positions[0]
	sideBroken := false
	for _, position := range positions[1:] {
		_, rightBlocked := positionMap[Position{y: position.y, x: position.x + 1}]
		if rightBlocked {
			sideBroken = true
		} else {
			if lastPosition.x == position.x {
				if lastPosition.y-1 != position.y || sideBroken {
					// new side
					sides++
					sideBroken = false
				}
			} else {
				// new side
				sides++
				sideBroken = false
			}
		}
		lastPosition = position
	}
	return
}

func getGarden() (garden []string) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	garden = append(garden, lines...)
	return
}
