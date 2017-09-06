package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	baseUrl = "https://api.coinmarketcap.com/v1"
	graphUrl = "https://graphs.coinmarketcap.com/currencies"
	url     string
)

//GetMarketData - Get information about the global market data of the cryptocurrencies
func GetMarketData() (GlobalMarketData, error) {
	url = fmt.Sprintf(baseUrl + "/global/")

	resp, err := makeReq(url)

	var data GlobalMarketData
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println(err)
	}

	return data, nil
}

//GetCoinInfo - Get information about single crypto currency
func GetCoinInfo(coin string) (Coin, error) {
	url = fmt.Sprintf("%s/ticker/%s", baseUrl, coin)
	resp, err := makeReq(url)
	if err != nil {
		log.Println(err)
		return Coin{}, err
	}
	var data []Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println(err)
		return Coin{}, err
	}

	return data[0], nil
}

//GetAllCoinInfo - Get information about all coins listed in Coin Market Cap. If you want to limit the search to top 10 coins pass 10 as int, if you want all - pass 0 == No Limit
func GetAllCoinInfo(limit int) (map[string]Coin, error) {
	var l string
	if limit > 0 {
		l = fmt.Sprintf("?limit=%v", limit)
	}
	url = fmt.Sprintf("%s/ticker/%s", baseUrl, l)

	resp, err := makeReq(url)

	var data []Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println(err)
	}
	//creating map from the array
	allCoins := make(map[string]Coin)
	for i := 0; i < len(data); i++ {
		allCoins[data[i].ID] = data[i]
	}

	return allCoins, nil
}

//GetCoinGraph - Get graph data points for crypto currency
func GetCoinGraph(coin string) (CoinGraph, error) {
	start := "1483228800000"
	end := "1504670068000"
	url = fmt.Sprintf("%s/%s/%s/%s", graphUrl, coin, start, end)
	resp, err := makeReq(url)
	if err != nil {
		log.Println(err)
		return CoinGraph{}, err
	}
	var data CoinGraph
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println(err)
		return CoinGraph{}, err
	}

	return data, nil
}

//Client
func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func makeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := doReq(req)
	if err != nil {
		log.Println(err)
	}

	return resp, err
}
