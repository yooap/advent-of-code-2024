package main

import (
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	nodeMap, bounds := getNodes()

	antinodes := map[Coord]bool{}
	antinodesWithHarmonics := map[Coord]bool{}

	for _, coords := range nodeMap {
		for i, node1 := range coords {
			for ii, node2 := range coords {
				if i == ii {
					continue
				}

				distanceX := node2.x - node1.x
				antinodeX := node1.x - distanceX

				distanceY := node2.y - node1.y
				antinodeY := node1.y - distanceY

				antinode := Coord{x: antinodeX, y: antinodeY}

				antinodesWithHarmonics[node1] = true
				first := true
				for !outOfBounds(antinode, bounds) {
					if first {
						antinodes[antinode] = true
						first = false
					}
					antinodesWithHarmonics[antinode] = true
					newAntinodeX := antinode.x - distanceX
					newAntinodeY := antinode.y - distanceY
					antinode = Coord{x: newAntinodeX, y: newAntinodeY}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
	fmt.Println(len(antinodesWithHarmonics))
}

func outOfBounds(antinode, bounds Coord) bool {
	return antinode.x < 0 || antinode.x > bounds.x || antinode.y < 0 || antinode.y > bounds.y
}

func getNodes() (map[string][]Coord, Coord) {
	nodeMap := map[string][]Coord{}
	var maxX, maxY, y int

	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		nodes := strings.Split(line, "")
		for x, node := range nodes {
			maxX = max(x, maxX)
			if node == "." {
				continue
			}
			assertKeyExists(nodeMap, node)
			nodeMap[node] = append(nodeMap[node], Coord{x: x, y: y})
		}
		y++
	}

	maxY = y - 1
	bounds := Coord{x: maxX, y: maxY}
	return nodeMap, bounds
}

func assertKeyExists(nodeMap map[string][]Coord, node string) {
	if _, ok := nodeMap[node]; !ok {
		nodeMap[node] = []Coord{}
	}
}
