package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	secrets := getSercrets()
	for i := range secrets {
		secrets[i] = evolve(secrets[i], 2000)
	}
	sum := sum(secrets)
	fmt.Println(sum)

	// part 2
	secrets = getSercrets()
	prices, changes := getPricesAndChanges(secrets, 2000)
	priceAtSequenceForEachSecret := getPriceAtSequenceForEachSecret(changes, prices)
	allSequences := getAllKeys(priceAtSequenceForEachSecret)
	max := calcualteMax(allSequences, priceAtSequenceForEachSecret)

	fmt.Println(max)
}

func calcualteMax(allSequences map[string]bool, priceAtSequenceForEachSecret []map[string]int) int {
	max := 0
	for sequence := range allSequences {
		sum := 0
		for _, priceAtSequenceForSecret := range priceAtSequenceForEachSecret {
			if price, ok := priceAtSequenceForSecret[sequence]; ok {
				sum += price
			}
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func getAllKeys(priceAtSequenceForEachSecret []map[string]int) map[string]bool {
	allSequences := map[string]bool{}
	for _, priceAtSequenceForSecret := range priceAtSequenceForEachSecret {
		for sequence := range priceAtSequenceForSecret {
			allSequences[sequence] = true
		}
	}
	return allSequences
}

func getPriceAtSequenceForEachSecret(changes [][]int, prices [][]int) []map[string]int {
	priceAtSequenceForEachSecret := []map[string]int{}
	for i, changesForSecret := range changes {
		for ii := range changesForSecret {
			if ii < 3 {
				continue
			}
			if ii == 3 {
				priceAtSequenceForEachSecret = append(priceAtSequenceForEachSecret, map[string]int{})
			}

			sequence := []int{changesForSecret[ii-3], changesForSecret[ii-2], changesForSecret[ii-1], changesForSecret[ii]}
			cacheKey := getCacheKey(sequence)
			if _, ok := priceAtSequenceForEachSecret[i][cacheKey]; !ok {
				priceAtSequenceForEachSecret[i][cacheKey] = prices[i][ii]
			}
		}
	}
	return priceAtSequenceForEachSecret
}

func getCacheKey(sequence []int) string {
	return fmt.Sprintf("%d_%d_%d_%d", sequence[0], sequence[1], sequence[2], sequence[3])
}

func getPricesAndChanges(secrets []int, iterations int) (prices [][]int, changes [][]int) {
	for si, secret := range secrets {
		prices = append(prices, []int{})
		changes = append(changes, []int{})

		oldPrice := secret % 10
		for i := 0; i < iterations; i++ {
			secret = evolveOnce(secret)
			price := secret % 10
			prices[si] = append(prices[si], price)
			changes[si] = append(changes[si], price-oldPrice)
			oldPrice = price
		}
	}

	return
}

func sum(secrets []int) (sum int) {
	for _, secret := range secrets {
		sum += secret
	}
	return
}

func evolve(secret int, iterations int) int {
	for i := 0; i < iterations; i++ {
		secret = evolveOnce(secret)
	}
	return secret
}

func evolveOnce(secret int) int {
	secret = ((secret * 64) ^ secret) % 16777216
	secret = ((secret / 32) ^ secret) % 16777216
	secret = ((secret * 2048) ^ secret) % 16777216
	return secret
}

func getSercrets() (secrets []int) {
	data, _ := os.ReadFile("input.txt")
	secretsAsString := strings.Split(string(data), "\n")
	for _, secretAsString := range secretsAsString {
		secret, _ := strconv.Atoi(secretAsString)
		secrets = append(secrets, secret)
	}
	return
}
