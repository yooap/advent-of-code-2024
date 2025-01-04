package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	y, x int
}
type Vector struct {
	y, x int
}

func main() {
	warehouse, moves := getWarehouseAndMoves()
	warehouse = moveRobot(warehouse, moves)
	result := getGpsValueSum(warehouse, 'O')
	fmt.Println(result)

	// part 2
	warehouse, moves = getWarehouseAndMoves()
	warehouse = expandWarehouse(warehouse)
	warehouse = moveRobotAfterExpansion(warehouse, moves)
	result = getGpsValueSum(warehouse, '[')
	fmt.Println(result)
}

// func printWarehouse(warehouse [][]rune, robotPosition Position) {
// 	warehouse[robotPosition.y][robotPosition.x] = '@'
// 	for _, line := range warehouse {
// 		fmt.Println(string(line))
// 	}
// 	fmt.Println()
// }

func moveRobotAfterExpansion(warehouse [][]rune, moves string) [][]rune {
	robotPosition := getRobotPosition(warehouse)
	for _, move := range moves {
		if move == '^' {
			warehouse, robotPosition = moveUpAfterExpansion(warehouse, robotPosition)
		} else if move == '>' {
			warehouse, robotPosition = moveRightAfterExpansion(warehouse, robotPosition)
		} else if move == 'v' {
			warehouse, robotPosition = moveDownAfterExpansion(warehouse, robotPosition)
		} else if move == '<' {
			warehouse, robotPosition = moveLeftAfterExpansion(warehouse, robotPosition)
		}
	}
	return warehouse
}

func moveUpAfterExpansion(warehouse [][]rune, robotPosition Position) ([][]rune, Position) {
	return moveWithVectorAfterExpansion(warehouse, robotPosition, Vector{-1, 0})
}

func moveRightAfterExpansion(warehouse [][]rune, robotPosition Position) ([][]rune, Position) {
	return moveWithVectorAfterExpansion(warehouse, robotPosition, Vector{0, 1})

}

func moveDownAfterExpansion(warehouse [][]rune, robotPosition Position) ([][]rune, Position) {
	return moveWithVectorAfterExpansion(warehouse, robotPosition, Vector{1, 0})
}

func moveLeftAfterExpansion(warehouse [][]rune, robotPosition Position) ([][]rune, Position) {
	return moveWithVectorAfterExpansion(warehouse, robotPosition, Vector{0, -1})
}

func moveWithVectorAfterExpansion(warehouse [][]rune, robotPosition Position, vector Vector) ([][]rune, Position) {
	canMove := true
	initialMove := Position{robotPosition.y + vector.y, robotPosition.x + vector.x}
	positionsToMove := []Position{initialMove}
	moveCache := map[Position]bool{initialMove: true}

	for i := 0; true; i++ {
		if len(positionsToMove) == i {
			break
		}
		positionToMove := positionsToMove[i]
		moveCache[positionToMove] = true
		valueToMove := warehouse[positionToMove.y][positionToMove.x]

		if valueToMove == '#' {
			canMove = false
			break
		}

		if valueToMove == '.' {
			continue
		}

		// moving horizontally
		if vector.x != 0 {
			positionsToMove = append(positionsToMove, Position{positionToMove.y, positionToMove.x + vector.x})
			continue
		}

		// moving vertically
		if valueToMove == '[' {
			otherHalfOfBox := Position{positionToMove.y, positionToMove.x + 1}
			if _, ok := moveCache[otherHalfOfBox]; !ok {
				positionsToMove = append(positionsToMove, otherHalfOfBox)
			}
		} else if valueToMove == ']' {
			otherHalfOfBox := Position{positionToMove.y, positionToMove.x - 1}
			if _, ok := moveCache[otherHalfOfBox]; !ok {
				positionsToMove = append(positionsToMove, otherHalfOfBox)
			}
		}
		nextPosition := Position{positionToMove.y + vector.y, positionToMove.x}
		if _, ok := moveCache[nextPosition]; !ok {
			positionsToMove = append(positionsToMove, nextPosition)
		}
	}

	if canMove {
		warehouse[robotPosition.y][robotPosition.x] = '.'
		for len(positionsToMove) > 0 {
			positionToMove := positionsToMove[len(positionsToMove)-1]
			positionsToMove = positionsToMove[:len(positionsToMove)-1]

			if warehouse[positionToMove.y][positionToMove.x] == '.' {
				continue
			} else {
				warehouse[positionToMove.y+vector.y][positionToMove.x+vector.x] = warehouse[positionToMove.y][positionToMove.x]
				warehouse[positionToMove.y][positionToMove.x] = '.'
			}
		}

		robotPosition = Position{robotPosition.y + vector.y, robotPosition.x + vector.x}
	}

	return warehouse, robotPosition
}

