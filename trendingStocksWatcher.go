package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/vvksh/amigo"
)

var IEX_API_KEY = ""

func init() {
	apikey, exists := os.LookupEnv("IEX_API_KEY")
	if !exists {
		log.Panicf("Couldn't find IEX_API_KEY env variable")
	}

	IEX_API_KEY = apikey
	var tsw TrendingStocksWatcher
	register(&tsw)
}

// type ActiveStocks []ActiveStock

// TrendingStocksWatcher implements Watcher interface
type TrendingStocksWatcher func()

func (tsw *TrendingStocksWatcher) Check() []string {
	iexGainersListUrl := getGainersApiUrl(IEX_API_KEY)

	var gainers []StockInfo

	err := amigo.CallHttpGetEndpoint(iexGainersListUrl, &gainers)

	if err != nil {
		return []string{err.Error()}
	}

	stockSymbols := fmt.Sprintf("%-10s|", "Symbols   ")
	latestPrice := fmt.Sprintf("%-14s|", "Price     ")
	changePercent := fmt.Sprintf("%-10s|", "Change%")
	// volume := fmt.Sprintf("%-10s", "Vol")
	rhURL := fmt.Sprintf("%-16s|", "url:")

	for _, gainer := range gainers {
		stockSymbols += fmt.Sprintf("%-12s|", gainer.Symbol)
		latestPrice += fmt.Sprintf("%-12.2f|", gainer.LatestPrice)
		changePercent += fmt.Sprintf("%-11.2f|", gainer.ChangePercent)
		// volume += fmt.Sprintf("%-12d|", gainer.LatestVolume)
		rhURL += fmt.Sprintf("<%s|m> , <%s|w>        |", getRHMobileStockQuoteUrl(gainer.Symbol), getRHWebStockQuoteUrl(gainer.Symbol))
	}
	updates := strings.Join([]string{stockSymbols, latestPrice, changePercent, rhURL}, "\n")
	return []string{updates}
}

func (tsw *TrendingStocksWatcher) Interval() time.Duration {

	if isAfterhours() {
		return getDurationTillStart()
	}
	return time.Minute * 30

}

func (tsw *TrendingStocksWatcher) SlackChannel() string {
	return "stockmovement"
}
