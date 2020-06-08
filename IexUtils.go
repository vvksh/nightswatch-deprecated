package main

import "fmt"

type StockInfo struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"companyName"`
	// PrimaryExchange        string      `json:"primaryExchange"`
	// CalculationPrice       string      `json:"calculationPrice"`
	// Open                   float64     `json:"open"`
	// OpenTime               int64       `json:"openTime"`
	// OpenSource             string      `json:"openSource"`
	Close float64 `json:"close"`
	// CloseTime              int64       `json:"closeTime"`
	// CloseSource            string      `json:"closeSource"`
	// High                   float64     `json:"high"`
	// HighTime               int64       `json:"highTime"`
	// HighSource             string      `json:"highSource"`
	// Low                    float64     `json:"low"`
	// LowTime                int64       `json:"lowTime"`
	// LowSource              string      `json:"lowSource"`
	LatestPrice float64 `json:"latestPrice"`
	// LatestSource           string      `json:"latestSource"`
	// LatestTime             string      `json:"latestTime"`
	// LatestUpdate           int64       `json:"latestUpdate"`
	LatestVolume int `json:"latestVolume"`
	// IexRealtimePrice       interface{} `json:"iexRealtimePrice"`
	// IexRealtimeSize        interface{} `json:"iexRealtimeSize"`
	// IexLastUpdated         interface{} `json:"iexLastUpdated"`
	// DelayedPrice           float64     `json:"delayedPrice"`
	// DelayedPriceTime       int64       `json:"delayedPriceTime"`
	// OddLotDelayedPrice     float64     `json:"oddLotDelayedPrice"`
	// OddLotDelayedPriceTime int64       `json:"oddLotDelayedPriceTime"`
	// ExtendedPrice          float64     `json:"extendedPrice"`
	// ExtendedChange         float64     `json:"extendedChange"`
	// ExtendedChangePercent  float64     `json:"extendedChangePercent"`
	// ExtendedPriceTime      int64       `json:"extendedPriceTime"`
	// PreviousClose          float64     `json:"previousClose"`
	// PreviousVolume         int         `json:"previousVolume"`
	// Change                 float64     `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	// Volume                 int         `json:"volume"`
	// IexMarketPercent       int         `json:"iexMarketPercent"`
	// IexVolume              interface{} `json:"iexVolume"`
	// AvgTotalVolume         int         `json:"avgTotalVolume"`
	// IexBidPrice            interface{} `json:"iexBidPrice"`
	// IexBidSize             interface{} `json:"iexBidSize"`
	// IexAskPrice            interface{} `json:"iexAskPrice"`
	// IexAskSize             interface{} `json:"iexAskSize"`
	// IexOpen                interface{} `json:"iexOpen"`
	// IexOpenTime            interface{} `json:"iexOpenTime"`
	// IexClose               float64     `json:"iexClose"`
	// IexCloseTime           int64       `json:"iexCloseTime"`
	// MarketCap              int64       `json:"marketCap"`
	// PeRatio                int         `json:"peRatio"`
	// Week52High             float64     `json:"week52High"`
	// Week52Low              float64     `json:"week52Low"`
	// YtdChange              float64     `json:"ytdChange"`
	// LastTradeTime          int64       `json:"lastTradeTime"`
	IsUSMarketOpen bool `json:"isUSMarketOpen"`
}

func (stockInfo *StockInfo) String() string {
	return fmt.Sprintf("%s Latest Price: %f \n Change Percent: %f \n Volume: %d \n <%s|Mobile> <%s|Web> \n\n", stockInfo.Symbol, stockInfo.LatestPrice, stockInfo.ChangePercent,
		stockInfo.LatestVolume, getRHMobileStockQuoteUrl(stockInfo.Symbol), getRHWebStockQuoteUrl(stockInfo.Symbol))
}

func getMostActiveApiUrl(IEX_API_KEY string) string {
	return fmt.Sprintf("https://cloud.iexapis.com/stable/stock/market/list/mostactive?token=%s&displayPercent=true", IEX_API_KEY)
}

func getGainersApiUrl(IEX_API_KEY string) string {
	return fmt.Sprintf("https://cloud.iexapis.com/stable/stock/market/list/gainers?token=%s&displayPercent=true", IEX_API_KEY)
}
