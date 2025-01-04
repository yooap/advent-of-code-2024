package main

import (
	"fmt"
	"os"
	"strings"
)

var CACHE = map[string]int{"": 1}

func main() {
	patterns, designs := getPatternsAndDesigns()
	possibleCount := 0
	allPossibilitiesCount := 0
	for _, design := range designs {
		count := getPossiblePickCount(design, patterns)
		if count > 0 {
			possibleCount++
			allPossibilitiesCount += count
		}
	}
	fmt.Println(possibleCount)
	fmt.Println(allPossibilitiesCount)
}

func getPossiblePickCount(design string, patterns []string) (count int) {
	if _, ok := CACHE[design]; ok {
		return CACHE[design]
	}

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			remainingDesign := design[len(pattern):]
			count += getPossiblePickCount(remainingDesign, patterns)
		}
	}

	CACHE[design] = count
	return count
}

func getPatternsAndDesigns() (patterns, designs []string) {
	data, _ := os.ReadFile("input.txt")
	splitData := strings.Split(string(data), "\n\n")
	patterns = strings.Split(splitData[0], ", ")
	designs = strings.Split(splitData[1], "\n")
	return
}
