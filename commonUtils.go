package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
)

func getRHMobileStockQuoteUrl(stock string) string {
	return fmt.Sprintf("https://robinhood.com/applink/instrument/?symbol=%s", stock)
}

func getRHWebStockQuoteUrl(stock string) string {
	return fmt.Sprintf("https://robinhood.com/stocks/%s", stock)
}

func getMdMessage(title string, link string, date string, stock string) string {
	return fmt.Sprintf("%s <%s|mobile> <%s|web> \n <%s|%s: %s> \n\n\n", stock, getRHMobileStockQuoteUrl(stock), getRHWebStockQuoteUrl(stock), link, date, title)
}

func getArxivFormattedMessage(title string, link string, date string, descriptionHTML string) string {

	desc := sanitizeDescription(descriptionHTML)

	return fmt.Sprintf("%s \n <%s|%s> \n %s \n\n", date, link, title, desc)
}

func getStockTickr(item *gofeed.Item) string {
	for _, category := range item.Categories {
		if strings.Contains(category, "Nasdaq") || strings.Contains(category, "Nyse") {
			stockTickr := strings.Split(category, ":")[1]
			return stockTickr
		}
	}
	return ""
}

func isAfterhours() bool {
	return !(time.Now().Hour() > 6 && time.Now().Hour() < 13)
}

func getDurationTillStart() time.Duration {
	diffHr := 24 - (time.Now().Hour()) + 6
	diffMin := (time.Now().Minute())
	return time.Hour*time.Duration(diffHr) - time.Minute*time.Duration(diffMin)
}

func sanitizeDescription(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
		return html
	}

	output := ""
	// Find the review items
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			output += s.Text()
		}
	})
	if len(output) == 0 {
		log.Println("Sanitize didnt work")
		return html
	}

	if len(output) > 300 {
		output = output[0:300]
		output += "..."
	}
	return output
}
