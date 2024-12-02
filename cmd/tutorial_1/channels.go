package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE = 5

func main() {
	var chickenChannel = make(chan string)
	var websites = []string{"a.com", "b.com", "c.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}

	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep((time.Second * 1))
		var chickenPrice = rand.Int() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("\nFound a deal on chicken at %s", <-chickenChannel)
}
