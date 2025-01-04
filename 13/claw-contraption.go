package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	y, x int
}

type Vector = Position

type Machine struct {
	aButton, bButton Vector
	prize            Position
}

func main() {
	machines := getMachines()
	tokens1, tokens2 := 0, 0
	for _, machine := range machines {
		tokens1 += calculateOptimalTokenUseWithEquations(machine)
		machine = fixUnitConversionError(machine)
		tokens2 += calculateOptimalTokenUseWithEquations(machine)
	}
	fmt.Println(tokens1)
	fmt.Println(tokens2)
}

func fixUnitConversionError(machine Machine) Machine {
	machine.prize.x = machine.prize.x + 10000000000000
	machine.prize.y = machine.prize.y + 10000000000000
	return machine
}

// 12176 = A*66 + B*21
// 12748 = A*26 + B*67
// A = 135,395252838
// 12748 = ((12176 - 21B) / 66)*26 + B*67
// 12748 = 12176/66*26 - 21/66*26*B + B*67
// 12748 - 12176/66*26 = (-21/66*26 + 67)*B
// B = (12748 - 12176/66*26) / (-21/66*26 + 67)
// B = 135,395252838
// A = 141,4045407637

// 5400 = A*34 + B*67
// 8400 = A*94 + B*22
// A = (5400 - 67B) / 34
// 8400 = ((5400 - 67B) / 34)*94 + B*22
// 8400 = 5400/34*94 - 67/34*94*B + B*22
// 8400 - 5400/34*94 = (-67/34*94 + 22)*B
// B = (8400 - 5400/34*94) / (-67/34*94 + 22)
// B = 40
// A = 80
func calculateOptimalTokenUseWithEquations(machine Machine) int {
	roundingPercision := 1000.0

	ax, ay := float64(machine.aButton.x), float64(machine.aButton.y)
	bx, by := float64(machine.bButton.x), float64(machine.bButton.y)
	px, py := float64(machine.prize.x), float64(machine.prize.y)

	bPresses := (px - py*ax/ay) / (-1*by*ax/ay + bx)
	bPresses = math.Round(bPresses*roundingPercision) / roundingPercision
	if bPresses != math.Trunc(bPresses) {
		return 0
	}
	aPresses := (px - bPresses*bx) / ax
	aPresses = math.Round(aPresses*roundingPercision) / roundingPercision
	if aPresses != math.Trunc(aPresses) {
		return 0
	}

	return int(aPresses)*3 + int(bPresses)
}

func getMachines() (machines []Machine) {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	machineInputs := strings.Split(input, "\n\n")
	for _, machineInput := range machineInputs {
		machineData := strings.Split(machineInput, "\n")
		aButton := getVector(machineData[0])
		bButton := getVector(machineData[1])
		prize := getPosition(machineData[2])
		machines = append(machines, Machine{aButton: aButton, bButton: bButton, prize: prize})
	}
	return
}

func getVector(line string) Vector {
	line = line[10:]
	line = strings.ReplaceAll(line, "+", "")
	split := strings.Split(line, ", ")
	x, _ := strconv.Atoi(split[0][1:])
	y, _ := strconv.Atoi(split[1][1:])
	return Vector{y: y, x: x}
}

func getPosition(line string) Position {
	line = line[7:]
	line = strings.ReplaceAll(line, "=", "")
	split := strings.Split(line, ", ")
	x, _ := strconv.Atoi(split[0][1:])
	y, _ := strconv.Atoi(split[1][1:])
	return Position{y: y, x: x}
}
