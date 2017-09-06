package main

import (
	"fmt"
	"log"

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
	getBTC, err := coinApi.GetCoinInfo("bitcoin")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getBTC)

	}
	//GetAllCoinInfo(0) for all coins, GetAllCoinInfo(10) for top 10 coins & etc.
	getCoins, err := coinApi.GetAllCoinInfo(0)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCoins["ethereum"])
	}

	//GetCoinGraph for coin
	getGraph, err := coinApi.GetCoinGraph()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getGraph)
	}

}
