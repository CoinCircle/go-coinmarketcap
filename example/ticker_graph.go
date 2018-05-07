package main

import (
	"fmt"
	"log"
	"time"

	cmc "github.com/coincircle/go-coinmarketcap"
)

func main() {
	threeMonths := int64(60 * 60 * 24 * 90)
	now := time.Now()
	secs := now.Unix()
	start := secs - threeMonths
	end := secs

	// Get graph data for coin
	graph, err := cmc.TickerGraph(&cmc.TickerGraphOptions{
		Symbol: "ETH",
		Start:  start,
		End:    end,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(graph)
}
