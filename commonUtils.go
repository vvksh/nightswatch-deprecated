package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
	"github.com/slack-go/slack"
)

func CallApi(apiEndpoint string, responseObject interface{}) error {
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return readErr
	}
	return json.Unmarshal(body, responseObject)
}

func getRHMobileStockQuoteUrl(stock string) string {
	return fmt.Sprintf("https://robinhood.com/applink/instrument/?symbol=%s", stock)
}

func getRHWebStockQuoteUrl(stock string) string {
	return fmt.Sprintf("https://robinhood.com/stocks/%s", stock)
}

func getMdMessage(title string, link string, date string, stock string) string {
	return fmt.Sprintf("%s <%s|mobile> <%s|web> \n <%s|%s: %s> \n\n\n", stock, getRHMobileStockQuoteUrl(stock), getRHWebStockQuoteUrl(stock), link, date, title)
}

func SendSlackNotification(msg string, channel string) error {
	webhookURL, exists := os.LookupEnv("SLACK_WEBHOOK")
	if !exists {
		log.Panicf("Environment variable SLACK_WEBHOOK not found\n")
	}
	webHookMessage := slack.WebhookMessage{}
	webHookMessage.Text = msg
	webHookMessage.Channel = channel
	return slack.PostWebhook(webhookURL, &webHookMessage)
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
