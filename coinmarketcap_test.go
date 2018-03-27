package coinmarketcap

import (
	"testing"
	"time"
)

func TestGetMarketData(t *testing.T) {
	market, err := GetMarketData()
	if err != nil {
		t.FailNow()
	}

	if market.TotalMarketCapUSD == 0 {
		t.FailNow()
	}
}

func TestGetMarketGraphData(t *testing.T) {
	var threeMonths int64 = (60 * 60 * 24 * 90)
	end := time.Now().Unix()
	start := end - threeMonths

	graph, err := GetMarketGraphData(start, end)
	if err != nil {
		t.FailNow()
	}

	if graph.VolumeUSD[0][0] == 0 {
		t.FailNow()
	}
}

func TestGetCoinData(t *testing.T) {
	coin, err := GetCoinData("ethereum")
	if err != nil {
		t.FailNow()
	}

	if coin.PriceUSD == 0 {
		t.FailNow()
	}
}

func TestGetAllCoinData(t *testing.T) {
	coins, err := GetAllCoinData(10)
	if err != nil {
		t.FailNow()
	}

	if len(coins) != 10 {
		t.FailNow()
	}
}

func TestGetCoinGraphData(t *testing.T) {
	var threeMonths int64 = (60 * 60 * 24 * 90)
	end := time.Now().Unix()
	start := end - threeMonths

	graph, err := GetCoinGraphData("ethereum", start, end)
	if err != nil {
		t.FailNow()
	}

	if graph.PriceUSD[0][0] == 0 {
		t.FailNow()
	}
}

func TestGetCoinPriceUSD(t *testing.T) {
	price, err := GetCoinPriceUSD("ethereum")
	if err != nil {
		t.FailNow()
	}
	if price <= 0 {
		t.FailNow()
	}
}

func TestGetCoinMarkets(t *testing.T) {
	markets, err := GetCoinMarkets("ethereum")
	if err != nil {
		t.FailNow()
	}
	if len(markets) == 0 {
		t.FailNow()
	}

	market := markets[0]
	if market.Rank == 0 {
		t.FailNow()
	}
	if market.Exchange == "" {
		t.FailNow()
	}
	if market.Pair == "" {
		t.FailNow()
	}
	if market.VolumeUSD == 0 {
		t.FailNow()
	}
	if market.Price == 0 {
		t.FailNow()
	}
	if market.VolumePercent == 0 {
		t.FailNow()
	}
	if market.Updated == "" {
	}
}

func TestDoReq(t *testing.T) {
	// TODO
}

func TestMakeReq(t *testing.T) {
	// TODO
}

func TestToInt(t *testing.T) {
	v := toInt("5")
	if v != 5 {
		t.FailNow()
	}
}

func TestToFloat(t *testing.T) {
	v := toFloat("5.2")
	if v != 5.2 {
		t.FailNow()
	}
}
