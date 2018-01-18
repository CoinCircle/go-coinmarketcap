# go-coinmarketcap

# coinmarketcap

    import "github.com/reteps/go-coinmarketcap"

Coin Market Cap API + Scraper in golang

## Usage

#### func  GetAllCoinData

```go
func GetAllCoinData(limit int) (map[string]Coin, error)
```
Get information about all coins listed in Coin Market Cap.

#### type Coin

```go
type Coin struct {
        ID               string  `json:"id"`
        Name             string  `json:"name"`
        Symbol           string  `json:"symbol"`
        Rank             int     `json:"rank,string"`
        PriceUsd         float64 `json:"price_usd,string"`
        PriceBtc         float64 `json:"price_btc,string"`
        Usd24hVolume     float64 `json:"24h_volume_usd,string"`
        MarketCapUsd     float64 `json:"market_cap_usd,string"`
        AvailableSupply  float64 `json:"available_supply,string"`
        TotalSupply      float64 `json:"total_supply,string"`
        PercentChange1h  float64 `json:"percent_change_1h,string"`
        PercentChange24h float64 `json:"percent_change_24h,string"`
        PercentChange7d  float64 `json:"percent_change_7d,string"`
        LastUpdated      string  `json:"last_updated"`
}
```

Coin struct

#### func  GetCoinData

```go
func GetCoinData(coin string) (Coin, error)
```
Get information about a crypto currency.

#### type CoinGraph

```go
type CoinGraph struct {
        MarketCapByAvailableAupply [][]float64 `json:"market_cap_by_available_supply"`
        PriceBtc                   [][]float64 `json:"price_btc"`
        PriceUsd                   [][]float64 `json:"price_usd"`
        VolumeUsd                  [][]float64 `json:"volume_usd"`
}
```

CoinGraph struct

#### func  GetCoinGraphData

```go
func GetCoinGraphData(coin string, start int64, end int64) (CoinGraph, error)
```
Get graph data points for a crypto currency.

#### type GlobalMarketData

```go
type GlobalMarketData struct {
        TotalMarketCapUsd            float64 `json:"total_market_cap_usd"`
        Total24hVolumeUsd            float64 `json:"total_24h_volume_usd"`
        BitcoinPercentageOfMarketCap float64 `json:"bitcoin_percentage_of_market_cap"`
        ActiveCurrencies             int     `json:"active_currencies"`
        ActiveAssets                 int     `json:"active_assets"`
        ActiveMarkets                int     `json:"active_markets"`
}
```

GlobalMarketData struct

#### func  GetMarketData

```go
func GetMarketData() (GlobalMarketData, error)
```

#### func CoinMarkets

```go
func CoinMarkets(coin string) ([]Markets, error)
```

# License

MIT
