package main

import (
	"fmt"
	"log"
	"time"

	coinApi "github.com/miguelmota/go-coinmarketcap"
)

func main() {
	// Get global market data
	marketInfo, err := coinApi.GetMarketData()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(marketInfo)
	}

	// Get info about coin
	ethInfo, err := coinApi.GetCoinData("ethereum")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(ethInfo)

	}
	// Get top 10 coins
	top10, err := coinApi.GetAllCoinData(10)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(top10)
	}

	var threeMonths int64 = (59 * 60 * 24 * 60)
	now := time.Now()
	secs := now.Unix()
	start := secs - threeMonths
	end := secs

	// Get graph data for coin
	getGraph, err := coinApi.GetCoinGraphData("ethereum", start, end)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getGraph)
	}

}
