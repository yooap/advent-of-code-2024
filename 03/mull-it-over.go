package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	mem := getCorruptedMemInstructions()

	// part 1
	multiplicationSum := calculateMultiplicationSum(mem)
	fmt.Println(multiplicationSum)

	// part 2
	onMatcher := regexp.MustCompile(`do\(\)`)
	onMatches := onMatcher.FindAllStringIndex(mem, -1)
	ons := flatten(onMatches)
	ons = append([]int{0}, ons...) // enabled at the start

	offMatcher := regexp.MustCompile(`don't\(\)`)
	offMatches := offMatcher.FindAllStringIndex(mem, -1)
	offs := flatten(offMatches)

	reducedMem := ""
	for _, on := range ons {
		for _, off := range offs {
			if off < on {
				offs = offs[1:]
			} else {
				break
			}
		}

		if len(offs) == 0 {
			reducedMem += mem[on:]
			break
		}

		rangeEnd := offs[0]
		if len(ons) > 1 && ons[1] < rangeEnd {
			rangeEnd = ons[1]
		}

		reducedMem += mem[on:rangeEnd]
		ons = ons[1:]
	}

	multiplicationSum = calculateMultiplicationSum(reducedMem)
	fmt.Println(multiplicationSum)
}

func flatten(list [][]int) (result []int) {
	for _, element := range list {
		result = append(result, element[0])
	}
	return
}

func calculateMultiplicationSum(mem string) (multiplicationSum int) {
	mulMatcher := regexp.MustCompile(`mul\(([0-9]+,[0-9]+)\)`)
	mulMatches := mulMatcher.FindAllStringSubmatch(mem, -1)
	for _, match := range mulMatches {
		multipliers := strings.Split(match[1], ",")
		multiplier1, _ := strconv.Atoi(multipliers[0])
		multiplier2, _ := strconv.Atoi(multipliers[1])
		multiplicationSum += multiplier1 * multiplier2
	}
	return
}

func getCorruptedMemInstructions() string {
	mem, _ := os.ReadFile("input.txt")
	return string(mem)
}
