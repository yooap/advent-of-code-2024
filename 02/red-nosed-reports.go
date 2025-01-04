package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := getReports()
	safeCount := 0
	safeCountWithDampener := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
			safeCountWithDampener++
		} else {
			for i := range report {
				dampenedReport := []int{}
				dampenedReport = append(dampenedReport, report[:i]...)
				dampenedReport = append(dampenedReport, report[i+1:]...)
				if isSafe(dampenedReport) {
					safeCountWithDampener++
					break
				}
			}
		}
	}
	fmt.Println(safeCount)
	fmt.Println(safeCountWithDampener)
}

func isSafe(report []int) bool {
	lastDiff := 0
	for i, level := range report[1:] {
		lastLevel := report[i]
		diff := lastLevel - level
		if diff == 0 || math.Abs(float64(diff)) > 3 || !orderConsistant(lastDiff, diff) {
			return false
		}

		lastDiff = diff
	}
	return true
}

func orderConsistant(lastDiff, diff int) bool {
	if lastDiff == 0 {
		return true
	}
	return (lastDiff < 0 && diff < 0) || (lastDiff > 0 && diff > 0)
}

func getReports() (reports [][]int) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		levels := strings.Split(line, " ")
		report := []int{}
		for _, level := range levels {
			levelInt, _ := strconv.Atoi(level)
			report = append(report, levelInt)
		}
		reports = append(reports, report)
	}
	return
}
