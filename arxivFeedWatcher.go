package main

import (
	"time"

	"github.com/mmcdole/gofeed"
)

var feedURLs = []string{"http://arxiv.org/rss/cs.DB?version=2.0", "http://arxiv.org/rss/cs.DS?version=2.0"}

var checkedArticles = make(map[string]bool)

func init() {
	var arxivWatcher ArxivWatcher
	register(&arxivWatcher)
}

// GlobeNewsWireWatcher implements Watcher interface
type ArxivWatcher func()

func (arxivWatcher *ArxivWatcher) Check() []string {
	fp := gofeed.NewParser()
	// One update per feed
	updates := []string{}
	for _, rssfeedURL := range feedURLs {
		feed, _ := fp.ParseURL(rssfeedURL)
		date := feed.Published
		for _, item := range feed.Items {
			if _, ok := checkedArticles[item.Title]; !ok {
				update := getArxivFormattedMessage(item.Title, item.Link, date, item.Description)
				updates = append(updates, update)

				checkedArticles[item.Title] = true
			}
		}
	}

	//clear map if really large
	if len(checkedArticles) > 1000 {
		checkedArticles = make(map[string]bool)
	}
	return updates
}

func (arxivWatcher *ArxivWatcher) Interval() time.Duration {
	return time.Hour * 12
}

func (arxivWatcher *ArxivWatcher) SlackChannel() string {
	return "arxiv"
}
