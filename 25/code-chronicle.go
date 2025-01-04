package main

import (
	"fmt"
	"os"
	"strings"
)

type Lock = [5]int
type Key = [5]int

func main() {
	locks, keys := getLocksAndKeys()
	count := countPairsThatFit(locks, keys)
	fmt.Println(count)
}

func countPairsThatFit(locks []Lock, keys []Key) (count int) {
	for _, lock := range locks {
		for _, key := range keys {
			if fit(lock, key) {
				count++
			}
		}
	}
	return
}

func fit(lock Lock, key Key) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func getLocksAndKeys() (locks []Lock, keys []Key) {
	data, _ := os.ReadFile("input.txt")
	locksAndKeys := strings.Split(string(data), "\n\n")

	for _, lockOrKey := range locksAndKeys {
		parsedValue := parseKeyOrLock(lockOrKey)
		isLock := lockOrKey[0] == '#'
		if isLock {
			locks = append(locks, parsedValue)
		} else {
			keys = append(keys, parsedValue)
		}
	}

	return
}

func parseKeyOrLock(keyOrLockData string) (key [5]int) {
	rows := strings.Split(keyOrLockData, "\n")
	for y := 1; y < len(rows)-1; y++ {
		for x := 0; x < len(rows[y]); x++ {
			if rows[y][x] == '#' {
				key[x]++
			}
		}
	}
	return
}
