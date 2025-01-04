package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stones := getStones()
	stoneMap := map[int]int{}
	for _, stone := range stones {
		changeCount(stoneMap, stone, 1)
	}

	for i := 0; i < 75; i++ {
		updatedStoneMap := map[int]int{}
		for stone, count := range stoneMap {
			if stone == 0 {
				changeCount(updatedStoneMap, 1, count)
			} else {
				stoneString := strconv.Itoa(stone)
				if len(stoneString)%2 == 0 {
					newStone1, newStone2 := splitStone(stoneString)
					changeCount(updatedStoneMap, newStone1, count)
					changeCount(updatedStoneMap, newStone2, count)
				} else {
					newStone := stone * 2024
					changeCount(updatedStoneMap, newStone, count)
				}
			}
		}
		stoneMap = updatedStoneMap
		if i == 24 {
			fmt.Println(getStoneCount(stoneMap))
		}
	}
	fmt.Println(getStoneCount(stoneMap))
}

func splitStone(stoneString string) (int, int) {
	newStone1String, newStone2String := stoneString[:len(stoneString)/2], stoneString[len(stoneString)/2:]
	newStone1, _ := strconv.Atoi(newStone1String)
	newStone2, _ := strconv.Atoi(newStone2String)
	return newStone1, newStone2
}

func changeCount(stoneMap map[int]int, stone, countChange int) {
	if count, ok := stoneMap[stone]; ok {
		stoneMap[stone] = count + countChange
	} else {
		stoneMap[stone] = countChange
	}
}

func getStoneCount(stoneMap map[int]int) (totalCount int) {
	for _, count := range stoneMap {
		totalCount += count
	}
	return totalCount
}

func getStones() (stones []int) {
	data, _ := os.ReadFile("input.txt")
	stonesString := strings.Split(string(data), " ")
	for _, stoneString := range stonesString {
		stone, _ := strconv.Atoi(stoneString)
		stones = append(stones, stone)
	}
	return
}
