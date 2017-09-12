// Coin Market Cap API in golang
package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	baseUrl  = "https://api.coinmarketcap.com/v1"
	graphUrl = "https://graphs.coinmarketcap.com/currencies"
	url      string
)

// Get information about the global market data of the cryptocurrencies.
func GetMarketData() (GlobalMarketData, error) {
	url = fmt.Sprintf(baseUrl + "/global/")

	resp, err := makeReq(url)

	var data GlobalMarketData
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return GlobalMarketData{}, err
	}

	return data, nil
}

// Get information about a crypto currency.
func GetCoinData(coin string) (Coin, error) {
	url = fmt.Sprintf("%s/ticker/%s", baseUrl, coin)
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

// Get information about all coins listed in Coin Market Cap.
func GetAllCoinData(limit int) (map[string]Coin, error) {
	var l string
	if limit > 0 {
		l = fmt.Sprintf("?limit=%v", limit)
	}
	url = fmt.Sprintf("%s/ticker/%s", baseUrl, l)

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

// Get graph data points for a crypto currency.
func GetCoinGraphData(coin string, start int64, end int64) (CoinGraph, error) {
	url = fmt.Sprintf("%s/%s/%d/%d", graphUrl, coin, start*1000, end*1000)
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
