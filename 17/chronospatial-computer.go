package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	a, b, c uint // use uint, negative values are impossible
	program []uint
}

func main() {
	computer := getComputer()

	output := run(computer)

	outputAsStrings := []string{}
	for _, value := range output {
		outputAsStrings = append(outputAsStrings, strconv.Itoa(int(value)))
	}
	outputString := strings.Join(outputAsStrings, ",")
	fmt.Println(outputString)

	// part 2
	a := uint(0)
	prevA := a
	for i := 0; i < len(computer.program); i++ {
		programOutput := computer.program[len(computer.program)-1-i]
		for a>>3 == prevA {
			b := a & 7
			b = b ^ 5
			c := a >> b
			b = b ^ 6 ^ c
			if b&7 == programOutput {
				if i == len(computer.program)-1 {
					break
				}
				prevA = a
				a = a << 3
				break
			}
			a++
		}
		if a>>3 != prevA {
			// go back
			i -= 2
			a = prevA
			prevA = a >> 3
			a++
		}
	}
	fmt.Println(a)
}

func run(computer Computer) (out []uint) {
	for pointer := 0; pointer < len(computer.program); pointer += 2 {
		opcode, operand := computer.program[pointer], computer.program[pointer+1]
		switch opcode {
		case 0: // adv
			computer.a = doDivisionWithRegsiterA(operand, computer)
		case 1: // bxl
			computer.b = computer.b ^ operand
		case 2: // bst
			computer.b = getComboOperand(operand, computer) % 8
		case 3: // jnz
			if computer.a != 0 {
				pointer = int(operand) - 2
			}
		case 4: // bxc
			computer.b = computer.b ^ computer.c
		case 5: // out
			out = append(out, getComboOperand(operand, computer)%8)
		case 6: // bdv
			computer.b = doDivisionWithRegsiterA(operand, computer)
		case 7: // cdv
			computer.c = doDivisionWithRegsiterA(operand, computer)
		}
	}

	return
}

func doDivisionWithRegsiterA(operand uint, computer Computer) uint {
	denominator := float64(getComboOperand(operand, computer))
	a := float64(computer.a) / math.Pow(2, denominator)
	return uint(a)
}

func getComboOperand(operand uint, computer Computer) uint {
	switch operand {
	case 4:
		return computer.a
	case 5:
		return computer.b
	case 6:
		return computer.c
	}
	return operand
}

func getComputer() (computer Computer) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	registerA, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	registerB, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	registerC, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])
	computer.a = uint(registerA)
	computer.b = uint(registerB)
	computer.c = uint(registerC)

	computer.program = []uint{}
	for _, value := range strings.Split(strings.Split(lines[4], ": ")[1], ",") {
		instruction, _ := strconv.Atoi(value)
		computer.program = append(computer.program, uint(instruction))
	}

	return
}
