package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pair struct {
	left, right string
}

type Triplet struct {
	left, middle, right string
}

func main() {
	pairs := getPairs()
	individualUnits := getIndividualUnits(pairs)
	relations := generateRelationMap(individualUnits, pairs)
	triplets := findThreeWayRelations(relations)
	tripletsStartingWithT := 0
	for triplet := range triplets {
		if anyStartWith(triplet, 't') {
			tripletsStartingWithT++
		}
	}
	fmt.Println(tripletsStartingWithT)

	// part 2
	unitsWithMostMutualRelations := findUnitsWithMostMutualRelations(relations)
	pw := strings.Join(unitsWithMostMutualRelations, ",")
	fmt.Println(pw)
}

func findUnitsWithMostMutualRelations(relations map[string]map[string]bool) []string {
	networks := []map[string]bool{}
	for unit := range relations {
		networks = append(networks, map[string]bool{unit: true})
	}

	for _, network := range networks {
		for unit := range relations {
			if unitFitsInNetwork(unit, network, relations) {
				network[unit] = true
			}
		}
	}

	max := 0
	maxIndex := 0
	for i, network := range networks {
		if len(network) > max {
			max = len(network)
			maxIndex = i
		}
	}

	bestNetwork := networks[maxIndex]
	bestNetworkSlice := []string{}
	for unit := range bestNetwork {
		bestNetworkSlice = append(bestNetworkSlice, unit)
	}
	slices.Sort(bestNetworkSlice)
	return bestNetworkSlice
}

func unitFitsInNetwork(unit string, network map[string]bool, relations map[string]map[string]bool) bool {
	if _, ok := network[unit]; ok {
		// unit alredy in net
		return true
	}

	relationsForUnit := relations[unit]

	for unitInNet := range network {
		if _, ok := relationsForUnit[unitInNet]; !ok {
			return false
		}
	}
	return true
}

func anyStartWith(triplet Triplet, char byte) bool {
	return triplet.left[0] == char || triplet.middle[0] == char || triplet.right[0] == char
}

func findThreeWayRelations(relations map[string]map[string]bool) map[Triplet]bool {
	triplets := map[Triplet]bool{}
	for unit := range relations {
		relationsForUnit := relations[unit]
		for unit2 := range relationsForUnit {
			relationsForUnit2 := relations[unit2]
			for unit3 := range relationsForUnit2 {
				relationsForUnit3 := relations[unit3]
				if _, ok := relationsForUnit3[unit]; ok {
					triplet := createTriplet(unit, unit2, unit3)
					triplets[triplet] = true
				}
			}
		}
	}
	return triplets
}

func createTriplet(unit1, unit2, unit3 string) (triplet Triplet) {
	slice := []string{unit1, unit2, unit3}
	slices.Sort(slice)
	triplet.left = slice[0]
	triplet.middle = slice[1]
	triplet.right = slice[2]
	return triplet
}

func generateRelationMap(individualUnits map[string]bool, pairs []Pair) map[string]map[string]bool {
	relations := map[string]map[string]bool{}
	for unit := range individualUnits {
		relations[unit] = map[string]bool{}
		for _, pair := range pairs {
			if pair.left == unit {
				relations[unit][pair.right] = true
			} else if pair.right == unit {
				relations[unit][pair.left] = true
			}
		}
	}
	return relations
}

func getIndividualUnits(pairs []Pair) map[string]bool {
	units := map[string]bool{}
	for _, pair := range pairs {
		units[pair.left] = true
		units[pair.right] = true
	}
	return units
}

func getPairs() (pairs []Pair) {
	data, _ := os.ReadFile("input.txt")
	pairRows := strings.Split(string(data), "\n")
	for _, pairRow := range pairRows {
		pair := strings.Split(pairRow, "-")
		pairs = append(pairs, Pair{pair[0], pair[1]})
	}
	return
}
