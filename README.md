# go-coinmarketcap

> Coin Market Cap API in golang

# Install

```bash
go get -u -v github.com/miguelmota/go-coinmarketcap
```

# Import

```go
```

# Usage

Get single coin info

```go
ethInfo, err := cmcAPI.GetCoinInfo("ethereum")
```

Get Allcoins info (all/limit)

```go
getCoins, err := cmcAPI.GetAllCoinInfo(0)
```

If you want to limit results to top 10, use `cmcAPI.GetAllCoinInfo(10)`


Get Global Market Data

```go
getMarket, err := cmcAPI.GetMarketData()
```

# License

MIT