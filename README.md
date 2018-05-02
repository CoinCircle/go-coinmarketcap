<h1 align="center">
  <img src="https://user-images.githubusercontent.com/168240/39501128-e66e2a18-4d6d-11e8-9e16-88655102da6c.png" alt="go-coinmarketcap" width="600" />
</h1>

> The Unofficial [CoinMarketCap](https://coinmarketcap.com/) API client for [Go](https://golang.org/).

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/go-coinmarketcap/master/LICENSE.md) [![Build Status](https://travis-ci.org/miguelmota/go-coinmarketcap.svg?branch=master)](https://travis-ci.org/miguelmota/go-coinmarketcap) [![Go Report Card](https://goreportcard.com/badge/github.com/miguelmota/go-coinmarketcap?)](https://goreportcard.com/report/github.com/miguelmota/go-coinmarketcap) [![GoDoc](https://godoc.org/github.com/miguelmota/go-coinmarketcap?status.svg)](https://godoc.org/github.com/miguelmota/go-coinmarketcap)

## Documentation

[https://godoc.org/github.com/miguelmota/go-coinmarketcap](https://godoc.org/github.com/miguelmota/go-coinmarketcap)

## Install

```bash
go get -u github.com/miguelmota/go-coinmarketcap
```

## Getting started

```go
package main

import (
	"fmt"
	"log"

	cmc "github.com/miguelmota/go-coinmarketcap"
)

func main() {
	coins, err := cmc.GetAllCoinData(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, coin := range coins {
		fmt.Println(coin.Symbol, coin.PriceUSD)
	}
}
```

## Examples

Check out the [`./example`](./example) directory and documentation.

## License

MIT
