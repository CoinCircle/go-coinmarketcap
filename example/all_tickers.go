package main

import (
	"fmt"
	"log"

	cmc "github.com/coincircle/go-coinmarketcap"
)

func main() {
	// get data for all coins
	coins, err := cmc.Tickers(&cmc.TickersOptions{
		Start:   0,
		Limit:   100,
		Convert: "USD",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, coin := range coins {
		fmt.Println(coin.Symbol, coin.Quotes["USD"].Price)
	}
}
