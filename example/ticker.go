package main

import (
	"fmt"
	"log"

	cmc "github.com/coincircle/go-coinmarketcap"
)

func main() {
	// Get info about coin
	ticker, err := cmc.Ticker(&cmc.TickerOptions{
		Symbol:  "ETH",
		Convert: "EUR",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ticker.Name, ticker.Quotes["EUR"].Price)
}
