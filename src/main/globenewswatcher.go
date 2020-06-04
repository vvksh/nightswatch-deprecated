package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "https://www.globenewswire.com"

const searchResultsURL = baseURL + "/Search/NewsSearch?lang=en&country=UNITED%20STATES&exchange=Nasdaq%2CNYSE"

// GlobeNewsWireWatcher implements Watcher interface
type GlobeNewsWireWatcher func()

func (gnwWatcher *GlobeNewsWireWatcher) Check() []Update {
	resp, err := http.Get(searchResultsURL)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	return parseGNWSearchResultsHTML(resp.Body)
}

func (gnwWatcher *GlobeNewsWireWatcher) Interval() time.Duration {
	return time.Minute * 5
}

func parseGNWSearchResultsHTML(r io.Reader) []Update {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	updates := []Update{}
	// Find the review items
	doc.Find(".results-link").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		update := Update{}
		parseUpdate(s, &update)
		updates = append(updates, update)
	})
	return updates
}

func parseUpdate(selection *goquery.Selection, update *Update) {
	item := selection.Find("a")
	title := item.Text()
	link, _ := item.Attr("href")
	date := parseDate(link)
	fullLink := baseURL + link
	// body := selection.Find("p")
	// log.Printf("title:%s \n link:%s\n body:%s\n", title, link, body.Text())

	update.Date = date
	update.Title = title
	update.URL = fullLink
}

func parseDate(url string) string {
	subStrings := strings.Split(url, "/")
	if len(subStrings) > 5 {
		return fmt.Sprintf("%s-%s-%s", subStrings[2], subStrings[3], subStrings[4])

	}
	log.Printf("Couldn't parse date from %s", url)
	return time.Now().Local().Format("2006-01-02")

}
