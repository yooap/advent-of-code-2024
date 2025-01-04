package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var LAST_INDEX = 70

type Position struct {
	y, x int
}

func main() {
	bytes := getCorruptedBytes()
	byteMap := toMap(bytes[:1024])
	paths := dijkstra(byteMap)
	shortestPath := getShortestPathToEnd(paths)
	steps := len(shortestPath) - 1
	fmt.Println(steps)

	// part 2
	end := Position{LAST_INDEX, LAST_INDEX}
	byteIndexWhenPathBreaks := 1023
	for {
		byteIndexWhenPathBreaks := getCutOffByteIndex(shortestPath, bytes, byteIndexWhenPathBreaks)
		byteMap = toMap(bytes[:byteIndexWhenPathBreaks+1])
		paths = dijkstra(byteMap)
		if _, ok := paths[end]; !ok {
			lastByte := bytes[byteIndexWhenPathBreaks]
			fmt.Printf("%d,%d\n", lastByte.x, lastByte.y)
			break
		}
		shortestPath = getShortestPathToEnd(paths)
	}
}

func getCutOffByteIndex(shortestPath []Position, bytes []Position, lastByteIndex int) int {
	nodesOnShortestPath := toMap(shortestPath)
	for i := lastByteIndex + 1; true; i++ {
		byte := bytes[i]
		if _, ok := nodesOnShortestPath[byte]; ok {
			return i
		}
	}

	panic("out of bounds")
}

func getShortestPathToEnd(paths map[Position]Position) (shortestPath []Position) {
	node := Position{LAST_INDEX, LAST_INDEX}
	shortestPath = append(shortestPath, node)
	start := Position{0, 0}
	for node != start {
		node = paths[node]
		shortestPath = append(shortestPath, node)
	}
	return
}

func dijkstra(byteMap map[Position]bool) map[Position]Position {
	distances := map[Position]int{}
	queue := map[Position]bool{}
	previous := map[Position]Position{}
	for y := 0; y < LAST_INDEX+1; y++ {
		for x := 0; x < LAST_INDEX+1; x++ {
			pos := Position{y, x}
			if _, ok := byteMap[pos]; !ok {
				distances[pos] = -1
				queue[pos] = true
			}
		}
	}
	start, end := Position{0, 0}, Position{LAST_INDEX, LAST_INDEX}
	distances[start] = 0

	previousNode := Position{-1, -1}
	for len(queue) > 0 {
		node := findMin(distances, queue)
		delete(queue, node)
		if node == end || node == previousNode {
			break
		}

		for _, move := range []Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			movedToNode := Position{node.y + move.y, node.x + move.x}
			if outOfBounds(movedToNode) {
				continue
			}
			if _, ok := byteMap[movedToNode]; ok {
				continue // corrupted byte
			}
			if _, ok := queue[movedToNode]; !ok {
				continue // already traversed
			}

			newDistance := distances[node] + 1
			oldDistance := distances[movedToNode]
			if oldDistance == -1 || newDistance < oldDistance {
				distances[movedToNode] = newDistance
				previous[movedToNode] = node
			}
		}
		previousNode = node
	}

	return previous
}

func toMap(bytes []Position) map[Position]bool {
	byteMap := map[Position]bool{}
	for _, byte := range bytes {
		byteMap[byte] = true
	}
	return byteMap
}

func outOfBounds(node Position) bool {
	return node.y < 0 || node.y > LAST_INDEX || node.x < 0 || node.x > LAST_INDEX
}

func findMin(distance map[Position]int, queue map[Position]bool) (minNode Position) {
	min := -1
	for node := range queue {
		distance := distance[node]
		if distance == -1 {
			continue
		} else if min == -1 || distance < min {
			min = distance
			minNode = node
		}
	}
	return
}

func getCorruptedBytes() (bytes []Position) {
	data, _ := os.ReadFile("input.txt")
	bytesPositions := strings.Split(string(data), "\n")
	for _, byte := range bytesPositions {
		xy := strings.Split(byte, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		bytes = append(bytes, Position{y, x})
	}
	return
}
