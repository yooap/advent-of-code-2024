package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Position struct {
	y, x int
}

var numberToPosition = map[rune]Position{
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	' ': {3, 0},
	'0': {3, 1},
	'A': {3, 2},
}

var directionToPosition = map[rune]Position{
	' ': {0, 0},
	'^': {0, 1},
	'A': {0, 2},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}

type PATHS_CACHE_KEY struct {
	start, end Position
}

var PATHS_CACHE = map[PATHS_CACHE_KEY][][]rune{}

type RECURSION_CACHE_KEY struct {
	pathToken           string
	shortestPath, depth int
}

var RECURSION_CACHE = map[RECURSION_CACHE_KEY]int{}

func main() {
	codes := getCodes()
	complexitySum := 0
	for _, code := range codes {
		complexitySum += findComplexity(code)
	}
	fmt.Println(complexitySum)
}

func findComplexity(code []rune) int {
	keyStrokes := 0

	startKey := 'A'
	for _, numpadKey := range code {
		numpadPaths := getPossiblePathsForNumpad(numberToPosition[startKey], numberToPosition[numpadKey])
		shortestPath := -1
		for _, path := range numpadPaths {
			shortestPath = getShortestDirectionalKeyboardPressesRecursivly(path, shortestPath, 25)
		}
		startKey = numpadKey
		keyStrokes += shortestPath
	}
	mulitplier, _ := strconv.Atoi(string(code)[:3])
	return keyStrokes * mulitplier
}

func getShortestDirectionalKeyboardPressesRecursivly(path []rune, shortestPath, depth int) int {
	if cachedValue, ok := RECURSION_CACHE[RECURSION_CACHE_KEY{string(path), shortestPath, depth}]; ok {
		return cachedValue
	}
	shortestPathAtStartOfMethod := shortestPath

	if depth == 0 {
		if shortestPath == -1 || len(path) < shortestPath {
			shortestPath = len(path)
		}
	} else {
		pathLength := 0
		startKey := 'A'
		for _, key := range path {
			nextPaths := getPossiblePathsForDirectionalKeyboard(directionToPosition[startKey], directionToPosition[key])
			shortestPath := -1
			for _, nextPath := range nextPaths {
				shortestPath = getShortestDirectionalKeyboardPressesRecursivly(nextPath, shortestPath, depth-1)
			}
			pathLength += shortestPath
			startKey = key
		}
		if shortestPath == -1 || pathLength < shortestPath {
			shortestPath = pathLength
		}
	}

	RECURSION_CACHE[RECURSION_CACHE_KEY{string(path), shortestPathAtStartOfMethod, depth}] = shortestPath
	return shortestPath
}

func getPossiblePathsForDirectionalKeyboard(start, end Position) (paths [][]rune) {
	if cachedPaths, ok := PATHS_CACHE[PATHS_CACHE_KEY{start, end}]; ok {
		return cachedPaths
	}

	if start == end {
		paths = [][]rune{{'A'}}
	} else if start == directionToPosition[' '] {
		paths = [][]rune{}
	} else {
		current := start

		if current.x < end.x {
			next := Position{current.y, current.x + 1}
			nextPaths := deepClone(getPossiblePathsForDirectionalKeyboard(next, end))
			for i := range nextPaths {
				nextPaths[i] = slices.Insert(nextPaths[i], 0, '>')
			}
			paths = append(paths, nextPaths...)
		}
		if current.x > end.x {
			next := Position{current.y, current.x - 1}
			nextPaths := deepClone(getPossiblePathsForDirectionalKeyboard(next, end))
			for i := range nextPaths {
				nextPaths[i] = slices.Insert(nextPaths[i], 0, '<')
			}
			paths = append(paths, nextPaths...)
		}
		if current.y < end.y {
			next := Position{current.y + 1, current.x}
			nextPaths := deepClone(getPossiblePathsForDirectionalKeyboard(next, end))
			for i := range nextPaths {
				nextPaths[i] = slices.Insert(nextPaths[i], 0, 'v')
			}
			paths = append(paths, nextPaths...)
		}
		if current.y > end.y {
			next := Position{current.y - 1, current.x}
			nextPaths := deepClone(getPossiblePathsForDirectionalKeyboard(next, end))
			for i := range nextPaths {
				nextPaths[i] = slices.Insert(nextPaths[i], 0, '^')
			}
			paths = append(paths, nextPaths...)
		}
	}

	PATHS_CACHE[PATHS_CACHE_KEY{start, end}] = paths
	return
}

func deepClone(paths [][]rune) [][]rune {
	clone := make([][]rune, len(paths))
	for i, path := range paths {
		clone[i] = make([]rune, len(path))
		copy(clone[i], path)
	}
	return clone
}

func getPossiblePathsForNumpad(start, end Position) (paths [][]rune) {
	if start == end {
		return [][]rune{{'A'}}
	}
	if start == numberToPosition[' '] {
		return [][]rune{}
	}
	current := start

	if current.x < end.x {
		next := Position{current.y, current.x + 1}
		nextPaths := getPossiblePathsForNumpad(next, end)
		for i := range nextPaths {
			nextPaths[i] = slices.Insert(nextPaths[i], 0, '>')
		}
		paths = append(paths, nextPaths...)
	}
	if current.x > end.x {
		next := Position{current.y, current.x - 1}
		nextPaths := getPossiblePathsForNumpad(next, end)
		for i := range nextPaths {
			nextPaths[i] = slices.Insert(nextPaths[i], 0, '<')
		}
		paths = append(paths, nextPaths...)
	}
	if current.y < end.y {
		next := Position{current.y + 1, current.x}
		nextPaths := getPossiblePathsForNumpad(next, end)
		for i := range nextPaths {
			nextPaths[i] = slices.Insert(nextPaths[i], 0, 'v')
		}
		paths = append(paths, nextPaths...)
	}
	if current.y > end.y {
		next := Position{current.y - 1, current.x}
		nextPaths := getPossiblePathsForNumpad(next, end)
		for i := range nextPaths {
			nextPaths[i] = slices.Insert(nextPaths[i], 0, '^')
		}
		paths = append(paths, nextPaths...)
	}

	return
}

func getCodes() (codes [][]rune) {
	data, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(data), "\n")
	for _, row := range rows {
		codes = append(codes, []rune(row))
	}
	return
}
