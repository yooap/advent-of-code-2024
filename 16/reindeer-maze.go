package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	y, x int
}

func main() {
	maze := getMaze()
	neighbors, start, end := constuctGraph(maze)
	previousNodeMap, distanceMap := dijkstra(neighbors, start, end)
	score := distanceMap[end]
	fmt.Println(score)

	// part 2
	validTiles := getValidTiles(previousNodeMap, start, end)
	previousNodeMap, _ = dijkstra(neighbors, end, start) // reverse
	validTilesReverse := getValidTiles(previousNodeMap, end, start)
	for tile := range validTilesReverse {
		validTiles[tile] = true
	}
	fmt.Println(len(validTiles))
}

func getValidTiles(previousNodeMap map[Node]map[Node]bool, start, end Node) map[Node]bool {
	validTiles := map[Node]bool{start: true}
	tilesToCheck := []Node{end}
	for len(tilesToCheck) > 0 {
		tile := tilesToCheck[0]
		tilesToCheck = tilesToCheck[1:]
		validTiles[tile] = true

		for refTile := range previousNodeMap[tile] {
			tilesToCheck = append(tilesToCheck, refTile)
			continue
		}
	}

	// print
	// dimensions := 142
	// for y := 0; y < dimensions; y++ {
	// 	fmt.Println()
	// 	for x := 0; x < dimensions; x++ {
	// 		if _, ok := validTiles[Node{y, x}]; ok {
	// 			fmt.Print("O")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	}
	// }

	return validTiles
}

func dijkstra(neighbors map[Node][]Node, start, end Node) (previousNodeMap map[Node]map[Node]bool, distanceMap map[Node]int) {
	previousNodeMap = map[Node]map[Node]bool{}
	distanceMap = map[Node]int{}
	queue := map[Node]bool{}
	for node := range neighbors {
		distanceMap[node] = -1
		queue[node] = true
	}
	distanceMap[start] = 0

	for {
		node := findMin(distanceMap, queue)
		delete(queue, node)
		if node == end {
			break
		}

		neighborsForNode := neighbors[node]
		for _, neighbor := range neighborsForNode {
			if _, ok := queue[neighbor]; !ok {
				continue
			}

			previousNodes, ok := previousNodeMap[node]
			if !ok {
				if node.x+1 == neighbor.x {
					distanceMap[neighbor] = 1
				} else {
					distanceMap[neighbor] = 1001
				}
				previousNodeMap[neighbor] = map[Node]bool{node: true}
				continue
			}

			for previousNode := range previousNodes {
				newDistance := distanceMap[node] + distanceBetween(node, neighbor, previousNode, end, distanceMap)
				oldDistance := distanceMap[neighbor]
				if oldDistance == -1 || newDistance < oldDistance {
					distanceMap[neighbor] = newDistance
					previousNodeMap[neighbor] = map[Node]bool{node: true}
				} else if newDistance == oldDistance {
					previousNodeMap[neighbor][node] = true
				}
			}
		}
	}
	return
}

func distanceBetween(node, neighbor, previousNode, end Node, allNodeMap map[Node]int) int {
	if neighbor == end {
		return 1
	}

	if previousNode.x == node.x && node.x == neighbor.x {
		if _, ok := allNodeMap[Node{neighbor.y + neighbor.y - node.y, neighbor.x}]; !ok {
			return 1001
		}
		return 1
	} else if previousNode.y == node.y && node.y == neighbor.y {
		if _, ok := allNodeMap[Node{neighbor.y, neighbor.x + neighbor.x - node.x}]; !ok {
			return 1001
		}
		return 1
	}

	if previousNode.x == node.x {
		if _, ok := allNodeMap[Node{y: node.y + node.y - previousNode.y, x: node.x}]; ok {
			return 1001
		}
	} else if previousNode.y == node.y {
		if _, ok := allNodeMap[Node{y: node.y, x: node.x + node.x - previousNode.x}]; ok {
			return 1001
		}
	}
	return 1
}

func findMin(distanceMap map[Node]int, queue map[Node]bool) (minNode Node) {
	min := -1
	for node := range queue {
		distance := distanceMap[node]
		if distance == -1 {
			continue
		} else if min == -1 || distance < min {
			min = distance
			minNode = node
		}
	}
	return
}

func constuctGraph(maze [][]rune) (neighbors map[Node][]Node, start, end Node) {
	neighbors = map[Node][]Node{}
	for y := range maze {
		for x, value := range maze[y] {
			if value == '#' {
				continue
			}
			thisNode := Node{y, x}
			if value == 'S' {
				start = thisNode
			} else if value == 'E' {
				end = thisNode
			}

			neighbors[thisNode] = getNeigbors(thisNode, maze)
		}
	}
	return
}

func getNeigbors(source Node, maze [][]rune) (neighbors []Node) {
	// up
	if maze[source.y-1][source.x] != '#' {
		neighbors = append(neighbors, Node{source.y - 1, source.x})
	}
	// right
	if maze[source.y][source.x+1] != '#' {
		neighbors = append(neighbors, Node{source.y, source.x + 1})
	}
	// down
	if maze[source.y+1][source.x] != '#' {
		neighbors = append(neighbors, Node{source.y + 1, source.x})
	}
	// left
	if maze[source.y][source.x-1] != '#' {
		neighbors = append(neighbors, Node{source.y, source.x - 1})
	}
	return
}

func getMaze() (maze [][]rune) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		maze = append(maze, []rune(line))
	}
	return
}
