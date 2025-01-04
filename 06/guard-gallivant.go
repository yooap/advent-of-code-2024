package main

import (
	"fmt"
	"os"
	"strings"
)

type CoordsWithDirection struct {
	y, x      int
	direction rune
}

var ORIENTATION_PROGRESSION = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

var CHECKED_COORDS_FOR_OBSTICLE = map[CoordsWithDirection]bool{}

func main() {
	mapLayout := getMap()
	loops := executePatrol(mapLayout)
	count := countWalkedBlocks(mapLayout)
	fmt.Println(count)
	fmt.Println(loops)
}

func executePatrol(mapLayout [][]rune) (loops int) {
	coords := getInitialCoordinates(mapLayout)
	CHECKED_COORDS_FOR_OBSTICLE[CoordsWithDirection{y: coords.y, x: coords.x, direction: 0}] = true // do not care about direction here

	for !outOfBounds(coords, mapLayout) {
		coords, loops = move(coords, mapLayout, loops)
	}
	return
}

func getInitialCoordinates(mapLayout [][]rune) CoordsWithDirection {
	for y := range mapLayout {
		for x, value := range mapLayout[y] {
			if value == '^' {
				return CoordsWithDirection{y: y, x: x, direction: '^'}
			}
		}
	}
	panic("Missing guard")
}

func move(coords CoordsWithDirection, mapLayout [][]rune, loops int) (CoordsWithDirection, int) {
	for isPathObstructed(coords, mapLayout) {
		coords = turn(coords)
	}

	if canFormLoopIfObstructionAdded(coords, mapLayout) {
		loops++
	}

	// moved to coords
	nextCoords := getNextCoordsAccordingToDirection(coords)

	// set walked path unless initial
	mapLayout[coords.y][coords.x] = 'X'

	return nextCoords, loops
}

func getNextCoordsAccordingToDirection(coords CoordsWithDirection) CoordsWithDirection {
	direction := coords.direction
	var newCoords CoordsWithDirection
	if direction == '^' {
		newCoords = CoordsWithDirection{y: coords.y - 1, x: coords.x}
	} else if direction == '>' {
		newCoords = CoordsWithDirection{y: coords.y, x: coords.x + 1}
	} else if direction == 'v' {
		newCoords = CoordsWithDirection{y: coords.y + 1, x: coords.x}
	} else if direction == '<' {
		newCoords = CoordsWithDirection{y: coords.y, x: coords.x - 1}
	}
	newCoords.direction = coords.direction
	return newCoords
}

func canFormLoopIfObstructionAdded(coords CoordsWithDirection, mapLayout [][]rune) bool {
	obstructionCoords := getNextCoordsAccordingToDirection(coords)

	obstructionCoordsWithoutDirection := obstructionCoords
	obstructionCoordsWithoutDirection.direction = 0

	// global caching
	if CHECKED_COORDS_FOR_OBSTICLE[obstructionCoordsWithoutDirection] {
		return false
	}
	CHECKED_COORDS_FOR_OBSTICLE[obstructionCoordsWithoutDirection] = true

	if outOfBounds(obstructionCoords, mapLayout) {
		return false
	}

	// local caching
	visitedCoords := map[CoordsWithDirection]bool{}

	// add obstruction
	oldValue := mapLayout[obstructionCoords.y][obstructionCoords.x]
	mapLayout[obstructionCoords.y][obstructionCoords.x] = '#'

	for !visitedCoords[coords] && !outOfBounds(coords, mapLayout) {
		visitedCoords[coords] = true
		for isPathObstructed(coords, mapLayout) {
			coords = turn(coords)
		}
		coords = getNextCoordsAccordingToDirection(coords)
	}

	// remove obstruction
	mapLayout[obstructionCoords.y][obstructionCoords.x] = oldValue

	return visitedCoords[coords]
}

func isPathObstructed(coords CoordsWithDirection, mapLayout [][]rune) bool {
	direction := coords.direction
	if direction == '^' {
		return coords.y != 0 && mapLayout[coords.y-1][coords.x] == '#'
	}
	if direction == '>' {
		return coords.x != len(mapLayout[0])-1 && mapLayout[coords.y][coords.x+1] == '#'
	}
	if direction == 'v' {
		return coords.y != len(mapLayout)-1 && mapLayout[coords.y+1][coords.x] == '#'
	}
	if direction == '<' {
		return coords.x != 0 && mapLayout[coords.y][coords.x-1] == '#'
	}
	panic("Incorrect direction")
}

func turn(coords CoordsWithDirection) CoordsWithDirection {
	coords.direction = ORIENTATION_PROGRESSION[coords.direction]
	return coords
}

func outOfBounds(coords CoordsWithDirection, mapLayout [][]rune) bool {
	return coords.y < 0 || coords.x < 0 || coords.y >= len(mapLayout) || coords.x >= len(mapLayout[0])
}

func countWalkedBlocks(mapLayout [][]rune) (count int) {
	for y := range mapLayout {
		for x := range mapLayout[y] {
			if mapLayout[y][x] == 'X' {
				count++
			}
		}
	}
	return
}

func getMap() (mapLayout [][]rune) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		runes := []rune(line)
		mapLayout = append(mapLayout, runes)
	}
	return
}