func expandWarehouse(warehouse [][]rune) (expandedWarehouse [][]rune) {
	for _, line := range warehouse {
		expandedLine := []rune{}
		for _, value := range line {
			if value == 'O' {
				expandedLine = append(expandedLine, '[', ']')
			} else if value == '@' {
				expandedLine = append(expandedLine, '@', '.')
			} else {
				expandedLine = append(expandedLine, value, value)
			}
		}
		expandedWarehouse = append(expandedWarehouse, expandedLine)
	}
	return
}

func getGpsValueSum(warehouse [][]rune, valueToLookFor rune) (sum int) {
	for y := range warehouse {
		for x, value := range warehouse[y] {
			if value == valueToLookFor {
				sum += y*100 + x
			}
		}
	}
	return
}

func moveRobot(warehouse [][]rune, moves string) [][]rune {
	robotPosition := getRobotPosition(warehouse)
	for _, move := range moves {
		if move == '^' {
			warehouse, robotPosition = moveUp(warehouse, robotPosition)
		} else if move == '>' {
			warehouse, robotPosition = moveRight(warehouse, robotPosition)
		} else if move == 'v' {
			warehouse, robotPosition = moveDown(warehouse, robotPosition)
		} else if move == '<' {
			warehouse, robotPosition = moveLeft(warehouse, robotPosition)
		}
	}
	return warehouse
}

func moveUp(warehouse [][]rune, robotPosition Position) ([][]rune, Position) {
	return moveWithVector(warehouse, robotPosition, Vector{-1, 0})
}

func moveRight(warehouse [][]rune, robotPosition Position) ([][]rune, Position) {
	return moveWithVector(warehouse, robotPosition, Vector{0, 1})

}

func moveDown(warehouse [][]rune, robotPosition Position) ([][]rune, Position) {
	return moveWithVector(warehouse, robotPosition, Vector{1, 0})
}

func moveLeft(warehouse [][]rune, robotPosition Position) ([][]rune, Position) {
	return moveWithVector(warehouse, robotPosition, Vector{0, -1})
}

func moveWithVector(warehouse [][]rune, robotPosition Position, vector Vector) ([][]rune, Position) {
	for i := 1; true; i++ {
		nextValue := warehouse[robotPosition.y+vector.y*i][robotPosition.x+vector.x*i]
		if nextValue == '#' {
			// blocked
			break
		} else if nextValue == 'O' {
			// keep moving
			continue
		} else {
			warehouse[robotPosition.y][robotPosition.x] = '.'
			if i > 1 {
				warehouse[robotPosition.y+(vector.y*i)][robotPosition.x+(vector.x*i)] = 'O'
			}
			robotPosition = Position{y: robotPosition.y + vector.y, x: robotPosition.x + vector.x}
			break
		}
	}
	return warehouse, robotPosition
}

func getRobotPosition(warehouse [][]rune) Position {
	for y, _ := range warehouse {
		for x, value := range warehouse[y] {
			if value == '@' {
				return Position{y: y, x: x}
			}
		}
	}
	panic("Robot missing")
}

func getWarehouseAndMoves() (warehouse [][]rune, moves string) {
	data, _ := os.ReadFile("input.txt")
	warehouseAndMoves := strings.Split(string(data), "\n\n")
	warehouseLines := strings.Split(warehouseAndMoves[0], "\n")
	for _, line := range warehouseLines {
		warehouse = append(warehouse, []rune(line))
	}
	moves = strings.ReplaceAll(warehouseAndMoves[1], "\n", "")
	return
}
