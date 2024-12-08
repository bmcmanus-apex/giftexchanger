package main

import (
	"fmt"
	"math/rand"
	"time"
)

var invalidPairs = map[string][]string{
	"Addie":    {"Cayson", "Charlotte"},
	"Cayson":   {"Addie", "Charlotte"},
	"Charlotte": {"Addie", "Cayson"},
	"Alyssa":   {"Kendall"},
	"Kendall":  {"Alyssa"},
	"Colton":   {"Hope", "Caleb", "Nash"},
	"Hope":     {"Colton", "Caleb", "Nash"},
	"Caleb":    {"Colton", "Hope", "Nash"},
	"Nash":     {"Colton", "Hope", "Caleb"},
}

func isValidPair(giver, receiver string) bool {
	for _, invalidReceiver := range invalidPairs[giver] {
		if receiver == invalidReceiver {
			return false
		}
	}
	return true
}

func createGiftExchange(names []string) map[string]string {
	rand.Seed(time.Now().UnixNano())
	shuffledNames := make([]string, len(names))
	copy(shuffledNames, names)

	for {
		rand.Shuffle(len(shuffledNames), func(i, j int) {
			shuffledNames[i], shuffledNames[j] = shuffledNames[j], shuffledNames[i]
		})

		valid := true
		for i := range names {
			if names[i] == shuffledNames[i] || !isValidPair(names[i], shuffledNames[i]) {
				valid = false
				break
			}
		}

		if valid {
			break
		}
	}

	giftExchange := make(map[string]string)
	for i := range names {
		giftExchange[names[i]] = shuffledNames[i]
	}

	return giftExchange
}

func main() {
	names := []string{"Alyssa", "Colton", "Addie", "Noah", "Hope", "Kendall", "Caleb", "Charlotte", "Cayson", "Nash"}
	giftExchange := createGiftExchange(names)

	for giver, receiver := range giftExchange {
		fmt.Printf("%s will give a gift to %s\n", giver, receiver)
	}
}
