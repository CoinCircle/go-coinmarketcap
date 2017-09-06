package main

import (
	"fmt"
	"log"
	"time"

	coinApi "github.com/miguelmota/go-coinmarketcap"
)

func main() {
	//Get global market data
	getMarket, err := coinApi.GetMarketData()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getMarket)
	}
	//Get info about coin
	getBTC, err := coinApi.GetCoinData("bitcoin")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getBTC)

	}
	//GetAllCoinInfo(0) for all coins, GetAllCoinInfo(10) for top 10 coins & etc.
	getCoins, err := coinApi.GetAllCoinData(0)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCoins["ethereum"])
	}

	now := time.Now()
	secs := now.Unix()
	start := secs - (60 * 60 * 24 * 60)
	end := secs

	//GetCoinGraph for coin
	getGraph, err := coinApi.GetCoinGraphData("ethereum", start, end)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getGraph)
	}

}
