package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Operation struct {
	left, right, op, out string
}

func main() {
	gates, ops := getGatesAndOperations()
	run(gates, ops)
	z := getNumberForGatesStartingWith(gates, 'z')
	fmt.Println(z)

	// part 2
	swaps := []string{}
	for _, op := range ops {
		if op.out[0] == 'z' {
			if op.out == "z45" {
				if op.op != "OR" {
					swaps = append(swaps, op.out)
				}
			} else if op.out == "z00" {
				if op.op != "XOR" {
					swaps = append(swaps, op.out)
				} else if (op.left != "x00" && op.left != "y00") || (op.right != "x00" && op.right != "y00") {
					swaps = append(swaps, op.out)
				}
			} else if op.op != "XOR" {
				swaps = append(swaps, op.out)
			} else if !checkOneOfTheInputsIsCorrectXOR(op, ops) {
				swaps = append(swaps, op.out)
			}
		} else if op.op == "XOR" {
			if op.left[0] != 'x' && op.left[0] != 'y' {
				swaps = append(swaps, op.out)
			}
		}
	}
	slices.Sort(swaps)
	fmt.Println(strings.Join(swaps, ","))
}

func checkOneOfTheInputsIsCorrectXOR(op Operation, ops []Operation) bool { // op = ... XOR ... -> zXX
	desiredX := "x" + op.out[1:]
	for _, opToCheckAgainst := range ops {
		if opToCheckAgainst.out == op.left || opToCheckAgainst.out == op.right {
			if opToCheckAgainst.op == "XOR" && (opToCheckAgainst.left == desiredX || opToCheckAgainst.right == desiredX) {
				return true
			}
		}
	}
	// either this opertaion is swapped, or operations related to it
	return false
}

func getNumberForGatesStartingWith(gates map[string]uint, char byte) int64 {
	selectedGates := []string{}
	for gate := range gates {
		if gate[0] == char {
			selectedGates = append(selectedGates, gate)
		}
	}
	slices.Sort(selectedGates)
	slices.Reverse(selectedGates)
	binaryString := ""
	for _, gate := range selectedGates {
		value := gates[gate]
		binaryString += strconv.Itoa(int(value))
	}

	value, _ := strconv.ParseInt(binaryString, 2, 64)
	return value
}

func run(gates map[string]uint, ops []Operation) {
	var unusedOps []Operation
	for len(ops) > 0 {
		unusedOps = []Operation{}
		for _, op := range ops {
			if leftGate, ok := gates[op.left]; ok {
				if rightGate, ok := gates[op.right]; ok {
					gates[op.out] = getOperationResult(op, leftGate, rightGate)
					continue
				}
			}
			unusedOps = append(unusedOps, op)
		}
		ops = unusedOps
	}
}

func getOperationResult(op Operation, leftGate, rightGate uint) uint {
	switch op.op {
	case "AND":
		return leftGate & rightGate
	case "OR":
		return leftGate | rightGate
	case "XOR":
		return leftGate ^ rightGate
	}
	panic("unexpected op")
}

func getGatesAndOperations() (gates map[string]uint, ops []Operation) {
	gates = map[string]uint{}

	data, _ := os.ReadFile("input.txt")
	gatesAndOps := strings.Split(string(data), "\n\n")

	for _, gateRow := range strings.Split(gatesAndOps[0], "\n") {
		gateCodeAndValue := strings.Split(gateRow, ": ")
		gateCode := gateCodeAndValue[0]
		value, _ := strconv.Atoi(gateCodeAndValue[1])
		gates[gateCode] = uint(value)
	}

	for _, opRow := range strings.Split(gatesAndOps[1], "\n") {
		opParts := strings.Split(opRow, " ")
		op := Operation{opParts[0], opParts[2], opParts[1], opParts[4]}
		ops = append(ops, op)
	}

	return
}
