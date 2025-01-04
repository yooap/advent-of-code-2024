package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list1, list2 := getBothLists()
	sortList(list1)
	sortList(list2)

	totalDistances := 0
	for i := range list1 {
		value1, value2 := list1[i], list2[i]
		distance := int(math.Abs(float64(value1 - value2)))
		totalDistances += distance
	}
	fmt.Println(totalDistances)

	similarityScore := 0
	for _, value1 := range list1 {
		frequency := 0
		for _, value2 := range list2 {
			if value1 == value2 {
				frequency++
			}
		}
		similarityScore += value1 * frequency
	}
	fmt.Println(similarityScore)
}

func sortList(list []int) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}

func getBothLists() (list1 []int, list2 []int) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		values := strings.Split(line, "   ")
		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])
		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}
	return
}
