package main

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

var watchers []Watcher

func register(watcher Watcher) {
	watchers = append(watchers, watcher)
}

func main() {
	for _, watcher := range watchers {
		go watch(watcher)
	}
	// var techWatcher GNWTechWatcher
	// var pharmaWatcher GNWPharmaWatcher
	// go watch(&techWatcher)
	// go watch(&pharmaWatcher)
	select {}
}

func watch(watcher Watcher) {
	log.Println("watching now")
	for {
		updates := watcher.Check()
		handleUpdates(updates, watcher.SlackChannel())
		r := rand.Intn(500)
		time.Sleep(watcher.Interval() + time.Duration(r)*time.Millisecond)
	}

}

func handleUpdates(updates []string, channel string) {
	// for _, update := range
	fullMessage := strings.Join(updates, "\n")
	if len(updates) > 0 {
		err := SendSlackNotification(fullMessage, channel)
		if err != nil {
			log.Printf(err.Error())
		}
	} else {
		log.Printf("\nNo new updates\n")
	}
}
