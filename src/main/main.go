package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/willf/bloom"
)

const totalItems = uint(1000)

var filter = bloom.New(20*totalItems, 5) // load of 20, 5 keys

func main() {
	var gnwWatcher GlobeNewsWireWatcher
	watch(&gnwWatcher)

}

func watch(watcher Watcher) {
	for {
		updates := watcher.Check()
		handleUpdates(updates)
		r := rand.Intn(500)
		time.Sleep(watcher.Interval() + time.Duration(r)*time.Millisecond)
	}

}

func handleUpdates(updates []Update) {
	// for _, update := range
	fullMessage := ""
	for _, update := range updates {
		if !filter.Test([]byte(getFilterKey(update.Title))) {
			fullMessage += getMessage(update)
			log.Printf("%v\n", update)
			filter.Add([]byte(getFilterKey(update.Title)))
		}
	}

	if len(fullMessage) > 0 {
		SendSlackNotification(fullMessage)
	} else {
		log.Printf("\nNo new updates\n")
	}
}

func getFilterKey(s string) string {
	if len(s) > 10 {
		return s[0:10]
	}
	return s
}
func getMessage(update Update) string {
	return fmt.Sprintf("<%s|%s: %s> \n", update.URL, update.Date, update.Title)
}
