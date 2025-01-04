package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Equation struct {
	result int
	values []int
}

func main() {
	equations := getEquations()
	sum1, sum2 := 0, 0
	for _, equation := range equations {
		if isPossible(equation, false) {
			sum1 += equation.result
			sum2 += equation.result
		} else if isPossible(equation, true) {
			sum2 += equation.result
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func isPossible(equation Equation, useConcatenation bool) bool {
	values := equation.values
	expectedResult := equation.result
	return checkIsPossibleRecursivly(values, expectedResult, useConcatenation)
}

func checkIsPossibleRecursivly(values []int, expectedResult int, useConcatenation bool) bool {
	if len(values) == 1 {
		return values[0] == expectedResult
	}

	// sum
	sum := values[0] + values[1]
	reducedValues := slices.Replace(slices.Clone(values), 0, 2, sum)
	if checkIsPossibleRecursivly(reducedValues, expectedResult, useConcatenation) {
		return true
	}

	// multiply
	multiplication := values[0] * values[1]
	reducedValues = slices.Replace(slices.Clone(values), 0, 2, multiplication)
	if checkIsPossibleRecursivly(reducedValues, expectedResult, useConcatenation) {
		return true
	}

	if !useConcatenation {
		return false
	}

	// multiply
	concatenation := concatenation(values[0], values[1])
	reducedValues = slices.Replace(slices.Clone(values), 0, 2, concatenation)
	return checkIsPossibleRecursivly(reducedValues, expectedResult, useConcatenation)
}

func concatenation(a, b int) int {
	concatenationStr := strconv.Itoa(a) + strconv.Itoa(b)
	concatenation, _ := strconv.Atoi(concatenationStr)
	return concatenation
}

func getEquations() (equations []Equation) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		resultAndValues := strings.Split(line, ": ")
		result, _ := strconv.Atoi(resultAndValues[0])
		values := []int{}
		for _, valueStr := range strings.Split(resultAndValues[1], " ") {
			value, _ := strconv.Atoi(valueStr)
			values = append(values, value)
		}
		equation := Equation{result: result, values: values}
		equations = append(equations, equation)
	}
	return
}
