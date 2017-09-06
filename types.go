package coinmarketcap

//Coin struct
type Coin struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Rank             int     `json:"rank,string"`
	PriceUsd         float64 `json:"price_usd,string"`
	PriceBtc         float64 `json:"price_btc,string"`
	Two4HVolumeUsd   float64 `json:"24h_volume_usd,string"`
	MarketCapUsd     float64 `json:"market_cap_usd,string"`
	AvailableSupply  float64 `json:"available_supply,string"`
	TotalSupply      float64 `json:"total_supply,string"`
	PercentChange1H  float64 `json:"percent_change_1h,string"`
	PercentChange24H float64 `json:"percent_change_24h,string"`
	PercentChange7D  float64 `json:"percent_change_7d,string"`
	LastUpdated      string  `json:"last_updated"`
}

//GlobalMarketData struct
type GlobalMarketData struct {
	TotalMarketCapUsd float64 `json:"total_market_cap_usd"`
	Total24HVolumeUsd float64 `json:"total_24h_volume_usd"`
	BitcoinDominance  float64 `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies  int     `json:"active_currencies"`
	ActiveAssets      int     `json:"active_assets"`
	ActiveMarkets     int     `json:"active_markets"`
}
