package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Position struct {
	y, x int
}

func main() {
	nodes, start, end := getRaceTrackNodes()
	path := constructPath(nodes, start, end)

	timesSavedByShortcuts := getTimesSavedByShortcuts(path, nodes, 2)
	count := getNubmerOfShortcutsWhichSaveAtLeast(100, timesSavedByShortcuts)
	fmt.Println(count)

	// part 2
	timesSavedByShortcuts = getTimesSavedByShortcuts(path, nodes, 20)
	count = getNubmerOfShortcutsWhichSaveAtLeast(100, timesSavedByShortcuts)
	fmt.Println(count)
}

func getNubmerOfShortcutsWhichSaveAtLeast(limit int, timesSavedByShortcuts []int) int {
	for i := 0; i < len(timesSavedByShortcuts); i++ {
		if timesSavedByShortcuts[i] < limit {
			return i
		}
	}
	panic("not found")
}

func getTimesSavedByShortcuts(path []Position, nodes map[Position]bool, steps int) (shortcutSavedTime []int) {
	for i, node := range path {
		allNodesReachedByCuttingThrough := findAllNodesReachableByCuttingThrough(node, nodes, steps)
		for reachableNode := range allNodesReachedByCuttingThrough {
			distanceBetweenTheseNodes := abs(reachableNode.y-node.y) + abs(reachableNode.x-node.x)
			savedMoves := slices.Index(path, reachableNode) - i - distanceBetweenTheseNodes
			if savedMoves > 0 {
				shortcutSavedTime = append(shortcutSavedTime, savedMoves)
			}
		}
	}

	slices.Sort(shortcutSavedTime)
	slices.Reverse(shortcutSavedTime)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findAllNodesReachableByCuttingThrough(node Position, nodes map[Position]bool, steps int) map[Position]bool {
	possibleToReachNodes := map[Position]bool{}
	for yOffset := -steps; yOffset <= steps; yOffset++ {
		for xOffset := -steps + abs(yOffset); xOffset <= steps-abs(yOffset); xOffset++ {
			possibleNode := Position{node.y + yOffset, node.x + xOffset}
			if _, ok := nodes[possibleNode]; ok {
				possibleToReachNodes[possibleNode] = true
			}
		}
	}
	return possibleToReachNodes
}

func constructPath(nodes map[Position]bool, start, end Position) (path []Position) {
	path = append(path, start)
	current := start
	for current != end {
		var previousNode Position
		if len(path) < 2 {
			previousNode = start
		} else {
			previousNode = path[len(path)-2]
		}
		nextInPath := getNextInPath(nodes, current, previousNode)
		path = append(path, nextInPath)
		current = nextInPath
	}
	return
}

func getNextInPath(nodes map[Position]bool, current, previous Position) Position {
	for _, direction := range []Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		neighbor := Position{current.y + direction.y, current.x + direction.x}
		if neighbor == previous {
			continue
		}
		if _, ok := nodes[neighbor]; ok {
			return neighbor
		}
	}
	panic("No path found")
}

func getRaceTrackNodes() (trackNodes map[Position]bool, start, end Position) {
	trackNodes = make(map[Position]bool)
	data, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(data), "\n")
	for y, row := range rows {
		for x, char := range row {
			switch char {
			case 'S':
				start = Position{y, x}
			case 'E':
				end = Position{y, x}
			case '.':
				trackNodes[Position{y, x}] = true
			}
		}
	}
	trackNodes[start] = true
	trackNodes[end] = true
	return
}
