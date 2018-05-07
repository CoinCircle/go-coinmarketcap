package main

import (
	"fmt"
	"log"

	cmc "github.com/coincircle/go-coinmarketcap"
)

func main() {
	// Get info about coin
	price, err := cmc.Price(&cmc.PriceOptions{
		Symbol:  "ETH",
		Convert: "EUR",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(price)
}
