package main

import (
	"fmt"
	"log"
	"time"

	cmc "github.com/miguelmota/go-coinmarketcap"
)

func main() {
	// Get global market data
	marketInfo, err := cmc.GetGlobalMarketData()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(marketInfo)
	}

	// Get info about coin
	coinInfo, err := cmc.GetCoinData("ethereum")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(coinInfo)
	}

	// Get top 10 coins
	top10, err := cmc.GetAllCoinData(10)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(top10)
	}

	var threeMonths int64 = (60 * 60 * 24 * 90)
	now := time.Now()
	secs := now.Unix()
	start := secs - threeMonths
	end := secs

	// Get graph data for coin
	coinGraphData, err := cmc.GetCoinGraphData("ethereum", start, end)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(coinGraphData)
	}

}
