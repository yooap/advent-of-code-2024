package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	compressedData := getDiskFragment()
	data := uncompress(compressedData)
	dataBackup := append([]int{}, data...)
	rearrangedData := rearrangeIndividualBlocks(data)
	checksum := calculateChecksum(rearrangedData)
	fmt.Println(checksum)

	rearrangedData = rearrangeWholeFiles(dataBackup)
	fmt.Println(rearrangedData)
	checksum = calculateChecksum(rearrangedData)
	fmt.Println(checksum)

}

func uncompress(compressedData []rune) (data []int) {
	for i, blockAsRune := range compressedData {
		block := int(blockAsRune - '0')
		var id int
		if i%2 == 0 {
			id = i / 2
		} else {
			id = -1
		}
		for j := 0; j < block; j++ {
			data = append(data, id)
		}
	}
	return
}

func rearrangeIndividualBlocks(data []int) []int {
	for i, value := range data {
		if value == -1 {
			lastElement := data[len(data)-1]
			data = slices.Replace(data, i, i+1, lastElement)
			data = data[:len(data)-1]
			data = removeTrailingFreeSpace(data)
		}
		if i == len(data)-1 {
			break
		}
	}
	return data
}

func removeTrailingFreeSpace(data []int) []int {
	for {
		lastElement := data[len(data)-1]
		if lastElement == -1 {
			data = data[:len(data)-1]
		} else {
			break
		}
	}
	return data
}

func rearrangeWholeFiles(data []int) []int {
	currentFile := []int{}
	for i := range data {
		index := len(data) - 1 - i
		value := data[index]
		if len(currentFile) == 0 {
			if value == -1 {
				continue
			}
			currentFile = append(currentFile, value)
		} else {
			if value == currentFile[0] {
				currentFile = append(currentFile, value)
			} else {
				data = tryToMoveFile(data, currentFile, index+1)
				currentFile = []int{}
				value := data[index] // get value again in case it was just overwritten
				if value == -1 {
					continue
				} else {
					currentFile = append(currentFile, value)
				}
			}
		}
	}
	data = removeTrailingFreeSpace(data)
	return data
}

func tryToMoveFile(data []int, currentFile []int, fileIndex int) []int {
	freeSpaceLength := 0
	for i, value := range data {
		if i == fileIndex {
			break
		}
		if value == -1 {
			freeSpaceLength++
			if freeSpaceLength == len(currentFile) {
				data = slices.Replace(data, i-freeSpaceLength+1, i+1, currentFile...)
				data = slices.Replace(data, fileIndex, fileIndex+freeSpaceLength, emptySpace(freeSpaceLength)...)
				return data
			}
		} else {
			freeSpaceLength = 0
		}
	}
	return data
}

func emptySpace(freeSpaceLength int) (emptySpace []int) {
	for i := 0; i < freeSpaceLength; i++ {
		emptySpace = append(emptySpace, -1)
	}
	return
}

func calculateChecksum(data []int) (checksum int) {
	for i, value := range data {
		if value == -1 {
			continue
		}
		checksum += i * value
	}
	return
}

func getDiskFragment() []rune {
	data, _ := os.ReadFile("input.txt")
	return []rune(string(data))
}
