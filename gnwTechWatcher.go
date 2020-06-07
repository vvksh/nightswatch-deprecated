package main

import (
	"time"

	"github.com/mmcdole/gofeed"
)

const techRSS = "https://www.globenewswire.com/AtomFeed/industry/9000-Technology/feedTitle/GlobeNewswire%20-%20Industry%20News%20on%20Technology"

var checkedTech = make(map[string]bool)

func init() {
	var gnwTechWatcher GNWTechWatcher
	register(&gnwTechWatcher)
}

// GlobeNewsWireWatcher implements Watcher interface
type GNWTechWatcher func()

func (gnwWatcher *GNWTechWatcher) Check() []string {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(techRSS)
	updates := []string{}

	for _, item := range feed.Items {
		tickr := getStockTickr(item)
		if tickr != "" {
			if _, ok := checkedTech[item.Title]; !ok {
				update := getMdMessage(item.Title, item.Link, item.Updated, tickr)
				updates = append(updates, update)
				checkedTech[item.Title] = true
			}
		}
	}

	//clear map if really large
	if len(checkedPharma) > 1000 {
		checkedPharma = make(map[string]bool)
	}
	return updates
}

func (gnwWatcher *GNWTechWatcher) Interval() time.Duration {
	return time.Second * 10
}

func (gnwWatcher *GNWTechWatcher) SlackChannel() string {
	return "technology"
}
