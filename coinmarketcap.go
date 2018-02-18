// Package coinmarketcap Coin Market Cap API fo Golang
package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

var (
	baseURL  = "https://api.coinmarketcap.com/v1"
	graphURL = "https://graphs2.coinmarketcap.com/currencies"
)

// GetMarketData get information about the global market data of the cryptocurrencies
func GetMarketData() (GlobalMarketData, error) {
	url := fmt.Sprintf(baseURL + "/global/")

	resp, err := makeReq(url)

	var data GlobalMarketData
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return GlobalMarketData{}, err
	}

	return data, nil
}

// GetCoinData get information about a crypto currency
func GetCoinData(coin string) (Coin, error) {
	coin = strings.ToLower(coin)
	url := fmt.Sprintf("%s/ticker/%s", baseURL, coin)
	resp, err := makeReq(url)
	if err != nil {
		return Coin{}, err
	}
	var data []Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return Coin{}, err
	}

	return data[0], nil
}

// GetAllCoinData get information about all coins listed in Coin Market Cap
func GetAllCoinData(limit int) (map[string]Coin, error) {
	var l string
	if limit >= 0 {
		l = fmt.Sprintf("?limit=%v", limit)
	}
	url := fmt.Sprintf("%s/ticker/%s", baseURL, l)

	resp, err := makeReq(url)

	var data []Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	// creating map from the array
	allCoins := make(map[string]Coin)
	for i := 0; i < len(data); i++ {
		allCoins[data[i].ID] = data[i]
	}

	return allCoins, nil
}

// GetCoinGraphData get graph data points for a crypto currency
func GetCoinGraphData(coin string, start int64, end int64) (CoinGraph, error) {
	url := fmt.Sprintf("%s/%s/%d/%d", graphURL, coin, start*1000, end*1000)
	resp, err := makeReq(url)
	if err != nil {
		return CoinGraph{}, err
	}
	var data CoinGraph
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return CoinGraph{}, err
	}

	return data, nil
}

// GetCoinPriceUsd get USD price of crypto currency
func GetCoinPriceUsd(coin string) (float64, error) {
	data, err := GetCoinData(coin)
	if err != nil {
		return float64(0), nil
	}
	return data.PriceUsd, nil
}

// CoinMarkets get market data for a coin name.
func CoinMarkets(coin string) ([]Market, error) {
	url := fmt.Sprintf("https://coinmarketcap.com/currencies/%s/#markets", coin)
	var markets []Market
	response, err := soup.Get(url)
	if err != nil {
		return nil, err
	}
	rows := soup.HTMLParse(response).Find("table", "id", "markets-table").Find("tbody").FindAll("tr")
	for _, row := range rows {
		var data []string
		for colNum, column := range row.FindAll("td") {
			for _, link := range column.FindAll("a") {
				data = append(data, strings.TrimSpace(link.Text()))
			}
			if colNum == 0 || colNum == 5 || colNum == 6 {
				data = append(data, column.Text())
			}
			for _, span := range column.FindAll("span") {
				data = append(data, strings.TrimSpace(span.Text()))
			}
		}
		markets = append(markets, Market{Rank: toInt(data[0]), Exchange: data[1], Pair: data[2], Volume: toInt(data[3]), Price: toFloat(data[4]), PercentVolume: toFloat(data[5]), Updated: (data[6] == "Recently")})
	}
	return markets, nil
}

// HTTP Client
func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

// HTTP Request Helper
func makeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := doReq(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// helper Function for CoinMarkets
func toInt(rawInt string) int {
	parsed, _ := strconv.Atoi(strings.Replace(strings.Replace(rawInt, "$", "", -1), ",", "", -1))
	return parsed
}

// helper Function for CoinMarkets
func toFloat(rawFloat string) float64 {
	parsed, _ := strconv.ParseFloat(strings.Replace(strings.Replace(strings.Replace(rawFloat, "$", "", -1), ",", "", -1), "%", "", -1), 64)
	return parsed
}
