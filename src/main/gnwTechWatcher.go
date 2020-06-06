package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/willf/bloom"
)

const techRSS = "https://www.globenewswire.com/AtomFeed/industry/9000-Technology/feedTitle/GlobeNewswire%20-%20Industry%20News%20on%20Technology"
const techChannelWebhook = "https://hooks.slack.com/services/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"

var techFilter = bloom.New(20*uint(1000), 5) // load of 20, 5 keys

// GlobeNewsWireWatcher implements Watcher interface
type GNWTechWatcher func()

func (gnwWatcher *GNWTechWatcher) Check() []string {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(techRSS)
	updates := []string{}

	for _, item := range feed.Items {
		tickr := getStockTickr(item)
		if tickr != "" {
			if !techFilter.Test([]byte(item.Title)) {
				update := getMdMessage(item.Title, item.Link, item.Updated, tickr)
				updates = append(updates, update)
				techFilter.Add([]byte(item.Title))
			}
		}
	}
	return updates
}

func (gnwWatcher *GNWTechWatcher) Interval() time.Duration {
	return time.Minute * 5
}

func (gnwWatcher *GNWTechWatcher) WebHookUrl() string {
	return techChannelWebhook
}

func getMdMessage(title string, link string, date string, stock string) string {
	return fmt.Sprintf("%s <%s|mobile> <%s|web> \n <%s|%s: %s> \n\n\n", stock, getRHWebStockQuoteUrl(stock), getRHMobileStockQuoteUrl(stock), link, date, title)
}
