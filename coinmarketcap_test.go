package coinmarketcap

import (
	"testing"
)

func TestGetMarketData(t *testing.T) {

}

func TestGetCoinData(t *testing.T) {

}

func TestGetAllCoinData(t *testing.T) {

}

func TestGetCoinGraphData(t *testing.T) {

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

func TestCoinMarkets(t *testing.T) {

}

func TestDoReq(t *testing.T) {

}

func TestMakeReq(t *testing.T) {

}

func TestToInt(t *testing.T) {

}

func TestToFloat(t *testing.T) {

}
