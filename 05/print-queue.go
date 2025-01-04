package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, updates := getRulesAndUpdates()

	// part 1
	result1, result2 := 0, 0
	for _, update := range updates {
		valid := isUpdateValid(update, rules)
		if valid {
			result1 += getMiddlePage(update)
		} else {
			fixedUpdate := fixUpdate(update, rules)
			result2 += getMiddlePage(fixedUpdate)
		}
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

func isUpdateValid(update []int, rules map[int][]int) bool {
	for i, page := range update {
		afters := rules[page]
		for _, previousPage := range update[:i] {
			if slices.Contains(afters, previousPage) {
				return false
			}
		}
	}
	return true
}

func fixUpdate(update []int, rules map[int][]int) []int {
	for i := 0; i < len(update); i++ {
		page := update[i]
		afters := rules[page]
		for ii, previousPage := range update[:i] {
			if slices.Contains(afters, previousPage) {
				update = slices.Delete(update, ii, ii+1)
				update = slices.Insert(update, i, previousPage)
				i--
			}
		}
	}
	return update
}

func getMiddlePage(update []int) int {
	return update[(len(update)-1)/2]
}

func getRulesAndUpdates() (rules map[int][]int, updates [][]int) {
	rules = make(map[int][]int)

	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, "|") {
			xy := strings.Split(line, "|")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			rules[x] = append(rules[x], y)
		} else if strings.Contains(line, ",") {
			pages := strings.Split(line, ",")
			update := []int{}
			for _, pageStr := range pages {
				page, _ := strconv.Atoi(pageStr)
				update = append(update, page)
			}
			updates = append(updates, update)
		}
	}
	return
}
