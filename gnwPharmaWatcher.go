package main

import (
	"time"

	"github.com/mmcdole/gofeed"
)

const pharmaRSS = "https://www.globenewswire.com/RssFeed/industry/4577-Pharmaceuticals/feedTitle/GlobeNewswire%20-%20Industry%20News%20on%20Pharmaceuticals"
const pharmaChannelWebhook = "https://hooks.slack.com/services/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"

var checkedPharma = make(map[string]bool)

// GlobeNewsWireWatcher implements Watcher interface
type GNWPharmaWatcher func()

func init() {
	var gnwPharmaWatcher GNWPharmaWatcher
	register(&gnwPharmaWatcher)
}

func (gnwWatcher *GNWPharmaWatcher) Check() []string {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(pharmaRSS)
	updates := []string{}

	for _, item := range feed.Items {
		tickr := getStockTickr(item)
		if tickr != "" {
			if _, ok := checkedPharma[item.Title]; !ok {
				update := getMdMessage(item.Title, item.Link, item.Updated, tickr)
				updates = append(updates, update)
				checkedPharma[item.Title] = true
			}
		}

	}

	//clear map if really large
	if len(checkedPharma) > 1000 {
		checkedPharma = make(map[string]bool)
	}
	return updates
}

func (gnwWatcher *GNWPharmaWatcher) Interval() time.Duration {
	return time.Second * 10
}

func (gnwWatcher *GNWPharmaWatcher) SlackChannel() string {
	return "pharma"
}
