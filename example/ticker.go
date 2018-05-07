package main

import (
	"fmt"
	"log"

	cmc "github.com/coincircle/go-coinmarketcap"
)

func main() {
	// Get info about coin
	coin, err := cmc.Ticker(&cmc.TickerOptions{
		Symbol:  "ETH",
		Convert: "EUR",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coin.Name, coin.Quotes["EUR"].Price)
}
